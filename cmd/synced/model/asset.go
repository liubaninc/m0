package model

type Asset struct {
	ID           uint   `json:"-" gorm:"primarykey" `
	Name         string `json:"denom" gorm:"unique"`
	Initiator    string `json:"initiator"`
	RefTxID      string `json:"ref_tx"`
	RefMsgOffset int32  `json:"ref_msg"`
	Time         string `json:"time"`
	MintAmount   string `json:"mint_amount"`
	BurnAmount   string `json:"burn_amount"`
}
