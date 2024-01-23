package spider

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"example.cn/grabblock/common"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/tools/log"
)

type CollectorOption struct{}
type Collector struct {
	wg                   *sync.WaitGroup
	lock                 *sync.RWMutex
	requestCallBacks     []RequestCallBack
	blockCallBacks       []BlockCallBack
	transactionCallBacks []TransactionCallBack
	jsonCallBacks        []JsonCallBack
	errorCallbacks       []ErrorCallback
	backend              *httpBackend
	loadBalance          []string
	loadIndex            int
}
type RequestCallBack func(*Request)
type BlockCallBack func(*Response, *entity.RetBlockMessage)
type TransactionCallBack func(*Response, *entity.RetTransactionMessage)
type JsonCallBack func(*Response, *entity.RetBlockMessage)
type ErrorCallback func(*Response, error)

func NewCollector(options ...CollectorOption) *Collector {
	s := &Collector{}
	s.Init()
	return s
}
func (s *Collector) SetRequestTimeout(timeout time.Duration) {
	s.backend.Client.Timeout = timeout
}
func (s *Collector) Init() {
	s.wg = &sync.WaitGroup{}
	s.lock = &sync.RWMutex{}
	s.backend = &httpBackend{}
}
func (s *Collector) OnRequest(f RequestCallBack) {
	s.lock.Lock()
	if s.requestCallBacks == nil {
		s.requestCallBacks = make([]RequestCallBack, 0, 4)
	}
	s.requestCallBacks = append(s.requestCallBacks, f)
	s.lock.Unlock()
}

func (s *Collector) handleOnRequest(r *Request) {
	for _, f := range s.requestCallBacks {
		f(r)
	}
}

func (s *Collector) OnBlockStorage(f BlockCallBack) {
	s.lock.Lock()
	if s.blockCallBacks == nil {
		s.blockCallBacks = make([]BlockCallBack, 0, 4)
	}
	s.blockCallBacks = append(s.blockCallBacks, f)
	s.lock.Unlock()
}

func (s *Collector) handleOnBlockStorage(r *Response, m *entity.RetBlockMessage) {
	for _, f := range s.blockCallBacks {
		f(r, m)
	}
}

func (s *Collector) OnTransactionStorage(f TransactionCallBack) {
	s.lock.Lock()
	if s.transactionCallBacks == nil {
		s.transactionCallBacks = make([]TransactionCallBack, 0, 4)
	}
	s.transactionCallBacks = append(s.transactionCallBacks, f)
	s.lock.Unlock()
}

func (s *Collector) handleOnTransactionStorage(r *Response, m *entity.RetTransactionMessage) {
	for _, f := range s.transactionCallBacks {
		f(r, m)
	}
}

func (s *Collector) OnBlock(j JsonCallBack) {
	s.lock.Lock()
	if s.jsonCallBacks == nil {
		s.jsonCallBacks = make([]JsonCallBack, 0, 4)
	}
	s.jsonCallBacks = append(s.jsonCallBacks, j)
	s.lock.Unlock()
}

func (s *Collector) handleOnBlock(r *Response, m *entity.RetBlockMessage) {
	for _, f := range s.jsonCallBacks {
		f(r, m)
	}
}
func (s *Collector) OnError(f ErrorCallback) {
	s.lock.Lock()
	if s.errorCallbacks == nil {
		s.errorCallbacks = make([]ErrorCallback, 0, 4)
	}
	s.errorCallbacks = append(s.errorCallbacks, f)
	s.lock.Unlock()
}

func (s *Collector) handleOnError(response *Response, request *Request, err error) error {
	if err == nil && response.StatusCode < 203 {
		return nil
	}
	if err == nil && response.StatusCode >= 203 {
		err = errors.New(http.StatusText(response.StatusCode))
	}

	if response == nil {
		response = &Response{
			Request: request,
		}
	}
	for _, f := range s.errorCallbacks {
		f(response, err)
	}
	return err
}

func (s *Collector) String() string {
	return fmt.Sprintf(
		"Callbacks Status: OnRequest: %d,OnBlockStorage: %d,OnTransactionStorage: %d,OnBlock: %d,OnError: %d",
		len(s.requestCallBacks), len(s.blockCallBacks), len(s.transactionCallBacks), len(s.jsonCallBacks), len(s.errorCallbacks),
	)
}

func (s *Collector) Wait() {
	s.wg.Wait()
}

func (s *Collector) SetLoadBalance(urls string) {
	s.loadIndex = -1
	s.loadBalance = strings.Split(urls, ",")
}

func (s *Collector) GetLoadBalance() string {
	s.lock.Lock()
	s.loadIndex = (s.loadIndex + 1) % len(s.loadBalance)
	s.lock.Unlock()
	return s.loadBalance[s.loadIndex]
}

func (s *Collector) GetHttpBackEnd() *httpBackend {
	return s.backend
}

func (s *Collector) Visit(request *Request) {
	s.wg.Add(1)
	go s.fetch(request)
}

func (s *Collector) fetch(request *Request) error {
	defer s.wg.Done()
	response, err := s.backend.Cache(request)
	if err := s.handleOnError(response, request, err); err != nil {
		log.Errorf("response %s", err.Error())
		return err
	}
	// blockinfo
	if strings.Contains(response.Request.Url, common.URL_BlOCK_INFO) {
		blockMessage := &entity.RetBlockMessage{}
		err := json.Unmarshal(response.Body, &blockMessage)
		if err := s.handleOnError(response, request, err); err != nil {
			log.Errorf("%s Error in JSON unmarshalling from json marshalled object:%s", response.Request.Url, err)
			return err
		}
		s.handleOnBlockStorage(response, blockMessage)
		s.handleOnBlock(response, blockMessage)
	}
	if strings.Contains(response.Request.Url, common.URL_TRANS_INFO) {
		transaction := &entity.RetTransactionMessage{}
		err := json.Unmarshal(response.Body, &transaction)
		if err := s.handleOnError(response, request, err); err != nil {
			log.Errorf("%s Error in JSON unmarshalling from json marshalled object:%s", response.Request.Url, err)
			return err
		}
		s.handleOnTransactionStorage(response, transaction)
	}
	return err
}
