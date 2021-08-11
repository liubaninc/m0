package model

type Events struct {
	ID       uint   `json:"-" gorm:"primarykey" `
	Operator string `json:"operator"`
	Route    string `json:"route" gorm:"index" `
	Type     string `json:"action" gorm:"index" `
	Object   string `json:"object"`
	Detail   string `json:"deatil" `
	Hash     string `json:"hash"`
	Height   int64  `json:"height"`
	Time     string `json:"time" `
}
