package model

type BlockChainChart struct {
	ID          uint   `json:"-" gorm:"primarykey" `
	Time        string `json:"time" gorm:"unique"`
	BlockNum    int    `json:"block_num"`
	TxNum       int    `json:"tx_num"`
	MsgNum      int    `json:"msg_num"`
	AssetNum    int    `json:"asset_num"`
	ContractNum int    `json:"contract_num"`
}
type BlockChainTPSChart struct {
	ID     uint   `json:"-" gorm:"primarykey" `
	Time   string `json:"time" gorm:"unique"`
	TxNum  int    `json:"tx_num"`
	MsgNum int    `json:"msg_num"`
}
