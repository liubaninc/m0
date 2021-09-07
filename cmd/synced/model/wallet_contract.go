package model

type MContract struct {
	ID           uint       `json:"id" gorm:"primarykey"`
	CreatedAt    TimeNormal `json:"created_at"`
	UpdatedAt    TimeNormal `json:"updated_at"`
	Name         string     `json:"name" gorm:"type:varchar(255);comment:合约名称"`
	Args         string     `json:"args" gorm:"type:varchar(255);comment:合约参数"`
	Description  string     `json:"description" gorm:"type:varchar(800);comment:合约描述"`
	Version      string     `json:"version" gorm:"type:varchar(16);comment:合约版本"`
	Type         int8       `json:"type" gorm:"type:int(1);comment:生成方式：1 自定义合约上传 2 模板合约"`
	TemplateId   uint       `json:"template_id" gorm:"comment:模板合约，模板Id"`
	FileName     string     `json:"file_name" gorm:"type:varchar(255);comment:自定义合约，合约文件名称"`
	Status       int8       `json:"status" gorm:"type:int(1);comment:合约状态0待部署 1部署中 2已部署 3部署失败 4已冻结 5已解冻 6升级中 7升级失败 8升级成功 9已删除 "`
	AllianceName string     `json:"alliance_name" gorm:"type:varchar(100);default:M0;comment:链名称"`
	Address      string     `json:"address" gorm:"type:varchar(255);not null;comment:账户地址"`
	Fees         string     `json:"fees"`
	Memo         string     `json:"memo"`
	Hash         string     `json:"hash" gorm:""`
	Raw          string     `json:"-"`
	Log          string     `json:"-"`
	Mode         string     `json:"mode" gorm:"type:varchar(32);comment;合约操作 undeploy删除合约 deploy部署合约 upgrade升级合约 freeze冻结合约 unfreeze解冻合约"`
}
