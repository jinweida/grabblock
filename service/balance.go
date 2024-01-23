package service

import (
	"encoding/json"
	"fmt"
	"time"

	"example.cn/grabblock/common"

	//"example.cn/grabblock/common"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/models"
	"example.cn/grabblock/spider"
	"example.cn/grabblock/tools"
	"example.cn/grabblock/tools/log"

	//"fmt"
	"strings"
)

var (
	Queue  *tools.DataContainer
	Coller *spider.Collector
)

func NewBalance(coller *spider.Collector) {
	Queue = tools.NewDataContainer(10)
	Coller = coller
}

var MAX_WAIT_TIME = 10 * time.Millisecond

// 入队列
func PushTrans(trans []*models.MainTransaction) {
	for _, tran := range trans {
		Queue.Push(tran, MAX_WAIT_TIME)
	}
}
func FethBalance() {
	for {
		v := Queue.Pop(MAX_WAIT_TIME)
		if v == nil {
			continue
		}
		tran := v.(*models.MainTransaction)
		if tran.Status == common.TRANSACTION_SUCCESS {
			list := make([]*models.MainAccountBalance, 0)
			if tran.InnerCodetype == common.INNER_CODE_TYPE_RC20 {
				if len(tran.ToAddress) > 0 && len(tran.CtAddress) > 0 {
					result := getTokenBalance(fmt.Sprintf(Coller.GetLoadBalance()+"%s?bd={\"token_address\":\"%s\",\"owner_address\":\"%s\"}", common.URL_TOKEN_BALANCE, tran.ToAddress, tran.CtAddress))

					if result.RetCode == common.RET_SUCCESS_CODE {
						list = append(list, &models.MainAccountBalance{
							Address:      tran.CtAddress,
							Balance:      tools.HexToBigInt(string(result.Value.Balance)).String(),
							TokenSymbol:  result.Symbol,
							TokenAddress: tran.ToAddress,
							AccountType:  2,
						})
					}
					if tran.SubCodeType == common.INNER_CODE_TYPE_RC20_TRANSFER {
						result = getTokenBalance(fmt.Sprintf(Coller.GetLoadBalance()+"%s?bd={\"token_address\":\"%s\",\"owner_address\":\"%s\"}", common.URL_TOKEN_BALANCE, tran.ToAddress, tran.FromAddress))

						if result.RetCode == common.RET_SUCCESS_CODE {
							list = append(list, &models.MainAccountBalance{
								Address:      tran.FromAddress,
								Balance:      tools.HexToBigInt(string(result.Value.Balance)).String(),
								TokenSymbol:  result.Symbol,
								TokenAddress: tran.ToAddress,
								AccountType:  2,
							})
						}
					}
				}
			}
			if tran.InnerCodetype == common.INNER_CODE_TYPE_TRANSFER {

				result := getTransferBalance(fmt.Sprintf(Coller.GetLoadBalance()+"%s?bd={\"address\":\"%s\"}", common.URL_TRANSFER_BALANCE, FormatAddress(tran.FromAddress)))
				if result.RetCode == common.RET_SUCCESS_CODE {
					list = append(list, &models.MainAccountBalance{
						Address:     tran.FromAddress,
						Balance:     tools.HexToBigInt(result.Balance).String(),
						AccountType: 1,
					})
				}
				result = getTransferBalance(fmt.Sprintf(Coller.GetLoadBalance()+"%s?bd={\"address\":\"%s\"}", common.URL_TRANSFER_BALANCE, FormatAddress(tran.ToAddress)))
				if result.RetCode == common.RET_SUCCESS_CODE {
					list = append(list, &models.MainAccountBalance{
						Address:     tran.ToAddress,
						Balance:     tools.HexToBigInt(result.Balance).String(),
						AccountType: 1,
					})
				}
			}
			if len(list) > 0 {

				if err := models.AddMainAccountBalance(list); err != nil {
					log.Errorf("account balance update error %s,%+v", err.Error(), tran)
					//再次插入队列
					//Queue.Push(tran)
				}
			}
		}
	}
}

func getTokenBalance(url string) *entity.RespQueryRC20TokenValue {
	body, err := tools.Http(url)
	if err != nil {
		log.Errorf("balance Http token error:%s", err.Error())
	}
	result := &entity.RespQueryRC20TokenValue{}
	err = json.Unmarshal(body, result)
	if err != nil {
		log.Errorf("balance Unmarshal token error : %s", err.Error())
	}
	return result
}
func getTransferBalance(url string) *entity.RetAccountMessage {
	body, err := tools.Http(url)
	if err != nil {
		log.Errorf("balance Http error:%s", err.Error())
	}
	result := &entity.RetAccountMessage{}
	err = json.Unmarshal(body, &result)

	if err != nil {
		log.Errorf("balance Unmarshal error : %s = %s", err.Error(), url)
	}
	return result
}

func FormatAddress(address string) string {
	return strings.ReplaceAll(address, "0x", "")
}
