package model

type MsgUTXO struct {
	ID                 uint             `json:"-" gorm:"primarykey" `
	Inputs             []*MsgInput      `json:"inputs,omitempty" `
	Outputs            []*MsgOutput     `json:"outputs,omitempty" `
	ContractRequests   []*InvokeRequest `json:"requests,omitempty" `
	InputNum           int              `json:"input_num,omitempty" `
	OutputNum          int              `json:"output_num,omitempty" `
	ContractRequestNum int              `json:"request_num,omitempty" `
	Type               string           `json:"type" `
	Desc               string           `json:"desc" `
	AuthRequire        string           `json:"auth,omitempty" `
	Assets             string           `json:"-" `
	Contracts          string           `json:"-" `
	Addresses          string           `json:"-" `
	TransactionID      uint             `json:"-" gorm:"index" `
}

type MsgInput struct {
	ID           uint   `json:"-" gorm:"primarykey" `
	RefTxID      string `json:"ref_tx_id" `
	RefMsgOffset int32  `json:"ref_msg_offset" `
	RefOffset    int32  `json:"ref_offset" `
	Address      string `json:"address" `
	Amount       string `json:"amount" `
	FrozenHeight int64  `json:"frozen_height,omitempty" `
	MsgUTXOID    uint   `json:"-" gorm:"index" `
}

type MsgOutput struct {
	ID           uint   `json:"-" gorm:"primarykey" `
	Address      string `json:"address" `
	Amount       string `json:"amount" `
	FrozenHeight int64  `json:"frozen_height,omitempty" `
	MsgUTXOID    uint   `json:"-" gorm:"index" `
}

type InvokeRequest struct {
	ID           uint   `json:"-" gorm:"primarykey" `
	ModuleName   string `json:"module_name" `
	ContractName string `json:"contract_name" `
	MethodName   string `json:"method_name" `
	Args         string `json:"args" `
	Amount       string `json:"amount" `
	MsgUTXOID    uint   `json:"-" gorm:"index" `
}
