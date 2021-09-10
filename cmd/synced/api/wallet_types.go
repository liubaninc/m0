package api

type UserRegisterRequest struct {
	Name      string `json:"name" binding:"required"`      // 用户名
	Password  string `json:"password" binding:"required"`  // 用户密码
	Nick      string `json:"nick"`                         // 用户昵称
	Email     string `json:"email"`                        // 用户邮件
	Mobile    string `json:"mobile"`                       // 用户手机
	CaptchaId string `json:"captchaId" binding:"required"` // 验证码图形文件
	Captcha   string `json:"captcha" binding:"required"`   // 验证码
}

type UserLoginRequest struct {
	Name        string `json:"name" binding:"required"`      // 用户名
	Password    string `json:"password" binding:"required"`  // 用户密码
	CaptchaId   string `json:"captchaId" binding:"required"` // 验证码图形文件
	Captcha     string `json:"captcha" binding:"required"`   // 验证码
	ExpDuration int64  `json:"exp_duration" binding:""`
}

type UserResponse struct {
	Name   string `json:"name"`
	Nick   string `json:"nick"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
}

type AccountImportRequest struct {
	Name            string `json:"name" binding:"required"`     // 账户名
	Password        string `json:"password" binding:"required"` // 账户密码
	Mnemonic        string `json:"mnemonic"`                    // 助记词
	PrivateKey      string `json:"private_key"`                 // 私钥 hex/base64
	Algo            string `json:"algo" binding:"required"`
	PrivateKeyArmor string `json:"private_key_armor"` // 私钥 hex/base64
}

type MultiAccountImportRequest struct {
	Name      string   `json:"name" binding:"required"`      // 账户名
	Related   string   `json:"related" binding:"required"`   // 关联秘钥的账户名
	MultiSig  []string `json:"multi_sig" binding:"required"` // 多签公钥列表
	Threshold int      `json:"threshold" binding:"required"` // 公钥列表
	Sort      bool     `json:"sort"`                         // 多签公钥是否排序
}

type AccountResponse struct {
	Name      string   `json:"name"`       // 账户名
	Address   string   `json:"address"`    // 账户地址
	PublicKey string   `json:"public_key"` // 公钥地址
	MultiSig  []string `json:"multi_sig"`  // 多签公钥列表
	Threshold int      `json:"threshold"`  // 公钥列表
	Related   string   `json:"related"`    // 关联秘钥的账户名
	Algo      string   `json:"algo"`       // 算法
}

type AccountExportRequest struct {
	Name     string `json:"name" binding:"required"`     // 账户名
	Password string `json:"password" binding:"required"` // 账户密码
}

type AccountExportResponse struct {
	*AccountResponse
	Mnemonic        string `json:"mnemonic"`          // 助记词
	PrivateKey      string `json:"private_key"`       // 私钥 hex
	PrivateKeyArmor string `json:"private_key_armor"` // 私钥 armor
}

type Receiver struct {
	To     string `json:"to,omitempty"` // 接收地址 burn 为空， mint、transfer 不能为空
	Amount string `json:"amount"`       // 接收金额及币种
}

type UTXORequest struct {
	From      string     `json:"from" binding:"required"`     // 发送方
	Receivers []Receiver `json:"tos"`                         // 接收列表
	Fees      []string   `json:"fees"`                        // 手续费
	Desc      string     `json:"desc"`                        // 消息描述
	Memo      string     `json:"memo"`                        // 交易描述
	Commit    bool       `json:"commit"`                      // 是否提交到节点
	Password  string     `json:"password" binding:"required"` // 账户密码
}

type SignRequest struct {
	//Address      string `json:"address"`          // 签名地址
	//MultiAddress string `json:"multi_address"`    // 多钱地址
	//MultiPublic  string `json:"multi_public_key"` // 多钱地址公钥
	Name string `json:"name" binding:"required"` // 账户名
	// Tx           string `json:"tx"`               // 交易
	Hash     string `json:"hash" binding:"required"`     // 交易哈希
	Commit   bool   `json:"commit"`                      // 是否提交到节点
	Password string `json:"password" binding:"required"` // 账户密码
}

type TxResponse struct {
	Hash string `json:"hash,omitempty"` // 交易哈希
	//Tx           string `json:"tx,omitempty"`     // 交易
	MultiAddress string   `json:"multi_address"`    // 多签地址
	MultiPublic  string   `json:"multi_public_key"` // 多签地址公钥
	Signatures   []string `json:"signatures"`       // 已签名数
	Threshold    int      `json:"threshold"`        // 需签名数
}

type TxRequest struct {
	Hash string `json:"hash" binding:"required"` // 交易哈希
}

type TxsRequest struct {
	PageRequest
	Address string `json:"address"` // 地址
	Coin    string `json:"coin"`    // 地址
}

type BalancesRequest struct {
	PageRequest
	Address string `json:"address"` // 地址
}

type ClaimRequest struct {
	Name     string `json:"name" binding:"required"` // 存证名
	Content  string `json:"info"`                    // 存证信息
	Memo     string `json:"memo"`                    // 存证备注
	FileName string `json:"file"`                    // 存证文件
	Commit   bool   `json:"commit"`                  // 是否提交到节点
	Password string `json:"password"`                // 账户密码
}

type MContractRequest struct {
	AccountName string `json:"account_name"` //账户名称
	Name        string `json:"name"`         // 合约名称
	Args        string `json:"args"`         //合约参数
	Version     string `json:"version"`      // 合约版本
	Description string `json:"description"`  // 合约描述
	Type        int8   `json:"type"`         //生成方式：1 自定义合约上传 2 模板合约
	FileName    string `json:"file_name"`    // 创建自定义合约，合约文件
	TemplateId  uint   `json:"template_id"`  //创建模板合约，模板Id
}
type MContractSignRespose struct {
	Name        string   `json:"name" binding:"required"` // 合约名称
	Args        string   `json:"args"`                    //合约参数
	Version     string   `json:"version"`                 // 合约版本
	Description string   `json:"description"`             // 合约描述
	Address     string   `json:"address"`                 //发起人
	Mode        string   `json:"mode"`                    //合约操作
	Signature   string   `json:"signature"`               //签名地址
	Signatures  []string `json:"signatures"`              // 已签名数
	Hash        string   `json:"hash"`
}

const (
	AllianceName                     string = "M0"            //链名称
	MContractCustomType              int8   = 1               //自定义合约上传
	MContractTemplateType            int8   = 2               //模板合约
	WASMContractSizeLimit            int64  = 5 * 1024 * 1024 //合约文件大小限制
	WASMContractStatusPending        int8   = 0               //待部署
	WASMContractStatusGoing          int8   = 1               //部署中
	WASMContractStatusSuccess        int8   = 2               //已部署
	WASMContractStatusFail           int8   = 3               //部署失败
	WASMContractStatusFrozen         int8   = 4               //已冻结
	WASMContractStatusUnfrozen       int8   = 5               //已解冻
	WASMContractStatusUpgradePending int8   = 6               //升级中
	WASMContractStatusUpgradeFail    int8   = 7               //升级失败
	WASMContractStatusUpgradeSuccess int8   = 8               //升级成功
	WASMContractStatusDeleted        int8   = 9               //已删除
	WASMContractHandleDeploy         string = "deploy"
	WASMContractHandleUpgrade        string = "upgrade"
	WASMContractHandleFreeze         string = "freeze"
	WASMContractHandleUnfreeze       string = "unfreeze"
	WASMContractHandleUndeploy       string = "undeploy"
)

var (
	ERROR_REQ          = "请求参数不正确"
	ERROR_USER_EXIST   = "用户已存在"
	ERROR_USER_NO      = "用户不存在"
	ERROR_ACCT_EXIST   = "账户已存在"
	ERROR_ACCT_NO      = "账户不存在"
	ERROR_FILE_EXIST   = "文件已存在"
	ERROR_FILE_NO      = "文件不存在"
	ERROR_FILE_READ    = "读取文件错误"
	ERROR_FILE_WRITE   = "读取写入错误"
	ERROR_FILE_FORMAT  = "文件格式不正确"
	ERROR_FILE_SIZE    = "文件大于要求"
	ERROR_NO           = "查询不存在"
	ERROR_EXIST        = "名称已存在"
	ERROR_DB           = "执行错误"
	ERROR_CAPTCHA      = "验证码错误"
	ERROR_PASSWORD     = "密码错误"
	ERROR_PUBKEY       = "公钥格式不正确"
	ERROR_PRIVKEY      = "私钥格式不正确"
	ERROR_ALGO         = "签名算法不正确"
	ERROR_Threshold    = "签名数不正确"
	ERROR_ADDRESS      = "地址格式不正确"
	ERROR_ADDRESS_SELF = "接收方不能是自己"
	ERROR_COIN         = "资产格式不正确"
	ERROR_TX           = "交易格式不正确"
	ERROR_SIGN_TX      = "签名文件不存在"
	ERROR_SIGN         = "签名错误"
	ERROR_SEND         = "交易上链失败，请重试或联系开发人员"
)
