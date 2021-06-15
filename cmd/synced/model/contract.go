package model

type Contract struct {
	ID           uint   `json:"-" gorm:"primarykey" `
	Name         string `json:"name" gorm:"unique" `
	Initiator    string `json:"initiator"`
	Number       uint64 `json:"number"`
	Runtime      string `json:"runtime"`
	Compiler     string `json:"compiler"`
	Digest       string `json:"digest"`
	VmCompiler   string `json:"vm_compiler"`
	ContractType string `json:"contract_type"`
	Version      string `json:"version"`
	Time         string `json:"time"`
	Status       int    `json:"status"`
	Deploy       int64  `json:"num"`
	Invoke       int64  `json:"total"`
}

type ContractUpgrade struct {
	ID      uint   `json:"-" gorm:"primarykey" `
	Name    string `json:"name" `
	Version string `json:"version"`
	Hash    string `json:"hash" `
	Time    string `json:"time" `
}

type ContractInvoke struct {
	ID      uint   `json:"-" gorm:"primarykey" `
	Name    string `json:"name" `
	Version string `json:"version"`
	Hash    string `json:"hash" `
	Time    string `json:"time" `
	Size    int    `json:"size"`
	Height  int64  `json:"height"`
}
