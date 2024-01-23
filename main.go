package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"example.cn/grabblock/common"
	"example.cn/grabblock/conf"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/models"
	"example.cn/grabblock/service"
	"example.cn/grabblock/spider"
	"example.cn/grabblock/tools"
	"example.cn/grabblock/tools/log"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

var on, height int64

func init() {
	flag.Parse()
	//defer log.Sync()
	err := conf.ParseConf("config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("please config config.json")
			os.Exit(0)
		}
		log.Panicf("err:%s", err)
	}
	models.Setup()
	//33608212
	height = models.GetBlockMaxheight().Height
}

// https://gobyexample.com/
func main() {
	log.InitLogger()
	s := spider.NewCollector()
	service.NewBalance(s)
	s.SetLoadBalance(conf.Context.Node.Url)
	s.OnBlockStorage(func(r *spider.Response, m *entity.RetBlockMessage) {
		//log.Infof("common.URL_BlOCK_INFO storage %s", r.Request.Url)
		for _, blockInfo := range m.Block {
			models.AddMainBlock(blockInfo)
		}
	})
	s.OnTransactionStorage(func(r *spider.Response, m *entity.RetTransactionMessage) {
		log.Infof("common.URL_TRANS_INFO storage %s %d", r.Request.Url, m.RetCode)

		if m.RetCode == common.RET_SUCCESS_CODE {

			m.Transaction.Status = m.Status
			m.Transaction.Status.Hash = r.Request.BlockHash
			if m.Transaction.Body.InnerCodetype == common.INNER_CODE_TYPE_FILE {
				if !conf.Context.Node.IsEvfs {
					return
				}
			}
			accountTrans, err := models.AddMainTransaction(m.Transaction)
			if err != nil {
				log.Errorf("main_transaction add data fail:%s", err.Error())
			} else {
				service.PushTrans(accountTrans)
			}
		}
	})
	s.OnBlock(func(r *spider.Response, m *entity.RetBlockMessage) {
		if m.RetCode != int32(common.RET_SUCCESS_CODE) {
			log.Infof("%s retcode %d", r.Request.Url, m.RetCode)
			return
		}
		for _, block := range m.GetBlock() {
			if len(block.GetHeader().TxHashs) > 0 {
				for _, hash := range block.Header.TxHashs {
					request := &spider.Request{
						Url:       fmt.Sprintf(s.GetLoadBalance()+"%s?bd={\"hash\":\"%s\",\"height\":\"%d\"}", common.URL_TRANS_INFO, hash[2:], block.Header.Height),
						Hash:      hash[2:],
						BlockHash: block.GetHeader().Hash,
					}
					s.Visit(request)
				}
			}
		}
	})
	// catch error
	s.OnError(func(r *spider.Response, err error) {
		if r.Request.FailCount > conf.Context.Node.FailCount {
			if err := models.NewMainErrorUrl().Add(r); err != nil {
				log.Errorf("main_error_url add data fail:%s", err.Error())
			}
		} else {
			r.Request.FailCount = r.Request.FailCount + 1
			log.Infof("url:%s,failcount:%d", r.Request.Url, (r.Request.FailCount))
			s.Visit(r.Request)
		}
	})
	//start grabblock
	go func() {
		// 创建一个计时器
		for {
			if on == common.GRABBLOCK_OFF {
				log.Infof("grabblock stop; %s", s.String())
				time.Sleep(time.Second * conf.Context.Node.Interval)
			} else {
				start := time.Now()
				maxHeight := height + conf.Context.Node.Count
				blockHeight, err := BlockState(s.GetLoadBalance())
				//control maxheight
				if err != nil {
					log.Errorf("get blockstate err %s", err.Error())
					time.Sleep(time.Second * conf.Context.Node.Interval)
				} else {
					safeHeight := blockHeight - conf.Context.Node.SafeDistance
					if safeHeight < 0 {
						safeHeight = 0
					}
					if safeHeight < maxHeight {
						maxHeight = safeHeight
					}
					//创世块http参数与普通块儿不一样
					if height == 1 {
						var chuangshi int64 = 0
						url := fmt.Sprintf(s.GetLoadBalance()+"%s?bd={\"height\":%d,\"type\":%d}", common.URL_BlOCK_INFO, 0, 1)
						request := &spider.Request{
							Url:    url,
							Height: chuangshi,
						}
						s.Visit(request)
					}
					s.Wait()
					//TODO 处理缺失的块
					//普通块儿处理
					for i := height; i < maxHeight; i++ {
						url := fmt.Sprintf(s.GetLoadBalance()+"%s?bd={\"height\":\"%d\"}", common.URL_BlOCK_INFO, i)
						request := &spider.Request{
							Url:    url,
							Height: i,
						}
						s.Visit(request)
					}
					s.Wait()
					tc := time.Since(start) //计算耗时
					log.Infof("complete=[%d~%d];safeHeight=[%d];blockheight=[%d];threads=[%d];timecost=%v", height, maxHeight, safeHeight, blockHeight, maxHeight-height, tc)

					if (safeHeight - maxHeight) == 0 {
						log.Infof("安全块延迟[%v]", time.Second*conf.Context.Node.Interval)
						time.Sleep(time.Second * conf.Context.Node.Interval)
					}
					height = maxHeight
				}
			}
		}
	}()
	//是否解析合约
	if conf.Context.Node.Contractparsing {
		startup := service.NewStartup()
		newStatic := service.NewStatic()
		c := cron.New(cron.WithSeconds())
		//解析合约
		c.AddFunc(fmt.Sprintf("*/%d * * * * *", conf.Context.Node.Interval), startup.ParsingContract)
		//解析文件
		c.AddFunc(fmt.Sprintf("*/%d * * * * *", conf.Context.Node.Interval), startup.ParsingFile)
		//统计业务合约调用次数
		c.AddFunc(fmt.Sprintf("*/%d * * * * *", conf.Context.Node.Interval), newStatic.StaticBizContractCallTimes)
		//统计业务域存储信息
		c.AddFunc(fmt.Sprintf("*/%d * * * * *", conf.Context.Node.Interval), newStatic.StaticBiz)
		//统计业务域下系统存储信息
		c.AddFunc(fmt.Sprintf("*/%d * * * * *", conf.Context.Node.Interval), newStatic.StaticBizSys)
		c.Start()
		defer c.Stop()
	} else {
		// go service.FethBalance()
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.GET("/status", func(c *gin.Context) {
		flag := c.DefaultQuery("flag", "0")
		msg := "停止追块"
		if flag == strconv.Itoa(common.GRABBLOCK_OFF) {
			msg = "停止追块"
			on = common.GRABBLOCK_OFF
			log.Info(fmt.Sprintf("ononon=%d", on))
		}
		if flag == strconv.Itoa(common.GRABBLOCK_ON) {
			msg = "开启追块"
			on = common.GRABBLOCK_ON
		}
		c.JSON(200, gin.H{
			"code": "1",
			"msg":  msg,
		})
	})
	router.Run(fmt.Sprintf(":%d", conf.Context.Server.Port))
}
func BlockState(url string) (int64, error) {
	body, err := tools.Http(url)
	if err != nil {
		return 0, err
	}
	state := &entity.RetChainSummaryMessage{}
	err = json.Unmarshal(body, &state)
	if err != nil {
		log.Errorf("unmarshal blockstate : %s", err.Error())
		return 0, err
	}
	return state.Last.Height, nil
}
