package model

import (
	"time"
)

type MContracTemplate struct {
	ID                        uint `gorm:"primarykey"`
	CreatedAt                 time.Time
	Name                      string                     `json:"name" gorm:"type:varchar(255);comment:合约模板名称"`
	Description               string                     `json:"description" gorm:"type:varchar(2000);comment:合约描述"`
	Address                   string                     `json:"address" gorm:"type:varchar(255);comment:账户地址"`
	CodeFile                  string                     `json:"code_file" gorm:"type:varchar(800);comment:合约编码"`
	MContracTemplateFunctions []MContracTemplateFunction `json:"functions" gorm:"foreignKey:TemplateId"`
}
