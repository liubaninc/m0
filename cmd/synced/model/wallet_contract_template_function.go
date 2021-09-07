package model

type MContractTemplateFunction struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	CreatedAt   TimeNormal `json:"created_at"`
	Name        string     `json:"name" gorm:"type:varchar(255);comment:函数名"`
	Args        string     `json:"args" gorm:"type:varchar(800);comment:函数参数"`
	Description string     `json:"description" gorm:"type:varchar(800);comment:函数简介"`
	TemplateId  uint       `json:"template_id" gorm:"comment:模板Id"`
}
