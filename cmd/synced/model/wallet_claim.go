package model

import "gorm.io/gorm"

type Claim struct {
	gorm.Model
	Address string `json:"address" gorm:"uniqueIndex:idx_name;size:256"`
	Name    string `json:"name" gorm:"uniqueIndex:idx_name;size:256"`
	FileMD5 string `json:"md5"`
	Content string `json:"info"`
	Memo    string `json:"memo"`
	Hash    string `json:"hash" gorm:""`
	THash   string `json:"thash" gorm:"unique;not null"`

	UUID string `json:"-"`

	FileName string `json:"file"`
	FileSize int64  `json:"size"`
}
