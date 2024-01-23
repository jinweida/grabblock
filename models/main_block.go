package models

import (
	"fclink.cn/grabblock/entity"
)

// blocks
type MainBlock struct {
	BlockId              int64  `gorm:"AUTO_INCREMENT;column:block_id;type:bigint;primary_key"`
	BlockHash            string `gorm:"column:block_hash;type:varchar(100)"`
	ParentBlockHash      string `gorm:"column:parent_block_hash;type:varchar(100)"`
	MinnerAccountAddress string `gorm:"column:minner_account_address;type:varchar(100)"`
	Height               int64  `gorm:"column:height;type:bigint"`
	BlockSize            int64  `gorm:"column:block_size;type:bigint"`
	StateRoot            string `gorm:"column:state_root;type:varchar(100)"`
	ReceiptRoot          string `gorm:"column:receipt_root;type:varchar(100)"`
	Reward               string `gorm:"column:reward;type:varchar(50)"` //奖励金额
	ExtData              string `gorm:"column:ext_data;type:varchar(2000)"`
	Timestamp            int64  `gorm:"column:timestamp;"`
	TransactionCount     int    `gorm:"column:transaction_count;type:int"`
	ScanCount            int64  `gorm:"column:scan_count;type:int"` //处理完的交易数
	Model
}

func AddMainBlock(blockInfo *entity.BlockInfo) error {
	db.Create(&MainBlock{
		BlockHash:            blockInfo.Header.Hash,
		ParentBlockHash:      blockInfo.Header.ParentHash,
		MinnerAccountAddress: blockInfo.Miner.Address,
		Height:               blockInfo.Header.Height,
		BlockSize:            blockInfo.Header.Totaltriesize,
		StateRoot:            blockInfo.Header.StateRoot,
		ReceiptRoot:          blockInfo.Header.ReceiptRoot,
		ExtData:              blockInfo.Header.ExtraData,
		Reward:               blockInfo.Miner.Reward,
		Timestamp:            blockInfo.Header.Timestamp / 1000,
		TransactionCount:     len(blockInfo.Header.TxHashs),
	})

	return db.Error
}

func GetBlockMaxheight() *MainBlock {
	mainblock := &MainBlock{}
	db.Select("MAX( height ) AS height").Take(&mainblock)
	return mainblock
}
