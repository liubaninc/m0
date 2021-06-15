package model

type Block struct {
	ID       uint           `json:"-" gorm:"primarykey" `
	Height   int64          `json:"height" gorm:"unique"`
	Size     int            `json:"size"`
	Hash     string         `json:"hash" gorm:"unique"`
	Time     string         `json:"time"`
	Proposer string         `json:"proposer"`
	PrevHash string         `json:"prev_hash"`
	TxNum    int            `json:"tx_num"`
	TxList   []*Transaction `json:"txs" gorm:"foreignKey:Height;references:Height"`
}
