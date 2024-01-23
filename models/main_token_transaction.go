package models

import "math/big"

// trans
type MainTokenTransaction struct {
	TransactionId       string   `gorm:"column:transaction_id;type:varchar(100);primary_key;"`
	TransactionHash     string   `gorm:"column:transaction_hash;type:varchar(100);"`
	TransactionIndex    int      `gorm:"column:transaction_index"`
	TransactionOutCount int      `gorm:"column:transaction_out_count"`
	BlockHeight         int64    `gorm:"column:block_height;type:bigint"`
	BlockHash           string   `gorm:"column:block_hash;type:varchar(100);"`
	PeerId              string   `gorm:"column:peer_id;type:varchar(100);"`
	Nonce               int32    `gorm:"column:nonce"`
	Function            int      `gorm:"column:function"` //2=create 3=transfer
	ExtData             string   `gorm:"column:ext_data;type:varchar(2000)"`
	TokenAddress        string   `gorm:"column:token_address;type:varchar(100)"`
	FromAddress         string   `gorm:"column:from_address;type:varchar(100)"`
	ToAddress           string   `gorm:"column:to_address;type:varchar(100)"`
	Amount              *big.Int `gorm:"column:amount;DECIMAL(32,0)"`
	TotalAmount         *big.Int `gorm:"column:total_amount;"`
	Status              string   `gorm:"column:status;type:varchar(10)"`
	StatusResult        string   `gorm:"column:status_result;type:varchar(100)"`
	StatusHash          string   `gorm:"column:status_hash;type:varchar(100)"`
	StatusTime          int64    `gorm:"column:status_time;"`
	Timestamp           int64    `gorm:"column:timestamp;"`
	AcceptTime          int64    `gorm:"column:accept_time;"`
	NodeAddress         string   `gorm:"column:node_address;type:varchar(100)"`
	Model
}
