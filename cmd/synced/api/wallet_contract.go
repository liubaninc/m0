package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmos "github.com/tendermint/tendermint/libs/os"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type Contract struct {
	Name        string `json:"name"`
	Args        string `json:"args"`
	Description string `json:"description"`
	Version     string `json:"version"`
}

type MContractsResponse struct {
	PageResponse
	List interface{} `json:"list"`
}

const home = "/Users/admin/.synced/"

// @合约模板方法导入
// @Summary 合约模板方法导入
// @Description
// @Tags contractTemplate
// @Accept  multipart/form-data
// @Produce json
// @Param description formData string false "合约方法描述"
// @Param args formData string false "合约方法参数"
// @Param name formData string true "合约方法名称"
// @Param template_id formData string true "合约方模板Id"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/template/function/insert [POST]
func (api *API) MContractTemplateFunctionInsert(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	description := c.PostForm("description")
	name := c.PostForm("name")
	args := c.PostForm("args")
	template_id, _ := strconv.Atoi(c.PostForm("template_id"))

	if result := api.db.Save(&model.MContractTemplateFunction{Description: description, Name: name, Args: args, TemplateId: uint(template_id)}); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// @合约模板导入
// @Summary 合约模板导入
// @Description
// @Tags contractTemplate
// @Accept  multipart/form-data
// @Produce json
// @Param account formData string false "账户名称"
// @Param codefile formData file true "合约文件"
// @Param description formData string true "合约描述"
// @Param language formData string false "合约语言"
// @Param name formData string true "合约名称"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/template/insert/{account} [POST]
func (api *API) MContractTemplateInsert(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	description := c.PostForm("description")
	language := c.PostForm("language")
	name := c.PostForm("name")
	codefile, _ := c.FormFile("codefile")
	file, _ := codefile.Open()
	code, _ := ioutil.ReadAll(file)
	account := c.PostForm("account")
	acct := &model.Account{}
	if len(account) > 0 {
		acct = api.getAccountByAddress(c, account, api.userID(c))
		if acct == nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_ACCT_NO
			response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), account)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	}

	if result := api.db.Save(&model.MContractTemplate{
		Address:     acct.Address,
		Description: description,
		Language:    language,
		Name:        name,
		CodeFile:    code}); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, response)
}

// @合约模板列表
// @Summary 合约模板列表
// @Description
// @Tags contractTemplate
// @Accept  json
// @Produce json
// @Param prequest body PageRequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/template/list [POST]
func (api *API) MContractTemplateList(c *gin.Context) {
	var request PageRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	if request.PageNum < 1 {
		request.PageNum = 1
	}
	if request.PageSize < 1 {
		request.PageSize = 10
	}
	offset := (request.PageNum - 1) * request.PageSize

	var templates []*model.MContractTemplate
	if result := api.db.Table(new(model.MContractTemplate).TableName()).Select("id,created_at,name,description,language,address").Offset(offset).Limit(request.PageSize).Scan(&templates); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var total int64
	if result := api.db.Model(&model.MContractTemplate{}).Count(&total); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = MContractsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		List: templates,
	}
	c.JSON(http.StatusOK, response)
}

// @合约模板详情
// @Summary 合约模板详情
// @Description
// @Tags contractTemplate
// @Accept json
// @Produce json
// @Param id path string true "合约模板ID"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/template/get/{id} [get]
func (api *API) GetMContractTemplate(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var template model.MContractTemplate
	if result := api.db.Table(new(model.MContractTemplate).TableName()).Select("id,created_at,name,description,language,address").Where("id = ?", id).Scan(&template); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_FILE_NO
		response.Detail = fmt.Sprintf("template not exist")
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	var templatefunctions []model.MContractTemplateFunction
	if result := api.db.Where("template_id = ?", template.ID).Find(&templatefunctions); result.RowsAffected > 0 {
		template.MContractTemplateFunctions = templatefunctions
	}
	response.Data = template
	c.JSON(http.StatusOK, response)
}

// @合约详情
// @Summary 合约详情
// @Description
// @Tags contract
// @Accept json
// @Produce json
// @Param id path string true "合约ID"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/get/{id} [get]
func (api *API) GetMContract(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	id, _ := strconv.Atoi(c.Param("id"))
	var contract model.MContract
	if result := api.db.First(&contract, id); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_FILE_NO
		response.Detail = fmt.Sprintf("file not exist")
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = contract
	c.JSON(http.StatusOK, response)
}

// @合约历史列表
// @Summary 合约历史列表
// @Description
// @Tags contract
// @Accept  json
// @Produce json
// @Param contractName path string true "合约名称"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/history/list/{contractName} [get]
func (api *API) MContractHistoryList(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	contractName := c.Param("contractName")
	var contracts []*model.MContract
	if result := api.db.Where(&model.MContract{
		Name: contractName,
	}).Order("Version desc").Find(&contracts); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = contracts
	c.JSON(http.StatusOK, response)
}

// @合约列表
// @Summary 合约列表
// @Description
// @Tags contract
// @Accept  json
// @Produce json
// @Param account path string true "账户名"
// @Param prequest body PageRequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/list/{account} [POST]
func (api *API) MContractList(c *gin.Context) {
	var request PageRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	if request.PageNum < 1 {
		request.PageNum = 1
	}
	if request.PageSize < 1 {
		request.PageSize = 10
	}
	offset := (request.PageNum - 1) * request.PageSize

	name := c.Param("account")
	acct := api.getAccountByAddress(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var contracts []*model.MContract
	if result := api.db.Where("address = ? AND status <> ?", acct.Address, WASMContractStatusDeleted).Order("updated_at desc").Offset(offset).Limit(request.PageSize).Find(&contracts); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var total int64
	if result := api.db.Model(&model.MContract{}).Where(&model.MContract{
		Address: acct.Address,
	}).Count(&total); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = MContractsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		List: contracts,
	}
	c.JSON(http.StatusOK, response)
}

// @合约签名
// @Summary 签名合约
// @Description
// @Tags contract
// @Accept  json
// @Produce json
// @Param tx body SignRequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/tx/sign [post]
func (api *API) MContractSign(c *gin.Context) {
	var request SignRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	var acct model.Account
	if result := api.db.Where(&model.Account{
		Name:   request.Name,
		UserID: api.userID(c),
	}).First(&acct); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "user", api.userID(c), "account", request.Name, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var mtx model.MContract
	if result := api.db.First(&mtx, map[string]interface{}{"hash": request.Hash}); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_SIGN_TX
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "hash", request.Hash, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	tx, err := api.client.TxConfig.TxJSONDecoder()([]byte(mtx.Raw))
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_TX
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "tx", mtx.Raw, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	txBuilder, err := api.client.TxConfig.WrapTxBuilder(tx)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_TX
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "tx", mtx.Raw, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	multiAddress := acct.Address
	multiPublic := acct.PublicKey
	if len(acct.Related) > 0 {
		kb := api.getKeyBase(c)
		err := kb.ImportPrivKey(acct.Related, acct.PrivateKey, request.Password)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_PASSWORD
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "user", api.userID(c), "account", acct.Related, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		err = api.client.WithKeyring(kb).SignTx(acct.Related, multiAddress, txBuilder, false)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_SIGN
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	}

	var multiSigPub *multisig.LegacyAminoPubKey
	if len(multiAddress) > 0 {
		if len(multiPublic) > 0 {
			publicKeyBytes, err := hex.DecodeString(multiPublic)
			if err != nil {
				response.Code = ExecuteCode
				response.Msg = ERROR_PUBKEY
				response.Detail = err.Error()
				api.logger.Error(c.Request.URL.Path, "pub", multiPublic, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
			if err := api.client.LegacyAmino.UnmarshalBinaryBare(publicKeyBytes, &multiSigPub); err != nil {
				response.Code = ExecuteCode
				response.Msg = ERROR_PUBKEY
				response.Detail = err.Error()
				api.logger.Error(c.Request.URL.Path, "pub", multiPublic, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
		} else {
			res, err := api.client.GetAccount(multiAddress)
			if err != nil {
				response.Code = ExecuteCode
				response.Msg = err.Error()
				response.Detail = "not multi public address"
				api.logger.Error(c.Request.URL.Path, "address", multiAddress, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
			_ = res
			var acct authtypes.AccountI
			pub, ok := acct.GetPubKey().(*multisig.LegacyAminoPubKey)
			if !ok {
				response.Code = ExecuteCode
				response.Msg = ERROR_PUBKEY
				response.Detail = "not multi public address"
				api.logger.Error(c.Request.URL.Path, "address", multiAddress, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
			multiSigPub = pub
		}
	}

	var resp TxResponse
	if request.Commit {
		signatures, _ := txBuilder.GetTx().GetSignaturesV2()
		if len(multiAddress) > 0 {
			if err := api.client.MultiSignTx(txBuilder, multiSigPub, signatures...); err != nil {
				response.Code = ExecuteCode
				response.Msg = ERROR_SIGN
				response.Detail = err.Error()
				api.logger.Error(c.Request.URL.Path, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
		}
		// 广播交易
		result, err := api.client.BroadcastTx(txBuilder.GetTx())
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = err.Error()
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		if result.Code != 0 {
			response.Code = ExecuteCode
			response.Msg = result.RawLog
			response.Detail = result.RawLog
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		resp.Hash = result.TxHash
	} else {
		if len(multiAddress) > 0 {
			resp.MultiAddress = multiAddress
			resp.MultiPublic = multiPublic
			resp.Threshold = int(multiSigPub.Threshold)
		}
		resp.Signatures = []string{}
		signatures, _ := txBuilder.GetTx().GetSignaturesV2()
		for _, signature := range signatures {
			addr, err := sdk.AccAddressFromHex(signature.PubKey.Address().String())
			if err != nil {
				panic(err)
			}
			resp.Signatures = append(resp.Signatures, addr.String())
		}
		bts, _ := api.client.TxConfig.TxEncoder()(tx)
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
	}
	if len(resp.Hash) == 0 {
		bts, _ := api.client.TxConfig.TxEncoder()(txBuilder.GetTx())
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
	}
	if request.Commit {
		status := mtx.Status
		switch mtx.Mode {
		case WASMContractHandleDeploy:
			status = WASMContractStatusGoing
		case WASMContractHandleUpgrade:
			status = WASMContractStatusUpgradePending
		case WASMContractHandleFreeze:
			status = WASMContractStatusFrozen
		case WASMContractHandleUnfreeze:
			status = WASMContractStatusUnfrozen
		case WASMContractHandleUndeploy:
			status = WASMContractStatusDeleted
		default:
			status = -1
		}
		if result := api.db.First(&mtx).Updates(&model.MContract{
			Status: status,
			Hash:   resp.Hash}); result.Error != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = result.Error.Error()
			api.logger.Error("ContractTx", "error", response.Detail)
			c.JSON(http.StatusOK, response)
		}
	}
	response.Data = resp
	c.JSON(http.StatusOK, response)
}

// @合约部署/升级/冻结/解冻/删除
// @Summary 合约部署/升级/冻结/解冻/删除
// @Description
// @Tags contract
// @Accept  multipart/form-data
// @Produce json
// @Param id formData string true "合约Id"
// @Param mode formData string true "合约操作：undeploy删除合约 deploy部署合约 upgrade升级合约 freeze冻结合约 unfreeze解冻合约"
// @Param password formData string true "账户密码"
//@Param commit formData bool true "是否提交到节点"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/operate [post]
func (api *API) OperateContract(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	id, _ := strconv.Atoi(c.PostForm("id"))
	mode := c.PostForm("mode")
	commit, _ := strconv.ParseBool(c.PostForm("commit"))
	password := c.PostForm("password")
	var contract model.MContract
	if result := api.db.First(&contract, id); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_NO
		response.Detail = fmt.Sprintf("account %s 's contract %s not exist", contract.Address, contract.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	data, tx, err := api.contractTx(c, mode, password, commit, &contract)

	if err != nil {
		response.Code = ExecuteCode
		if strings.Contains(err.Error(), "ciphertext decryption failed") {
			response.Msg = ERROR_PASSWORD
		} else {
			response.Msg = ERROR_SEND
		}
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	status := contract.Status
	if commit {
		switch mode {
		case WASMContractHandleDeploy:
			status = WASMContractStatusGoing
		case WASMContractHandleUpgrade:
			status = WASMContractStatusUpgradePending
		case WASMContractHandleFreeze:
			status = WASMContractStatusFrozen
		case WASMContractHandleUnfreeze:
			status = WASMContractStatusUnfrozen
		case WASMContractHandleUndeploy:
			status = WASMContractStatusDeleted
		default:
			status = -1
		}

	}
	bts, _ := api.client.TxConfig.TxJSONEncoder()(tx)
	if result := api.db.Model(&contract).Updates(&model.MContract{
		Status: status,
		Hash:   data.Hash,
		Raw:    string(bts),
		Mode:   mode}); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error("ContractTx", "error", response.Detail)
		c.JSON(http.StatusOK, response)
	}
	response.Data = contract
	c.JSON(http.StatusOK, response)
}

// @合约创建
// @Summary 创建合约
// @Description
// @Tags contract
// @Accept  multipart/form-data
// @Produce json
// @Param file formData file false "合约文件"
// @Param account_name formData string true "账户名称"
// @Param name formData string true "合约名称"
// @Param args formData string false "合约参数"
// @Param version formData string true "合约版本"
// @Param description formData string false "合约描述"
// @Param type formData int true "生成方式：1 自定义合约上传 2 模板合约"
// @Param file_name formData string false "文件名称"
// @Param template_id formData int false "模板ID"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/create [post]
func (api *API) CreateMContract(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	fileName := c.PostForm("file_name")
	accountName := c.PostForm("account_name")
	name := c.PostForm("name")
	args := c.PostForm("args")
	if len(args) == 0 {
		args = "{}"
	}
	version := c.PostForm("version")
	description := c.PostForm("description")
	opreraType, _ := strconv.Atoi(c.PostForm("type"))
	templateId, _ := strconv.Atoi(c.PostForm("template_id"))

	acct := api.getAccountByAddress(c, accountName, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), accountName)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var contract model.MContract
	if result := api.db.Where(&model.MContract{
		Name: name,
		//Address: acct.Address,
		//Version: version,
	}).Find(&contract); result.RowsAffected > 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_EXIST
		response.Detail = fmt.Sprintf("contract %s alreay exist", name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	//生成方式为上传合约时，保存文件
	if int8(opreraType) == MContractCustomType {
		response = api.uploadMContractFile(c, name, acct, version)
		if response.Code != OKCode {
			c.JSON(http.StatusOK, response)
			return
		}
		if len(fileName) == 0 {
			file, _ := c.FormFile("file")
			fileName = file.Filename
		}
	}

	wasm := &model.MContract{
		Name:         name,
		Args:         args,
		Description:  description,
		Version:      version,
		Type:         int8(opreraType),
		TemplateId:   uint(templateId),
		FileName:     fileName,
		Status:       WASMContractStatusPending,
		AllianceName: AllianceName,
		Address:      acct.Address,
	}
	if result := api.db.Save(wasm); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error("ContractTx", "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = wasm
	c.JSON(http.StatusOK, response)
}

// @上传合约交易
// @Summary 上传合约交易
// @Description
// @Tags contract
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/tx/upload [post]
func (api *API) MContractUploadTx(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "file", header.Filename, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	bts, err := ioutil.ReadAll(file)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "file", header.Filename, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	tx, err := api.client.TxConfig.TxJSONDecoder()(bts)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_TX
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "tx", string(bts), "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	bts, _ = api.client.TxConfig.TxEncoder()(tx)
	hash := fmt.Sprintf("%X", tmhash.Sum(bts))
	var contract model.MContract
	if result := api.db.Where(&model.MContract{
		Hash: hash,
	}).Find(&contract); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_NO
		response.Detail = fmt.Sprintf("contract %s.%s no exist", contract.Name, contract.Version)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	mResponse := &MContractSignRespose{
		Name:        contract.Name,
		Args:        contract.Args,
		Version:     contract.Version,
		Description: contract.Description,
		Address:     contract.Address,
		Mode:        contract.Mode,
	}
	response.Data = mResponse
	c.JSON(http.StatusOK, response)
}

// @下载合约交易
// @Summary 下载合约交易
// @Description
// @Tags contract
// @Accept  json
// @Produce json
// @Param hash path string true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/tx/download/{hash} [get]
func (api *API) DownloadMContractTx(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	hash := c.Param("hash")
	var mtx model.MContract
	if result := api.db.First(&mtx, map[string]interface{}{"hash": hash}); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_NO
		api.logger.Error(c.Request.URL.Path, "hash", hash, "error", result.Error.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	fileContentDisposition := "attachment;filename=\"" + hash + ".json\""
	c.Header("Content-Type", "application/json")
	c.Header("Content-Disposition", fileContentDisposition)
	c.Data(http.StatusOK, "application/json", []byte(mtx.Raw))
}

// @下载合约文件
// @Summary 下载合约文件
// @Description
// @Tags contract
// @Accept  json
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /mcontract/download/{id} [get]
func (api *API) DownloadMContractFile(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	id, _ := strconv.Atoi(c.Param("id"))
	var contract model.MContract
	if result := api.db.First(&contract, id); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_FILE_NO
		response.Detail = fmt.Sprintf("file not exist")
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	//dst := filepath.Join(viper.GetString(flags.FlagHome), "upload/contract", contract.Version, contract.Address, contract.Name,contract.FileName)

	dst := filepath.Join(home, "upload/contract", contract.Version, contract.Address, contract.Name, contract.FileName)

	fileContentDisposition := "attachment;filename=\"" + contract.FileName + "\""
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fileContentDisposition)
	c.File(dst)
}

// 上传合约文件
func (api *API) uploadMContractFile(c *gin.Context, name string, acct *model.Account, version string) *Response {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	file, err := c.FormFile("file")
	if err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "file", file.Filename, "error", err.Error())
		return response
	}
	if ".wasm" != path.Ext(file.Filename) {
		response.Code = RequestCode
		response.Msg = ERROR_FILE_FORMAT
		response.Detail = fmt.Sprintf("file %s format not supported", file.Filename)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		return response
	}

	//上传的合约文件大小限制
	if file.Size > WASMContractSizeLimit {
		response.Code = RequestCode
		response.Msg = ERROR_FILE_SIZE
		response.Detail = fmt.Sprintf("file %s is too large", file.Filename)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		return response
	}

	//if err := tmos.EnsureDir(filepath.Join(viper.GetString(flags.FlagHome), "upload/contract/", version, acct.Address,name), 0700); err != nil {
	if err := tmos.EnsureDir(filepath.Join(home, "upload/contract/", version, acct.Address, name), 0700); err != nil {
		panic(err)
	}
	//dst := filepath.Join(viper.GetString(flags.FlagHome), "upload/contract", version, acct.Address,name, filepath.Base(file.Filename))
	dst := filepath.Join(home, "upload/contract", version, acct.Address, name, filepath.Base(file.Filename))
	if tmos.FileExists(dst) {
		response.Code = ExecuteCode
		response.Msg = ERROR_FILE_EXIST
		response.Detail = fmt.Sprintf("file %s alreay exist", file.Filename)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		return response
	}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "file", file.Filename, "error", err.Error())
		return response
	}
	response.Data = filepath.Base(dst)
	return response
}

func (api *API) contractTx(c *gin.Context, handle string, password string, commit bool, contract *model.MContract) (*TxResponse, signing.Tx, error) {
	var acct model.Account
	if result := api.db.Where(&model.Account{
		Address: contract.Address,
		UserID:  api.userID(c),
	}).First(&acct); result.Error != nil {
		return nil, nil, result.Error
	}

	var msg sdk.Msg
	var err error
	switch handle {
	case WASMContractHandleDeploy:
		var code []byte
		//模板部署
		if contract.Type == MContractTemplateType {
			var template model.MContractTemplate
			if result := api.db.First(&template, contract.TemplateId); result.RowsAffected > 0 {
				code = template.CodeFile
				contract.Description = template.Description
			}

		} else {
			//codeFile := filepath.Join(viper.GetString(flags.FlagHome), "upload/contract", contract.Version, contract.Address, contract.FileName)
			codeFile := filepath.Join(home, "upload/contract", contract.Version, contract.Address, contract.Name, contract.FileName)
			if !tmos.FileExists(codeFile) {
				return nil, nil, fmt.Errorf("file %s not exist", contract.FileName)
			}
			code, err = ioutil.ReadFile(codeFile)
			if err != nil {
				return nil, nil, err
			}
		}
		msg, err = api.client.DeployMsg(contract.Address, contract.Name, code, contract.Args, contract.Description, contract.Fees)
	case WASMContractHandleUpgrade:
		//codeFile := filepath.Join(viper.GetString(flags.FlagHome), "upload/contract", contract.Version, contract.Address,contract.FileName, contract.FileName)
		codeFile := filepath.Join(home, "upload/contract", contract.Version, contract.Address, contract.Name, contract.FileName)
		if !tmos.FileExists(codeFile) {
			return nil, nil, fmt.Errorf("file %s not exist", contract.FileName)
		}
		code, er := ioutil.ReadFile(codeFile)
		if er != nil {
			return nil, nil, er
		}
		msg, err = api.client.UpgradeMsg(contract.Address, contract.Name, code, contract.Description, contract.Fees)
	case WASMContractHandleFreeze:
		msg, err = api.client.FreezeMsg(contract.Address, contract.Name)
	case WASMContractHandleUnfreeze:
		msg, err = api.client.UnfreezeMsg(contract.Address, contract.Name)
	case WASMContractHandleUndeploy:
		msg, err = api.client.UndeployMsg(contract.Address, contract.Name)
	default:
		panic("unknown handle")
	}
	if err != nil {
		return nil, nil, err
	}

	bts, _ := json.Marshal(&Contract{
		Name:        contract.Name,
		Args:        contract.Args,
		Description: contract.Description,
		Version:     contract.Version,
	})
	var resp TxResponse
	txBuilder, err := api.client.GenerateTx(contract.Address, contract.Fees, string(bts), 0, msg)
	if err != nil {
		return nil, nil, err
	}
	if acct.Threshold > 0 { // 多签地址 仅仅构建交易
		if len(acct.Related) > 0 {
			kb := api.getKeyBase(c)
			if err := kb.ImportPrivKey(acct.Related, acct.PrivateKey, password); err != nil {
				return nil, nil, err
			}
			if err := api.client.WithKeyring(kb).SignTx(acct.Related, acct.Address, txBuilder, true); err != nil {
				return nil, nil, err
			}
		}
		if commit {
			publicKeyBytes, err := hex.DecodeString(acct.PublicKey)
			if err != nil {
				return nil, nil, err
			}
			var multiSigPub multisig.LegacyAminoPubKey
			if err := api.client.LegacyAmino.UnmarshalBinaryBare(publicKeyBytes, &multiSigPub); err != nil {
				return nil, nil, err
			}
			signatures, _ := txBuilder.GetTx().GetSignaturesV2()

			txBuilder, _ = api.client.TxConfig.WrapTxBuilder(txBuilder.GetTx())
			if err := api.client.MultiSignTx(txBuilder, &multiSigPub, signatures...); err != nil {
				return nil, nil, err
			}
			// 广播交易
			//api.client.WithBroadcastMode("sync")
			result, err := api.client.BroadcastTx(txBuilder.GetTx())
			if err != nil {
				return nil, nil, err
			}
			if result.Code != 0 {
				return nil, nil, fmt.Errorf(result.RawLog)
			}
			resp.Hash = result.TxHash
		} else {
			resp.MultiAddress = acct.Address
			resp.MultiPublic = acct.PublicKey
			resp.Threshold = acct.Threshold
			resp.Signatures = []string{}
			signatures, _ := txBuilder.GetTx().GetSignaturesV2()
			for _, signature := range signatures {
				addr, err := sdk.AccAddressFromHex(signature.PubKey.Address().String())
				if err != nil {
					panic(err)
				}
				resp.Signatures = append(resp.Signatures, addr.String())
			}
			// resp.Tx = string(api.client.Codec.MustMarshalJSON(tx))
		}
	} else {
		// 签名
		kb := api.getKeyBase(c)
		if err := kb.ImportPrivKey(acct.Name, acct.PrivateKey, password); err != nil {
			return nil, nil, err
		}
		if err := api.client.WithKeyring(kb).SignTx(acct.Name, "", txBuilder, true); err != nil {
			return nil, nil, err
		}

		if commit {
			// 广播交易
			//api.client.WithBroadcastMode("sync")
			result, err := api.client.BroadcastTx(txBuilder.GetTx())
			if err != nil {
				return nil, nil, err
			}
			if result.Code != 0 {
				return nil, nil, fmt.Errorf(result.RawLog)
			}
			resp.Hash = result.TxHash
		}
	}
	if len(resp.Hash) == 0 {
		bts, _ := api.client.TxConfig.TxEncoder()(txBuilder.GetTx())
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
	}
	return &resp, txBuilder.GetTx(), nil
}

func (api *API) getAccountByAddress(c *gin.Context, address string, usrID uint) *model.Account {
	var acct model.Account
	if result := api.db.Where(&model.Account{
		Address: address,
		UserID:  usrID,
	}).Find(&acct); result.RowsAffected == 0 {
		return nil
	}
	return &acct
}
