package api

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/memstore"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"github.com/gin-gonic/gin"
	msdk "github.com/liubaninc/m0/sdk"
	"github.com/tendermint/tendermint/libs/log"
	"gorm.io/gorm"
)

type API struct {
	logger log.Logger
	db     *gorm.DB
	client msdk.Client
}

func New(db *gorm.DB, client msdk.Client, logger log.Logger) *API {
	return &API{
		db:     db,
		client: client,
		logger: logger.With("module", "api"),
	}
}

func (api *API) Run(port int) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	store := memstore.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.Use(func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, Authorization")
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	})
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1 := router.Group("/api")
	v1.GET("/faucet/:address", api.Faucet)
	v1.GET("/search/:content", api.GetSearch)
	v1.GET("/charts", api.GetCharts)
	v1.GET("/peers", api.GetPeers)
	v1.GET("/tps/max", api.GetMaxTPS)
	v1.GET("/blockchain", api.GetBlockChain)
	v1.GET("/blocks", api.GetBlocks)
	v1.GET("/blocks/:id", api.GetBlock)
	v1.GET("/transactions", api.GetTransactions)
	v1.GET("/transactions/:hash", api.GetTransaction)
	v1.GET("/addresses", api.GetAddresses)
	v1.GET("/addresses/:name", api.GetAddress)
	v1.GET("/addresses/:name/assets", api.GetAddressAssets)
	v1.GET("/addresses/:name/contracts", api.GetAddressContracts)
	v1.GET("/addresses/:name/transactions", api.GetAddressTransactions)
	v1.GET("/assets", api.GetAssets)
	v1.GET("/assets/:name", api.GetAsset)
	v1.GET("/assets/:name/transactions", api.GetAssetTransactions)
	v1.GET("/contracts", api.GetContracts)
	v1.GET("/contracts/:name", api.GetContract)
	v1.GET("/contracts/:name/transactions", api.GetContractTransactions)
	v1.GET("/events/peer", api.GetPeerEvents)
	v1.GET("/events/validator", api.GetValidatorEvents)
	v1.GET("/events/cert", api.GetCertEvents)
	v1.GET("/events/account", api.GetAccountEvents)
	v1.GET("/events/contract", api.GetContractEvents)
	// wallet
	v1.GET("/captcha", GetCaptcha)
	v1.GET("/captcha/:captchaId", GetCaptchaPNG)
	v1.POST("/user/register", api.UserRegister)
	v1.POST("/user/login", api.UserLogin)
	v1.GET("/download/:hash", api.DownloadTx)
	v1.GET("/claims/download/:user/:account/:name", api.DownloadClaim)

	//钱包合约模块
	v1.GET("/mcontract/tx/download/:hash", api.DownloadMContractTx)
	v1.GET("/mcontract/download/:id", api.DownloadMContractFile)
	v1.GET("/mcontract/download/sdk", api.DownloadSDk)

	v1.Use(api.authorize)
	v1.POST("/account/mnemonic", api.AccountMnemonic)
	v1.POST("/account/create", api.AccountCreate)
	v1.POST("/account/create_multisig", api.AccountCreateMultiSig)
	v1.POST("/account/export", api.AccountExport)
	v1.POST("/accounts", api.AccountList)
	v1.POST("/accounts/:name", api.Account)
	v1.POST("/tx/mint", api.Mint)
	v1.POST("/tx/burn", api.Burn)
	v1.POST("/tx/transfer", api.Transfer)
	v1.POST("/tx/sign", api.Sign)
	//v1.GET("/tx/download/:hash", api.DownloadTx)
	v1.POST("/tx/upload", api.UploadTx)
	v1.POST("/claims/:account", api.ClaimList)
	v1.POST("/claims/:account/verify", api.ClaimVerify)
	v1.POST("/claims/:account/upload", api.UploadClaim)
	v1.POST("/claims/:account/tx", api.ClaimTx)
	v1.POST("/claims/:account/get", api.Claim)
	v1.POST("/user/logout", api.UserLogout)
	//钱包合约模块
	v1.GET("/mcontract/transactions/:hash", api.GetContractTx)
	v1.POST("/mcontract/template/list", api.MContractTemplateList)
	v1.GET("/mcontract/template/get/:id", api.GetMContractTemplate)
	v1.GET("/mcontract/get/:id", api.GetMContract)
	v1.GET("/mcontract/history/list/:contractName", api.MContractHistoryList)
	v1.POST("/mcontract/list/:account", api.MContractList)
	v1.POST("mcontract/tx/sign", api.MContractSign)
	v1.POST("/mcontract/operate", api.OperateContract)
	v1.POST("/mcontract/delete/:id", api.DeleteContract)
	v1.POST("/mcontract/create", api.CreateMContract)
	v1.POST("/mcontract/tx/upload", api.MContractUploadTx)

	v1.POST("/mcontract/template/insert/:account", api.MContractTemplateInsert)
	v1.POST("/mcontract/template/function/insert", api.MContractTemplateFunctionInsert)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		panic(err)
	}
}

var (
	// OKMsg msg
	OKMsg = "SUCCESS"
	// OKCode Ok
	OKCode   = 200
	PermCode = 201
	AuthCode = 3000
	//RequestCode Api request error
	RequestCode = 3001
	//ExecuteCode Api execute error
	ExecuteCode = 3002
)

type Response struct {
	Code   int         `json:"code"`
	Msg    string      `json:"msg"`
	Detail string      `json:"detail"`
	Data   interface{} `json:"data"`
}

type PageRequest struct {
	PageNum  int `form:"page_num" json:"page_num" xml:"page_num"`
	PageSize int `form:"page_size" json:"page_size" xml:"page_size"`
}

type PageResponse struct {
	Total     int64 `json:"total"`
	PageTotal int64 `json:"page_total"`
	PageNum   int   `json:"page_num"`
	PageSize  int   `json:"page_size"`
}
