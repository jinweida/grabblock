package spider

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"

	"example.cn/grabblock/conf"
	"example.cn/grabblock/tools/log"
	"golang.org/x/net/proxy"
)

type httpBackend struct {
	Client *http.Client
	lock   *sync.RWMutex
}

func (h *httpBackend) Init() {
	rand.Seed(time.Now().UnixNano())
	h.Client = &http.Client{
		Timeout: 10 * time.Second,
	}
	h.lock = &sync.RWMutex{}
}
func (h *httpBackend) Cache(req *Request) (*Response, error) {
	client := &http.Client{
		Timeout: conf.Context.Node.Interval * time.Second,
	}
	if len(conf.Context.Node.Proxy) > 0 {
		// create a socks5 dialer
		dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
		if err != nil {
			fmt.Fprintln(os.Stderr, "can't connect to the proxy:", err)
			os.Exit(1)
		}
		// setup a http client
		httpTransport := &http.Transport{
			Dial: dialer.Dial,
		}
		client.Transport = httpTransport
	}
	res, err := client.Get(req.Url)
	if err != nil {
		log.Errorf("cache host: %s ,err : %s", req.Url, err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("cache host: %s ,readall : %s", req.Url, err.Error())
		return nil, err
	}

	return &Response{
		StatusCode: res.StatusCode,
		Body:       body,
		Request:    req,
	}, nil
}
