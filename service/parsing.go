package service

import (
	"errors"
	"fmt"
	"reflect"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/entity"
	"example.cn/grabblock/tools"
	"example.cn/grabblock/tools/log"
)

type Parsing struct{}

func NewParsing() *Parsing {
	parsing := &Parsing{}
	//parsing.Init()
	return parsing
}

// 合约解码
func (c *Parsing) Decode(m *entity.TransactionInfo) (int64, error) {
	if m.Status.Status == constant.TRANSACTION_SUCCESS {
		if m.Body.InnerCodetype == int32(constant.INNER_CODE_TYPE_CVM) {
			if len(m.Body.Outputs) > 0 {
				//合约参数 解码
				contractAddress := m.Body.Outputs[0].Address
				sig := m.Body.CodeData[:10]
				if contratType, ok := ContractType[contractAddress+sig]; ok {
					v := reflect.ValueOf(contratType.ClassType).MethodByName(tools.FirstToUpper(contratType.Method))
					log.Infof("ClassType=[%s] method=[%s]", reflect.ValueOf(contratType.ClassType).String(), contratType.Method)
					if v.IsValid() {
						params := make([]reflect.Value, 2)
						params[0] = reflect.ValueOf(m)
						params[1] = reflect.ValueOf(contratType.Abi)
						rs := v.Call(params)
						if rs[0].IsNil() {
							return constant.PARSING_DONE, nil
						}
						return constant.PARSING_DONE, rs[0].Interface().(error)
					} else {
						log.Errorf("contractAddress:%s,sig:%s,method:%s 不需要解析!", contractAddress, sig, tools.FirstToUpper(contratType.Method))
						return constant.PARSING_NOT, nil
					}
				} else {
					log.Errorf("contractAddress:%s,sig:%s,method:%s 不需要解析!", contractAddress, sig, tools.FirstToUpper(contratType.Method))
					return constant.PARSING_NOT, nil
				}
			}
		}
	}
	return constant.PARSING_STATUS_FAIL, errors.New(fmt.Sprintf("status is not %s", m.Status.Status))
}
