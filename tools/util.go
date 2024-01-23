package tools

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"example.cn/grabblock/common"
	"example.cn/grabblock/conf"
	"example.cn/grabblock/tools/log"
	"golang.org/x/net/proxy"
)

func HexStrToInt64(val string) int64 {
	if val == "" {
		return 0
	}
	val = strings.ReplaceAll(val, "0x", "")
	n, err := strconv.ParseInt(val, 16, 64)
	fmt.Println(err)
	if err != nil {
		return 0
	}

	return n
}
func HexToBigInt(hex string) *big.Int {
	n := new(big.Int)
	hex = strings.ReplaceAll(hex, "0x", "")
	n, _ = n.SetString(hex, 16)
	return n
}

func Http(url string) ([]byte, error) {
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
	res, err := client.Get(url + common.URL_BLOCK_STATE)
	if err != nil {
		log.Errorf("get maxHeight err: %s", err.Error())
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func Copy(dst, src interface{}) error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(src); err != nil {
		return err
	}
	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)
}
