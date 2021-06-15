package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Nick     string `json:"nick"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`

	Accounts []Account `json:"accounts"`
}
