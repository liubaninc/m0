package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Name       string `json:"name" gorm:"gorm:"unique_index:unique_index_with_second"`
	Address    string `json:"address" gorm:"not null"`
	Mnemonic   string `json:"mnemonic"`
	MultiSig   string `json:"multi_sig"`
	Threshold  int    `json:"threshold"`
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Related    string `json:"related"`
	Info       string `json:"info"`

	UserID uint `json:"-" gorm:"unique_index:unique_index_with_first"`
}
