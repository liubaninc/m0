package model

import (
	"time"
)

type MContracTemplate struct {
	ID                        uint `gorm:"primarykey"`
	CreatedAt                 time.Time
	Name                      string                     `json:"name" gorm:"type:varchar(255);comment:合约模板名称"`
	Description               string                     `json:"description" gorm:"type:varchar(2000);comment:合约描述"`
	Language                  string                     `json:"language" gorm:"type:varchar(32);comment:合约语言"`
	Address                   string                     `json:"address" gorm:"type:varchar(255);comment:账户地址"`
	CodeFile                  []byte                     `json:"code_file" gorm:"type:longblob;comment:合约编码"`
	MContracTemplateFunctions []MContracTemplateFunction `json:"functions" gorm:"foreignKey:TemplateId"`
}
type Tabler interface {
	TableName() string
}

// TableName overrides the table name used by User to `profiles`
func (MContracTemplate) TableName() string {
	return "m_contrac_templates"
}
