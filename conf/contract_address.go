package conf

import (
	"encoding/json"
	"os"
)

/**
    创世块儿地址（启动会读取json配置文件，读取到内存）
 */

var ContractAddressSet *ContractAddressconfig

type ContractAddressconfig struct {
	ContractAddress ContractAddress `json:"contractAddress"`
	// 合约地址-改为参数配置
	Mode     string   `json:"mode"`
	Name     string   `json:"name"`
}

// 合约信息
type ContractAddress struct {
	ContractAdminAddress             string     `json:"contractAdminAddress"`
	ContractCommitteeAddress         string     `json:"contractCommitteeAddress"`
	ContractDataStorageAddress       string     `json:"contractDataStorageAddress"`
	ContractBussinessAddress         string     `json:contractBussinessAddress`
}

// 解析参数（json）
func ParseContractAddressConf(config string) error {
	var c ContractAddressconfig
	conf, err := os.Open(config)
	if err != nil {
		return err
	}
	err = json.NewDecoder(conf).Decode(&c)
	ContractAddressSet = &c
	return err
}
