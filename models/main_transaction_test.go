package models

import (
	"encoding/hex"
	"fmt"
	"testing"

	"example.cn/grabblock/entity"
	"github.com/golang/protobuf/proto"
)

var trans = `{
    "retCode": 1,
    "transaction": {
        "hash": "0xb58c2f95bb62bdb3020e0d0df4d89db896ccacb807830b0c667866f8468a790ddc",
        "body": {
            "nonce": 1,
            "address": "CVN07d2697074d522a2903be3e3cd4478f183040c5d",
            "outputs": [
                {
                    "address": "CVNabf95e24ab7a6790216ad1d06179a850f548f691",
                    "amount": "0x021e19e0c9bab2400000"
                },
                {
                    "address": "CVN1fd525a8de1005e01936c9d0a4df7c81208a980c",
                    "amount": "0x021e19e0c9bab2400000"
                },
                {
                    "address": "CVN67e86d2f6c7084b99f0d305694d5259ee9e81973",
                    "amount": "0x021e19e0c9bab2400000"
                }
            ],
            "fee_hi": 0,
            "fee_low": 0,
            "inner_codetype": 0,
            "timestamp": 1602127272588
        },
        "signature": "0x9ab8014286fc432481dd00d6b236cf3228524fba0277c15ee4bc9583212853912ae52c7ae5ae9aa9b460a4b596bcd854d02c01500f2aa994ed27e5ad8330cee507d2697074d522a2903be3e3cd4478f183040c5d636f04a3342e781495c6540bcf68f6d385ea341231baed319f6dbcfb6592ad1be5d8c54c3aa428ed8875d58b8d8cad83230b72400f37d766ad5465de6a23a22a",
        "status": {
            "status": "0x01",
            "hash": "0xd70a2d184276bd6f6a211982ed44e7f2b182b3c07def570fc72a715e86ef929f",
            "height": 820,
            "timestamp": 1602127275086
        },
        "node": {
            "nid": "V0sOAmJWA0CwthPzo4hEcvA2i8r7x",
            "address": "CVN5d9cdda85093d68c28573ae9875eb32dbad6f0e0"
        },
        "accepttimestamp": 1602127273525
    }
}`

func TestAddMainTransaction(t *testing.T) {
	//err := conf.ParseConf("../config.json")
	//if err != nil {
	//	if os.IsNotExist(err) {
	//		t.Error("please config config.json")
	//		os.Exit(0)
	//	}
	//	t.Error(err)
	//}
	//log.InitLogger()
	//Setup()
	//transaction := &entity.RetTransactionMessage{}
	//
	//if err := json.Unmarshal([]byte(trans), &transaction); err != nil {
	//	t.Error(err)
	//}

	//AddMainTransaction(transaction.Transaction)

	//0x08031a14ff8a88c5c4701f4308fab0b26e58e54fd753eb81221431303030303030303030303030303030303030306012

	rc20 := &entity.ContractRC20{}
	v, _ := hex.DecodeString("08031a14ff8a88c5c4701f4308fab0b26e58e54fd753eb81221431303030303030303030303030303030303030306012")
	proto.Unmarshal(v, rc20)
	fmt.Printf("rc= %v", rc20)
}
