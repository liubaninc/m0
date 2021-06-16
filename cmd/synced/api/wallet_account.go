package api

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/keys/sm2"

	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client/flags"

	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/bartekn/go-bip39"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	mnemonicEntropySize = 256
	bip39Password       = ""
)

// @生成助记词
// @Summary 生成助记词
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /account/mnemonic [post]
func (api *API) AccountMnemonic(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	entropy, err := bip39.NewEntropy(mnemonicEntropySize)
	if err != nil {
		panic(err)
	}
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		panic(err)
	}
	response.Data = mnemonic
	c.JSON(http.StatusOK, response)
}

func (api *API) userID(c *gin.Context) uint {
	session := sessions.Default(c)
	return session.Get(userIDKey).(uint)
}

func (api *API) userName(c *gin.Context) string {
	session := sessions.Default(c)
	return session.Get(userNameKey).(string)
}

func (api *API) getKeyBase(c *gin.Context) keyring.Keyring {
	userName := api.userName(c)
	kr, err := keyring.New(userName, keyring.BackendMemory, viper.GetString(flags.FlagHome), os.Stdin)
	if err != nil {
		panic(err)
	}
	return kr
}

func toAccount(acct *model.Account) *AccountResponse {
	return &AccountResponse{
		Name:      acct.Name,
		Address:   acct.Address,
		PublicKey: acct.PublicKey,
		MultiSig: func() []string {
			if len(acct.MultiSig) != 0 {
				return strings.Split(acct.MultiSig, ",")
			}
			return []string{}
		}(),
		Threshold: acct.Threshold,
		Related:   acct.Related,
		Algo:      acct.Info,
	}
}

// @创建/导入单签钱包账户
// @Summary 新建钱包账户
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Param account body AccountImportRequest true "钱包账户信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /account/create [post]
func (api *API) AccountCreate(c *gin.Context) {
	var request AccountImportRequest
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

	var account model.Account
	if result := api.db.Where(&model.Account{
		UserID: api.userID(c),
		Name:   request.Name,
	}).Find(&account); result.RowsAffected != 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_EXIST
		response.Detail = fmt.Sprintf("account %s alreay existed", request.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	kb := api.getKeyBase(c)
	signingAlgoList, _ := kb.SupportedAlgorithms()
	signingAlgo, err := keyring.NewSigningAlgoFromString(strings.ToLower(request.Algo), signingAlgoList)
	if err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_ALGO
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var info keyring.Info
	if len(request.Mnemonic) > 0 {
		if valid := bip39.IsMnemonicValid(request.Mnemonic); !valid {
			response.Code = RequestCode
			response.Msg = ERROR_PRIVKEY
			response.Detail = fmt.Sprintf("invalid mnemonic %s", request.Mnemonic)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		info, err = kb.NewAccount(request.Name, request.Mnemonic, bip39Password, sdk.GetConfig().GetFullFundraiserPath(), signingAlgo)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	} else if len(request.PrivateKeyArmor) > 0 {
		if err := kb.ImportPrivKey(request.Name, request.PrivateKeyArmor, request.Password); err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		info, err = kb.Key(request.Name)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	} else if len(request.PrivateKey) > 0 {
		privateKeyBytes, err := hex.DecodeString(request.PrivateKey)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_PRIVKEY
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		switch signingAlgo {
		case hd.Secp256k1:
			if len(privateKeyBytes) != secp256k1.PrivKeySize {
				response.Code = ExecuteCode
				response.Msg = ERROR_PRIVKEY
				response.Detail = fmt.Sprintf("invalid size for %s's PubKey Got %d expected %d", request.Algo, len(privateKeyBytes), secp256k1.PrivKeySize)
				api.logger.Error(c.Request.URL.Path, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
		case hd.SM2:
			if len(privateKeyBytes) != sm2.PrivKeySize {
				response.Code = ExecuteCode
				response.Msg = ERROR_PRIVKEY
				response.Detail = fmt.Sprintf("invalid size for %s's PubKey Got %d expected %d", request.Algo, len(privateKeyBytes), sm2.PrivKeySize)
				api.logger.Error(c.Request.URL.Path, "error", response.Detail)
				c.JSON(http.StatusOK, response)
				return
			}
		}
		privateKeyArmor := crypto.EncryptArmorPrivKey(signingAlgo.Generate()(privateKeyBytes), request.Password, string(signingAlgo.Name()))
		if err := kb.ImportPrivKey(request.Name, privateKeyArmor, request.Password); err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		info, err = kb.Key(request.Name)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		info, request.Mnemonic, err = kb.NewMnemonic(request.Name, keyring.English, sdk.GetConfig().GetFullFundraiserPath(), signingAlgo)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	}

	privateKeyArmor, err := kb.ExportPrivKeyArmor(request.Name, request.Password)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	privateKey, _, err := crypto.UnarmorDecryptPrivKey(privateKeyArmor, request.Password)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	acct := &model.Account{
		Name:       request.Name,
		Address:    info.GetAddress().String(),
		Mnemonic:   request.Mnemonic,
		PrivateKey: privateKeyArmor,
		PublicKey:  hex.EncodeToString(api.client.LegacyAmino.MustMarshalBinaryBare(info.GetPubKey())),
		Info:       string(signingAlgo.Name()),
		UserID:     api.userID(c),
	}

	if err := api.Faucet(acct.Address); err != nil {
		api.logger.Error(c.Request.URL.Path, "error", err)
	}

	if result := api.db.Create(acct); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Data = &AccountExportResponse{
		AccountResponse: toAccount(acct),
		Mnemonic:        request.Mnemonic,
		PrivateKeyArmor: privateKeyArmor,
		PrivateKey:      hex.EncodeToString(privateKey.Bytes()),
	}
	c.JSON(http.StatusOK, response)
}

// @创建多签钱包账户
// @Summary 新建钱包账户
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Param account body MultiAccountImportRequest true "钱包账户信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /account/create_multisig [post]
func (api *API) AccountCreateMultiSig(c *gin.Context) {
	var request MultiAccountImportRequest
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

	if request.Threshold < 1 {
		response.Code = ExecuteCode
		response.Msg = ERROR_Threshold
		response.Detail = "threshold must be a positive integer"
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	if nKeys := len(request.MultiSig); nKeys < request.Threshold {
		response.Code = RequestCode
		response.Msg = ERROR_Threshold
		response.Detail = fmt.Sprintf("threshold %d > multisignature %d", nKeys, request.Threshold)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var account model.Account
	if result := api.db.Where(&model.Account{
		UserID: api.userID(c),
		Name:   request.Name,
	}).Find(&account); result.RowsAffected != 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_EXIST
		response.Detail = fmt.Sprintf("account %s alreay existed", request.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var related model.Account
	if result := api.db.Where(&model.Account{
		UserID: api.userID(c),
		Name:   request.Related,
	}).Find(&related); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("account %s not existed", request.Related)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var pks []cryptotypes.PubKey
	for _, sig := range request.MultiSig {
		pubKeyBytes, err := hex.DecodeString(sig)
		if err != nil {
			response.Code = RequestCode
			response.Msg = ERROR_PUBKEY
			response.Detail = fmt.Sprintf("pub %s error %s", sig, err)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		var pk cryptotypes.PubKey
		if err := api.client.LegacyAmino.Amino.UnmarshalBinaryBare(pubKeyBytes, &pk); err != nil {
			response.Code = RequestCode
			response.Msg = ERROR_PUBKEY
			response.Detail = fmt.Sprintf("pub %s error %s", sig, err)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		pks = append(pks, pk)
	}

	kb := api.getKeyBase(c)

	var info keyring.Info
	var err error
	if len(pks) != 1 {
		if request.Sort {
			sort.Slice(pks, func(i, j int) bool {
				return bytes.Compare(pks[i].Address(), pks[j].Address()) < 0
			})
		}
		pk := multisig.NewLegacyAminoPubKey(request.Threshold, pks)
		info, err = kb.SaveMultisig(request.Name, pk)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	} else {
		info, err = kb.SavePubKey(request.Name, pks[0], hd.Secp256k1Type)
		if err != nil {
			response.Code = ExecuteCode
			response.Msg = ERROR_DB
			response.Detail = err.Error()
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
	}

	acct := &model.Account{
		Name:      request.Name,
		Address:   info.GetAddress().String(),
		PublicKey: hex.EncodeToString(api.client.LegacyAmino.MustMarshalBinaryBare(info.GetPubKey())),
		UserID:    api.userID(c),
	}

	if err := api.Faucet(acct.Address); err != nil {
		api.logger.Error(c.Request.URL.Path, "error", err)
	}

	if len(pks) > 1 {
		acct.Related = request.Related
		acct.PrivateKey = related.PrivateKey
		acct.MultiSig = strings.Join(request.MultiSig, ",")
		acct.Threshold = request.Threshold
	}
	if result := api.db.Create(acct); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = toAccount(acct)
	c.JSON(http.StatusOK, response)
}

// @钱包账户导出
// @Summary 显示钱包账户
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Param account body AccountExportRequest true "导出钱包账户信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /account/export [post]
func (api *API) AccountExport(c *gin.Context) {
	var request AccountExportRequest
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

	var account model.Account
	if result := api.db.Where(&model.Account{
		Name:   request.Name,
		UserID: api.userID(c),
	}).Find(&account); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("account %s not exist", request.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	if account.Threshold > 0 {
		var taccount model.Account
		fmt.Println(account.Related, api.userID(c))
		if result := api.db.Where(&model.Account{
			Name:   account.Related,
			UserID: api.userID(c),
		}).Find(&taccount); result.RowsAffected == 0 {
			response.Code = ExecuteCode
			response.Msg = ERROR_ACCT_NO
			response.Detail = fmt.Sprintf("account %s not exist", account.Related)
			api.logger.Error(c.Request.URL.Path, "error", response.Detail)
			c.JSON(http.StatusOK, response)
			return
		}
		account = taccount
	}

	privateKey, _, err := crypto.UnarmorDecryptPrivKey(account.PrivateKey, request.Password)
	if err != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_PASSWORD
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Data = &AccountExportResponse{
		AccountResponse: toAccount(&account),
		Mnemonic:        account.Mnemonic,
		PrivateKeyArmor: account.PrivateKey,
		PrivateKey:      hex.EncodeToString(privateKey.Bytes()),
	}
	c.JSON(http.StatusOK, response)
}

type AccountsResponse struct {
	PageResponse
	List []*AccountResponse `json:"accounts"`
}

// @钱包账户列表
// @Summary 钱包账户列表
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /accounts [post]
func (api *API) AccountList(c *gin.Context) {
	var request PageRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindQuery(&request); err != nil {
		response.Code = RequestCode
		response.Msg = err.Error()
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

	var accounts []*model.Account
	if result := api.db.Where(&model.Account{
		UserID: api.userID(c),
	}).Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&accounts); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_DB
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var total int64
	if result := api.db.Model(&model.Account{}).Where(&model.Account{
		UserID: api.userID(c),
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
	response.Data = AccountsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		List: func() []*AccountResponse {
			var list []*AccountResponse
			for _, account := range accounts {
				list = append(list, toAccount(account))
			}
			return list
		}(),
	}
	c.JSON(http.StatusOK, response)
}

// @钱包账户
// @Summary 钱包账户
// @Description
// @Tags account
// @Accept  json
// @Produce json
// @Param name path string true "账户名"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /accounts/{name} [post]
func (api *API) Account(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var account model.Account
	name := c.Param("name")
	if result := api.db.Where(&model.Account{
		Name:   name,
		UserID: api.userID(c),
	}).Find(&account); result.RowsAffected == 0 {
		response.Code = ExecuteCode
		response.Msg = ERROR_ACCT_NO
		response.Detail = fmt.Sprintf("account %s not exist", name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Data = toAccount(&account)
	c.JSON(http.StatusOK, response)
}
