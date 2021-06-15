package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmos "github.com/tendermint/tendermint/libs/os"

	"github.com/cosmos/cosmos-sdk/client/flags"
)

type Claim struct {
	MD5  string `json:"md5"`
	Info string `json:"info"`
	Memo string `json:"memo"`
}

type ClaimsResponse struct {
	PageResponse
	List []*model.Claim `json:"claims"`
}

// @存证列表
// @Summary 存证列表
// @Description
// @Tags claim
// @Accept  json
// @Produce json
// @Param account path string true "账户名"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/{account} [post]
func (api *API) ClaimList(c *gin.Context) {
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
	acct := api.getAccount(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var claims []*model.Claim
	if result := api.db.Where(&model.Claim{
		Address: acct.Address,
	}).Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&claims); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var total int64
	if result := api.db.Model(&model.Claim{}).Where(&model.Claim{
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
	response.Data = ClaimsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		List: claims,
	}
	c.JSON(http.StatusOK, response)
}

// @存证详情
// @Summary 存证详情
// @Description
// @Tags claim
// @Accept json
// @Produce json
// @Param account path string true "账户名"
// @Param name formData string true "存证名"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/{account}/get [post]
func (api *API) Claim(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var request map[string]string
	if err := c.BindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	name := c.Param("account")
	acct := api.getAccount(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var claim model.Claim
	if result := api.db.Where(&model.Claim{
		Address: acct.Address,
		Name:    request["name"],
	}).Find(&claim); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_NO
		response.Detail = fmt.Sprintf("account %s 's name %s not exist", acct.Name, request["name"])
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = claim
	c.JSON(http.StatusOK, response)
}

// @下载存证文件
// @Summary 下载存证文件
// @Description
// @Tags claim
// @Accept  json
// @Produce json
// @Param user path string true "用户名"
// @Param account path string true "账户名"
// @Param name path string true "文件名"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/download/{user}/{account}/{name} [get]
func (api *API) DownloadClaim(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var usr model.User
	if result := api.db.Where(&model.User{
		Name: c.Param("user"),
	}).Find(&usr); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_USER_NO
		response.Detail = fmt.Sprintf("user %s not exist", c.Param("user"))
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	name := c.Param("account")
	acct := api.getAccount(c, name, usr.ID)
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	dst := filepath.Join(viper.GetString(flags.FlagHome), "upload", acct.Address, c.Param("name"))
	//if !tmos.FileExists(dst) {
	//	response.Code = ExecuteCode
	//	response.Msg = ERROR_FILE_NO
	//	response.Detail = fmt.Sprintf("file %s alreay exist", c.Param("name"))
	//	api.logger.Error(c.Request.URL.Path, "error", response.Detail)
	//	c.JSON(http.StatusOK, response)
	//	return
	//}

	fileContentDisposition := "attachment;filename=\"" + c.Param("name") + "\""
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fileContentDisposition)
	c.File(dst)
	//content, err := ioutil.ReadFile(dst)
	//if err != nil {
	//	panic(err)
	//}
	//c.Writer.Write(content)
}

// @上传存证文件
// @Summary 上传存证文件
// @Description
// @Tags claim
// @Accept multipart/form-data
// @Produce json
// @Param account path string true "账户名"
// @Param verify formData bool false "file"
// @Param file formData file true "file"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/{account}/upload [post]
func (api *API) UploadClaim(c *gin.Context) {
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
		c.JSON(http.StatusOK, response)
		return
	}

	name := c.Param("account")
	acct := api.getAccount(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	if strings.ToLower(c.PostForm("verify")) == "true" {
		f, err := file.Open()
		if err != nil {
			panic(err)
		}
		bts, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		response.Data = map[string]string{
			"md5":  fmt.Sprintf("%X", tmhash.Sum(bts)),
			"file": filepath.Base(file.Filename),
		}
		c.JSON(http.StatusOK, response)
		return
	}

	if err := tmos.EnsureDir(filepath.Join(viper.GetString(flags.FlagHome), "upload", acct.Address), 0700); err != nil {
		panic(err)
	}
	dst := filepath.Join(viper.GetString(flags.FlagHome), "upload", acct.Address, filepath.Base(file.Filename))
	if tmos.FileExists(dst) {
		response.Code = ExecuteCode
		response.Msg = ERROR_FILE_EXIST
		response.Detail = fmt.Sprintf("file %s alreay exist", file.Filename)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	if err := c.SaveUploadedFile(file, dst); err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "file", file.Filename, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = filepath.Base(dst)
	c.JSON(http.StatusOK, response)
}

func MD5(fileName string) (string, int) {
	bts, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%X", tmhash.Sum(bts)), len(bts)
}

// @存证上链
// @Summary 存证上链
// @Description
// @Tags claim
// @Accept  json
// @Produce json
// @Param account path string true "账户名"
// @Param tx body ClaimRequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/{account}/tx [post]
func (api *API) ClaimTx(c *gin.Context) {
	var request ClaimRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if len(request.FileName) == 0 && len(request.Content) == 0 {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = "文件或内容必须存在一个"
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	name := c.Param("account")
	acct := api.getAccount(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var tclaim model.Claim
	if result := api.db.Where(&model.Claim{
		Name:    request.Name,
		Address: acct.Address,
	}).Find(&tclaim); result.RowsAffected > 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_EXIST
		response.Detail = fmt.Sprintf("account %s's claim %s alreay exist", name, request.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	md5 := ""
	sz := 0
	if len(request.FileName) != 0 {
		dst := filepath.Join(viper.GetString(flags.FlagHome), "upload", acct.Address, filepath.Base(request.FileName))
		if !tmos.FileExists(dst) {
			response.Code = ExecuteCode
			response.Msg = ERROR_FILE_NO
			response.Detail = fmt.Sprintf("file %s not exist", request.FileName)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		md5, sz = MD5(dst)
	}

	ev := &Claim{
		MD5:  md5,
		Info: request.Content,
		Memo: request.Memo,
	}
	bts, err := json.Marshal(ev)
	if err != nil {
		panic(err)
	}

	id := uuid()
	req := &UTXORequest{
		From: acct.Address,
		Receivers: []Receiver{
			{
				To:     acct.Address,
				Amount: "1m0token",
			},
		},
		Desc:     string(bts),
		Memo:     id,
		Commit:   request.Commit,
		Password: request.Password,
	}
	data, err := api.getUtxoResponse(c, "transfer", req)
	if err != nil {
		response.Code = ExecuteCode
		if err.Error() == "ciphertext decryption failed" {
			response.Msg = ERROR_PASSWORD
		} else {
			response.Msg = ERROR_SEND
		}
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	hash := ""
	if req.Commit {
		hash = data.Hash
	}

	claim := &model.Claim{
		Address:  acct.Address,
		Name:     request.Name,
		Content:  request.Content,
		Memo:     request.Memo,
		FileName: request.FileName,
		FileSize: int64(sz),
		FileMD5:  md5,
		UUID:     id,
		Hash:     hash,
		THash:    data.Hash,
	}
	if result := api.db.Save(claim); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error("ClaimTx", "error", response.Detail)
		c.JSON(http.StatusOK, response)
	}
	response.Data = claim
	c.JSON(http.StatusOK, response)
}

// @存证验证
// @Summary 存证验证
// @Description
// @Tags claim
// @Accept json
// @Produce json
// @Param account path string true "账户名"
// @Param name formData string true "存证名"
// @Param md5 formData string true "验证文件MD5"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /claims/{account}/verify [post]
func (api *API) ClaimVerify(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var request map[string]string
	if err := c.BindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	name := c.Param("account")
	acct := api.getAccount(c, name, api.userID(c))
	if acct == nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("user %s 's account %s not exist", api.userName(c), name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var claim model.Claim
	if result := api.db.Where(&model.Claim{
		Address: acct.Address,
		Name:    request["name"],
	}).Find(&claim); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_NO
		response.Detail = fmt.Sprintf("account %s 's name %s not exist", acct.Name, request["name"])
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = claim.FileMD5 == request["md5"]
	c.JSON(http.StatusOK, response)
}

func (api *API) getAccount(c *gin.Context, name string, usrID uint) *model.Account {
	var acct model.Account
	if result := api.db.Where(&model.Account{
		Name:   name,
		UserID: usrID,
	}).Find(&acct); result.RowsAffected == 0 {
		return nil
	}
	return &acct
}

func (api *API) Faucet(address string) error {
	info, err := api.client.Keyring.Key("faucet")
	if err != nil {
		panic(err)
	}

	if res, _ := api.client.GetAccountBalance(address, "m0token"); res.Balance != nil {
		return nil
	}

	result, err := api.client.BroadcastSendTx("faucet", []string{address}, []string{"100m0token"}, "faucet send", "faucet", "")
	if err != nil {
		return err
	}

	api.logger.Info(info.GetAddress().String(), "tx", string(api.client.JSONMarshaler.MustMarshalJSON(result)))
	return nil
}

func uuid() string {
	// generate 32 bits timestamp
	unix32bits := uint32(time.Now().UTC().Unix())

	buff := make([]byte, 12)

	numRead, err := rand.Read(buff)

	if numRead != len(buff) || err != nil {
		panic(err)
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x-%x\n", unix32bits, buff[0:2], buff[2:4], buff[4:6], buff[6:8], buff[8:])
}
