package model

type BlockChain struct {
	ID           uint  `json:"-" gorm:"primarykey" `
	BlockNum     int64 `json:"block_num"`
	TxNum        int64 `json:"tx_num"`
	MsgNum       int64 `json:"msg_num"`
	AssetNum     int64 `json:"asset_num"`
	ContractNum  int64 `json:"contract_num"`
	ValidatorNum int64 `json:"validator_num"`
	PeerNum      int64 `json:"peer_num"`
}
