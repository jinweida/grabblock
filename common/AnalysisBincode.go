package common

import (
	"encoding/hex"
	"strconv"
	"strings"

	"example.cn/grabblock/tools/log"
)

/**
  此函数是解析创世块儿使用
*/

var moban []byte = []byte{0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x6a, 0x6b, 0x6c, 0x6d, 0x6e, 0x6f,
	0x70, 0x71, 0x72, 0x73, 0x74, 0x75, 0x76, 0x77, 0x78, 0x79, 0x7a, 0x7b, 0x7c, 0x7d, 0x7e, 0x7f}

// 魔术地址（用于区分bincode与参数）
var magicEndStr string = "83398101"

func GetParamsHexStr(code string) string {
	replace_result := strings.Replace(code, "0x", "", 5)
	len, _ := calcByteLength(replace_result)
	if len > 0 {
		return replace_result[len*2:]
	}
	return ""
}

func calcByteLength(code string) (int64, error) {
	replace_result := strings.Replace(code, "0x", "", 5)
	end := strings.Index(replace_result, magicEndStr)
	if end > -1 {
		replace_result = replace_result[0:end]
	}
	data, _ := Hextobyte(replace_result)
	start := 0
	for i := len(data) - 2; i > -1; i-- {
		if data[i] == moban[start] {
			hexlen := replace_result[i*2+2:]
			log.Infof("hexLen:[%s]", hexlen)
			return strconv.ParseInt(hexlen, 16, 64)
		}
		start++
	}
	return 0, nil
}

func Hextobyte(str string) ([]byte, error) {
	if len(str) != 0 {
		if len(str)%2 == 1 {
			str = "0" + str
		}
	}
	log.Info("[str]---" + str)
	return hex.DecodeString(str)
}
