package model

type Validator struct {
	ID     uint   `json:"-" gorm:"primarykey" `
	Name   string `json:"name"`
	PubKey string `json:"pub_key"`
	Pow    int64  `json:"pow"`
	NodeID string `json:"node_id"`
}
