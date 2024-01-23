package tools

import (
	"encoding/hex"
	"strings"
	"unicode"
)

func BytesToString(info []byte) string {
	s := strings.ReplaceAll(string(info), "0x", "")
	if len(s)==0{
		return ""
	}
	v, _ := hex.DecodeString(s)
	return string(v[2:])
}
func FirstToUpper(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}
