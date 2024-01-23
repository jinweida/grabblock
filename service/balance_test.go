package service

import (
	"fclink.cn/grabblock/conf"
	"fclink.cn/grabblock/models"
	"fclink.cn/grabblock/spider"
	"fclink.cn/grabblock/tools/log"
	"os"
	"testing"
)

func TestNewBalance(t *testing.T) {}
func TestPushTrans(t *testing.T)  {}
func TestFethBalance(t *testing.T) {
	log.InitLogger()
	err := conf.ParseConf("../config.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("please config config.json")
			os.Exit(0)
		}
		log.Panic(err)
	}
	models.Setup()
	s := spider.NewCollector()
	s.SetLoadBalance(conf.Context.Node.Url)
	NewBalance(s)
	accountTrans := make([]*models.MainTransaction, 0)
	accountTrans = append(accountTrans, &models.MainTransaction{
		Status:        "0x01",
		InnerCodetype: 0,
		SubCodeType:   3,
		FromAddress:   "07d2697074d522a2903be3e3cd4478f183040c5d",
		ToAddress:     "CVNabf95e24ab7a6790216ad1d06179a850f548f691",
		CtAddress:     "0x3c1ea4aa4974d92e0eabd5d024772af3762720a0",
	})
	PushTrans(accountTrans)
	FethBalance()

	//fmt.Println(tools.HexToBigInt("0x845748f67cc37f7a00000"))
}
