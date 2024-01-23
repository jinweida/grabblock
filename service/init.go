package service

import (
	"os"
	"strings"

	constant "example.cn/grabblock/common"
	"example.cn/grabblock/conf"
	"example.cn/grabblock/tools/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

var FileType map[string]int32        //文件类型map
var ContractType map[string]Contract // 声明一个hashmap，还不能直接使用，必须使用make来初始化

// 合约处理--存储每一个合约处理对象（与ContractType组成命令模式）
type Contract struct {
	Name      string
	Abi       abi.ABI
	ClassType interface{}
	Method    string
	Sig       string
}

/*
*

	初始化各个合约处理对象。为后续抓块后（获得交易），解析合约信息准备
*/
func init() {
	err := conf.ParseContractAddressConf("contractAddress.json")
	if err != nil {
		if os.IsNotExist(err) {
			log.Error("please config contractAddress.json")
			os.Exit(0)
		}
		log.Panicf("err:%s", err)
	}
	ContractType = make(map[string]Contract, 0) // 初始化一个map
	FileType = make(map[string]int32, 0)
	//读取abi（合约说明）
	admin, err := abi.JSON(strings.NewReader(constant.AdminGroup))
	if err != nil {
		log.Errorf("AdminGroup is fail %s", err)
	}
	for method, _ := range admin.Methods {
		hash := crypto.Keccak256Hash([]byte(admin.Methods[method].Sig))
		sig := hash.Hex()[:10]
		//组装map（key=合约地址+方法签名，value=合约处理方法）
		ContractType[conf.ContractAddressSet.ContractAddress.ContractAdminAddress+sig] = Contract{
			Abi: admin, Method: admin.Methods[method].Name, Sig: sig, ClassType: NewAdminGroup(),
		}
	}
	committee, err := abi.JSON(strings.NewReader(constant.CommitteeGroup))
	if err != nil {
		log.Errorf("CommitteeGroup is fail %s", err)
	}
	for method, _ := range committee.Methods {
		hash := crypto.Keccak256Hash([]byte(committee.Methods[method].Sig))
		sig := hash.Hex()[:10]
		ContractType[conf.ContractAddressSet.ContractAddress.ContractCommitteeAddress+sig] = Contract{
			Abi: committee, Method: committee.Methods[method].Name, Sig: sig, ClassType: NewCommitteeGroup(),
		}
	}

	dataStorage, err := abi.JSON(strings.NewReader(constant.DataStorageDomain))
	if err != nil {
		log.Errorf("DataStorageDomain is fail %s", err)
	}

	for method, _ := range dataStorage.Methods {
		hash := crypto.Keccak256Hash([]byte(dataStorage.Methods[method].Sig))
		sig := hash.Hex()[:10]
		ContractType[conf.ContractAddressSet.ContractAddress.ContractDataStorageAddress+sig] = Contract{
			Abi: dataStorage, Method: dataStorage.Methods[method].Name, Sig: sig, ClassType: NewDataStorageDomain(),
		}
	}

	bussiness, err := abi.JSON(strings.NewReader(constant.BussinessDomain))
	if err != nil {
		log.Errorf("BussinessDomain is fail %s", err)
	}

	for method, _ := range bussiness.Methods {
		hash := crypto.Keccak256Hash([]byte(bussiness.Methods[method].Sig))
		sig := hash.Hex()[:10]
		ContractType[conf.ContractAddressSet.ContractAddress.ContractBussinessAddress+sig] = Contract{
			Abi: bussiness, Method: bussiness.Methods[method].Name, Sig: sig, ClassType: NewBussinessDomain(),
		}
	}
	//图片
	FileType["dwg"] = constant.FILE_TYPE_PIC
	FileType["dxf"] = constant.FILE_TYPE_PIC
	FileType["gif"] = constant.FILE_TYPE_PIC
	FileType["jp2"] = constant.FILE_TYPE_PIC
	FileType["jpe"] = constant.FILE_TYPE_PIC
	FileType["jpeg"] = constant.FILE_TYPE_PIC
	FileType["jpg"] = constant.FILE_TYPE_PIC
	FileType["png"] = constant.FILE_TYPE_PIC
	FileType["svf"] = constant.FILE_TYPE_PIC
	FileType["tif"] = constant.FILE_TYPE_PIC
	FileType["tiff"] = constant.FILE_TYPE_PIC
	FileType["webp"] = constant.FILE_TYPE_PIC
	FileType["bmp"] = constant.FILE_TYPE_PIC
	FileType["pcx"] = constant.FILE_TYPE_PIC
	FileType["tga"] = constant.FILE_TYPE_PIC
	FileType["exif"] = constant.FILE_TYPE_PIC
	FileType["fpx"] = constant.FILE_TYPE_PIC
	FileType["svg"] = constant.FILE_TYPE_PIC
	FileType["psd"] = constant.FILE_TYPE_PIC
	FileType["cdr"] = constant.FILE_TYPE_PIC
	FileType["ico"] = constant.FILE_TYPE_PIC
	//音频
	FileType["ac3"] = constant.FILE_TYPE_SOUND
	FileType["au"] = constant.FILE_TYPE_SOUND
	FileType["mp2"] = constant.FILE_TYPE_SOUND
	FileType["ogg"] = constant.FILE_TYPE_SOUND
	FileType["flac"] = constant.FILE_TYPE_SOUND
	FileType["ape"] = constant.FILE_TYPE_SOUND
	FileType["wav"] = constant.FILE_TYPE_SOUND
	FileType["mp3"] = constant.FILE_TYPE_SOUND
	FileType["aac"] = constant.FILE_TYPE_SOUND
	FileType["wma"] = constant.FILE_TYPE_SOUND
	//视频
	FileType["3gpp"] = constant.FILE_TYPE_VIDEO
	FileType["mp4"] = constant.FILE_TYPE_VIDEO
	FileType["mpeg"] = constant.FILE_TYPE_VIDEO
	FileType["3gp"] = constant.FILE_TYPE_VIDEO
	FileType["mpg"] = constant.FILE_TYPE_VIDEO
	FileType["wmv"] = constant.FILE_TYPE_VIDEO
	FileType["asf"] = constant.FILE_TYPE_VIDEO
	FileType["asx"] = constant.FILE_TYPE_VIDEO
	FileType["rm"] = constant.FILE_TYPE_VIDEO
	FileType["m4v"] = constant.FILE_TYPE_VIDEO
	FileType["mov"] = constant.FILE_TYPE_VIDEO
	FileType["avi"] = constant.FILE_TYPE_VIDEO
	FileType["dat"] = constant.FILE_TYPE_VIDEO
	FileType["mkv"] = constant.FILE_TYPE_VIDEO
	FileType["flv"] = constant.FILE_TYPE_VIDEO
	FileType["vob"] = constant.FILE_TYPE_VIDEO

	//其它
	FileType["doc"] = constant.FILE_TYPE_OTHER
	FileType["pdf"] = constant.FILE_TYPE_OTHER
	FileType["docx"] = constant.FILE_TYPE_OTHER
	FileType["tif"] = constant.FILE_TYPE_OTHER
	FileType["dwg"] = constant.FILE_TYPE_OTHER
	FileType["psd"] = constant.FILE_TYPE_OTHER
	FileType["rtf"] = constant.FILE_TYPE_OTHER
	FileType["xml"] = constant.FILE_TYPE_OTHER
	FileType["html"] = constant.FILE_TYPE_OTHER
	FileType["eml"] = constant.FILE_TYPE_OTHER
	FileType["xls"] = constant.FILE_TYPE_OTHER
	FileType["mdb"] = constant.FILE_TYPE_OTHER
	FileType["ps"] = constant.FILE_TYPE_OTHER
	FileType["rar"] = constant.FILE_TYPE_OTHER
	FileType["mid"] = constant.FILE_TYPE_OTHER
	FileType["asf"] = constant.FILE_TYPE_OTHER
	FileType["gz"] = constant.FILE_TYPE_OTHER

}
