package api

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"io/ioutil"
	"net/http"
	"strings"

	utxotypes "github.com/liubaninc/m0/x/utxo/types"
	wasmtypes "github.com/liubaninc/m0/x/wasm/types"

	"github.com/cosmos/cosmos-sdk/x/auth/signing"

	"github.com/liubaninc/m0/cmd/synced/syncer"
	"github.com/tendermint/tendermint/crypto/tmhash"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	"github.com/tendermint/tendermint/types/time"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (req *UTXORequest) ValidateBasic(self_check bool) error {
	if _, err := sdk.AccAddressFromBech32(req.From); err != nil {
		return fmt.Errorf("发送方 %s %s", req.From, ERROR_ADDRESS)
	}
	for _, receiver := range req.Receivers {
		if _, err := sdk.AccAddressFromBech32(receiver.To); err != nil {
			return fmt.Errorf("接收方 %s %s", receiver.To, ERROR_ADDRESS)
		}

		if self_check && strings.Compare(req.From, receiver.To) == 0 {
			return fmt.Errorf("接收方 %s %s", receiver.To, ERROR_ADDRESS_SELF)
		}

		if _, err := sdk.ParseCoinsNormalized(receiver.Amount); err != nil {
			return fmt.Errorf("接收金额 %s %s", receiver.Amount, ERROR_COIN)
		}
	}
	for _, fee := range req.Fees {
		if _, err := sdk.ParseCoinsNormalized(fee); err != nil {
			return fmt.Errorf("手续费 %s %s", fee, ERROR_COIN)
		}
	}

	return nil
}

func (api *API) getUtxoResponse(c *gin.Context, tp string, request *UTXORequest) (*TxResponse, error) {
	// 用户账户地址检查
	var acct model.Account
	if result := api.db.Where(&model.Account{
		Address: request.From,
		UserID:  api.userID(c),
	}).First(&acct); result.Error != nil {
		return nil, result.Error
	}

	tos := make([]string, len(request.Receivers))
	amounts := make([]string, len(request.Receivers))
	for index, receiver := range request.Receivers {
		tos[index] = receiver.To
		amounts[index] = receiver.Amount
	}
	fee := strings.Join(request.Fees, ",")

	var msg sdk.Msg
	var err error
	switch tp {
	case "mint":
		msg, err = api.client.IssueMsg(request.From, tos, amounts, request.Desc, fee)
	case "burn":
		msg, err = api.client.IssueMsg(request.From, nil, amounts, request.Desc, fee)
	case "transfer":
		msg, err = api.client.SendMsg(request.From, tos, amounts, request.Desc, fee)
	default:
		panic("unknown tp")
	}
	if err != nil {
		return nil, err
	}

	var resp TxResponse
	tx, err := api.client.GenerateTx(request.From, fee, request.Memo, 0, msg)
	if err != nil {
		return nil, err
	}
	if acct.Threshold > 0 {
		resp.MultiAddress = acct.Address
		resp.MultiPublic = acct.PublicKey
		resp.Threshold = acct.Threshold
		// 多签地址 仅仅构建交易
		if len(acct.Related) > 0 {
			kb := api.getKeyBase(c)
			if err := kb.ImportPrivKey(acct.Related, acct.PrivateKey, request.Password); err != nil {
				return nil, err
			}
			if err := api.client.SignTx(acct.Related, acct.Address, tx, true); err != nil {
				return nil, err
			}
		}
		resp.Signatures = []string{}
		signatures, _ := tx.GetTx().GetSignaturesV2()
		for _, signature := range signatures {
			addr, err := sdk.AccAddressFromHex(signature.PubKey.Address().String())
			if err != nil {
				panic(err)
			}
			resp.Signatures = append(resp.Signatures, addr.String())
		}
		// resp.Tx = string(api.client.Codec.MustMarshalJSON(tx))
		bts, _ := api.client.TxConfig.TxEncoder()(tx.GetTx())
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
	} else {
		// 签名
		kb := api.getKeyBase(c)
		if err := kb.ImportPrivKey(acct.Name, acct.PrivateKey, request.Password); err != nil {
			return nil, err
		}
		if err := api.client.SignTx(acct.Name, "", tx, true); err != nil {
			return nil, err
		}

		if request.Commit {
			// 广播交易
			result, err := api.client.BroadcastTx(tx.GetTx())
			if err != nil {
				return nil, err
			}
			if result.Code != 0 {
				return nil, fmt.Errorf(result.RawLog)
			}
			resp.Hash = result.TxHash
		} else {
			bts, _ := api.client.TxConfig.TxEncoder()(tx.GetTx())
			// resp.Tx = string(api.client.Codec.MustMarshalJSON(tx))
			resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
		}
	}
	if err := api.processTx(tx.GetTx(), request.Commit); err != nil {
		panic(err)
	}
	return &resp, nil
}

// @发行/增发资产
// @Summary 发行/增发资产
// @Description
// @Tags tx
// @Accept  json
// @Produce json
// @Param tx body UTXORequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /tx/mint [post]
func (api *API) Mint(c *gin.Context) {
	var request UTXORequest
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

	if err := request.ValidateBasic(false); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	data, err := api.getUtxoResponse(c, "mint", &request)
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
	response.Data = data
	c.JSON(http.StatusOK, response)
}

// @销毁资产
// @Summary 销毁资产
// @Description
// @Tags tx
// @Accept  json
// @Produce json
// @Param tx body UTXORequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /tx/burn [post]
func (api *API) Burn(c *gin.Context) {
	var request UTXORequest
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

	if err := request.ValidateBasic(false); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	data, err := api.getUtxoResponse(c, "burn", &request)
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
	response.Data = data
	c.JSON(http.StatusOK, response)
}

// @转移资产
// @Summary 转移资产
// @Description
// @Tags tx
// @Accept  json
// @Produce json
// @Param tx body UTXORequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /tx/transfer [post]
func (api *API) Transfer(c *gin.Context) {
	var request UTXORequest
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

	if err := request.ValidateBasic(true); err != nil {
		response.Code = ExecuteCode
		response.Msg = err.Error()
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	data, err := api.getUtxoResponse(c, "transfer", &request)
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
	response.Data = data
	c.JSON(http.StatusOK, response)
}

// @签名
// @Summary 签名
// @Description
// @Tags tx
// @Accept  json
// @Produce json
// @Param tx body SignRequest true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /tx/sign [post]
func (api *API) Sign(c *gin.Context) {
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

	var mtx model.MTransaction
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
		err = api.client.SignTx(acct.Related, multiAddress, txBuilder, false)
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
			if err := api.client.LegacyAmino.UnmarshalBinaryBare(publicKeyBytes, &multiSigPub);err != nil {
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
			if err := api.client.MultiSignTx(txBuilder, multiSigPub, signatures); err != nil {
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
		//api.logger.Info("broadcast", "tx", string(api.client.JSONMarshaler.MustMarshalJSON(txBuilder.GetTx())))
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
		//tx.Signatures = signatures
		bts, _ := api.client.TxConfig.TxEncoder()(tx)
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
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
		// resp.Tx = string(api.client.Codec.MustMarshalJSON(tx))
		bts, _ := api.client.TxConfig.TxEncoder()(tx)
		resp.Hash = fmt.Sprintf("%X", tmhash.Sum(bts))
	}
	if err := api.processTx(txBuilder.GetTx(), request.Commit); err != nil {
		panic(err)
	}
	response.Data = resp
	c.JSON(http.StatusOK, response)
}

func (api *API) processTx(tx signing.Tx, commit bool) error {
	bts, _ := api.client.TxConfig.TxEncoder()(tx)
	hash := fmt.Sprintf("%X", tmhash.Sum(bts))
	var ttx model.MTransaction
	if result := api.db.Find(&ttx, map[string]interface{}{"hash": hash}); result.Error != nil {
		return result.Error
	} else if result.RowsAffected > 0 {
		if uuid := tx.GetMemo(); len(uuid) > 0 {
			if result := api.db.Model(&model.Claim{}).Where("uuid = ?", uuid).Update("hash", hash); result.Error != nil {
				return result.Error
			}
		}
		return nil
	}
	fmt.Println("cccc", hash)

	db := api.db.Begin()
	mtx := &model.MTransaction{
		Hash:   hash,
		Size:   len(bts),
		Memo:   tx.GetMemo(),
		Fee:    tx.GetFee().String(),
		Status: false,
		Time:   time.Now().In(syncer.Local).Format(syncer.TIME_FORMAT),
		Height: 0,
		MsgNum: len(tx.GetMsgs()),
	}

	types := make([]string, mtx.MsgNum)
	assetsMap := map[string]bool{}
	var assets []string
	addressesMap := map[string]bool{}
	var addresses []string
	contractsMap := map[string]bool{}
	var contracts []string
	for index, msg := range tx.GetMsgs() {
		switch msg := msg.(type) {
		case *utxotypes.MsgIssue:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgIssue(msg))
		case *utxotypes.MsgDestroy:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgDestroy(msg))
		case *utxotypes.MsgSend:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgSend(msg))
		case *wasmtypes.MsgDeploy:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgDeploy(msg))
		case *wasmtypes.MsgInvoke:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgInvoke(msg))
		case *wasmtypes.MsgUpgrade:
			mtx.UTXOMsgs = append(mtx.UTXOMsgs, syncer.ProcessMsgUpgrade(msg))
		default:
			return fmt.Errorf("not support route %s type %s", msg.Route(), msg.Type())
		}
		umsg := mtx.UTXOMsgs[index]
		types[index] = umsg.Type
		taddresses := strings.Split(umsg.Addresses, ",")
		for _, addr := range taddresses {
			if len(addr) == 0 {
				continue
			}
			if _, ok := addressesMap[addr]; !ok {
				addressesMap[addr] = true
				addresses = append(addresses, addr)
			}
		}
		tassets := strings.Split(umsg.Assets, ",")
		for _, asset := range tassets {
			if len(asset) == 0 {
				continue
			}
			if _, ok := assetsMap[asset]; !ok {
				assetsMap[asset] = true
				assets = append(assets, asset)
			}
		}
		tcontracts := strings.Split(umsg.Contracts, ",")
		for _, contract := range tcontracts {
			if len(contract) == 0 {
				continue
			}
			if _, ok := contractsMap[contract]; !ok {
				contractsMap[contract] = true
				contracts = append(contracts, contract)
			} else if strings.HasPrefix(contract, ":") {
				contracts = append(contracts, contract)
			}
		}
	}
	if len(addresses) > 0 {
		mtx.Addresses = "," + strings.Join(addresses, ",") + ","
	}
	if len(assets) > 0 {
		mtx.Assets = "," + strings.Join(assets, ",") + ","
	}
	if len(contracts) > 0 {
		mtx.Contracts = "," + strings.Join(contracts, ",") + ","
	}
	mtx.Type = strings.Join(types, ",")

	// added
	if commit {
		mtx.Confirmed = 0
	} else {
		mtx.Confirmed = -1
	}
	var signatures []string
	signs, _ := tx.GetSignaturesV2()
	for _, signature := range signs {
		addr, err := sdk.AccAddressFromHex(signature.PubKey.Address().String())
		if err != nil {
			panic(err)
		}
		signatures = append(signatures, addr.String())
	}
	mtx.Signature = strings.Join(signatures, ",")
	bts, _ = api.client.TxConfig.TxJSONEncoder()(tx)
	mtx.Raw = string(bts)
	bts, err := json.Marshal(mtx.UTXOMsgs)
	if err != nil {
		return err
	}
	mtx.Msgs = string(bts)

	if result := db.Save(mtx); result.Error != nil {
		db.Rollback()
		return result.Error
	}
	db.Commit()
	return nil
}

// @下载交易
// @Summary 下载交易
// @Description
// @Tags tx
// @Accept  json
// @Produce json
// @Param hash path string true "请求信息"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /download/{hash} [get]
func (api *API) DownloadTx(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	hash := c.Param("hash")
	var mtx model.MTransaction
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

// @上传交易
// @Summary 上传交易
// @Description
// @Tags tx
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /tx/upload [post]
func (api *API) UploadTx(c *gin.Context) {
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

	if err := api.processTx(tx.(signing.Tx), false); err != nil {
		panic(err)
	}
	bts, _ = api.client.TxConfig.TxEncoder()(tx)
	hash := fmt.Sprintf("%X", tmhash.Sum(bts))
	response.Data = hash
	c.JSON(http.StatusOK, response)
}
