package api

import (
	"fmt"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"gorm.io/gorm"
)

type AddressesReqeust struct {
	PageRequest
	Token string `form:"coin" json:"coin" xml:"coin"`
}

type AddressesResponse struct {
	PageResponse
	AddressList []*model.Address `json:"addresses,omitempty"`
}

func (api *API) GetAddresses(c *gin.Context) {
	var request AddressesReqeust
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
	var addresses []*model.Address
	var result *gorm.DB
	total := int64(0)
	if len(request.Token) > 0 {
		api.db.Model(&model.Address{}).Where("balance LIKE ?", fmt.Sprintf("%%%s,%%", request.Token)).Count(&total)
		result = api.db.Where("balance LIKE ?", fmt.Sprintf("%%%s,%%", request.Token)).Offset(offset).Limit(request.PageSize).Find(&addresses)
	} else {
		var address model.Address
		if result := api.db.Last(&address); result.Error != nil {
			api.logger.Error("GetAddresses", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
		total = int64(address.ID)
		result = api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&addresses)
	}
	if result.Error != nil {
		api.logger.Error("GetAddresses", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &AddressesResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
		},
		AddressList: func() []*model.Address {
			addrs := make([]*model.Address, len(addresses))
			for index, addr := range addresses {
				if len(request.Token) > 0 {
					addr.Fill(request.Token)
				} else {
					addr.Fill()
				}
				addrs[index] = addr
			}
			return addrs
		}(),
	}
	c.JSON(http.StatusOK, response)
}

// @地址详情
// @Summary 地址详情
// @Description
// @Tags address
// @Accept  json
// @Produce json
// @Param name path string true "地址"
// @Param coin query string false "资产名, 过滤该资产"
// @Success 200 {object} Response
// @Router /addresses/{name} [get]
func (api *API) GetAddress(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["address"] = c.Param("name")
	var address model.Address
	if result := api.db.Find(&address, cond); result.Error != nil {
		api.logger.Error("GetAddress", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		if token := c.Query("coin"); token != "" {
			address.Fill(token)
		} else {
			address.Fill()
		}
		response.Data = address
	}
	c.JSON(http.StatusOK, response)
}

type AddressAssetsResponse struct {
	PageResponse
	Coins sdk.Coins `json:"coins"`
}

// @地址下资产余额列表
// @Summary 地址下资产余额列表
// @Description
// @Tags address
// @Accept  json
// @Produce json
// @Param name path string true "地址"
// @Success 200 {object} Response
// @Router /addresses/{name}/assets [get]
func (api *API) GetAddressAssets(c *gin.Context) {
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

	cond := map[string]interface{}{}
	cond["address"] = c.Param("name")
	var address model.Address
	if result := api.db.Find(&address, cond); result.Error != nil {
		api.logger.Error("GetAddress", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		address.Fill()
	}

	total := int64(len(address.Coins))
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}

	var coins sdk.Coins
	if int64(offset) > total {

	} else if int64(offset+request.PageSize) > total {
		coins = address.Coins[offset:]
	} else {
		coins = address.Coins[offset : offset+request.PageSize]
	}
	response.Data = &AddressAssetsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		Coins: coins,
	}
	c.JSON(http.StatusOK, response)
}

type AddressTransactionsRequest struct {
	PageRequest
	Token string `form:"coin" json:"coin" xml:"coin"`
	Type  string `form:"type" json:"type" xml:"type"`
}

func (api *API) GetAddressContracts(c *gin.Context) {
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
	cond := map[string]interface{}{}
	cond["initiator"] = c.Param("name")
	var contracts []*model.Contract
	if result := api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&contracts, cond); result.Error != nil {
		api.logger.Error("GetContracts", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := int64(0)
	api.db.Model(&model.Transaction{}).Where(cond).Count(&total)
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &ContractsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		ContractList: contracts,
	}
	c.JSON(http.StatusOK, response)
}

// @地址相关的交易
// @Summary 历史交易详情
// @Description
// @Tags address
// @Accept  json
// @Produce json
// @Param name path string true "地址"
// @Param coin query string false "资产名,  过滤该资产"
// @Param type query string false "交易类型,  过滤该类型"
// @Success 200 {object} Response
// @Router /addresses/{name}/transactions [get]
func (api *API) GetAddressTransactions(c *gin.Context) {
	var request AddressTransactionsRequest
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
	var transactions []*model.Transaction
	var result *gorm.DB
	if len(request.Token) > 0 {
		if len(request.Type) > 0 {
			result = api.db.Where("addresses LIKE ? AND assets LIKE ? AND type LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%,%s,%%", request.Token), fmt.Sprintf("%%%s%%", request.Type)).Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Preload("UTXOMsgs.Inputs").Preload("UTXOMsgs.Outputs").Preload("UTXOMsgs.ContractRequests").Find(&transactions)
		} else {
			result = api.db.Where("addresses LIKE ? AND assets LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%,%s,%%", request.Token)).Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Preload("UTXOMsgs.Inputs").Preload("UTXOMsgs.Outputs").Preload("UTXOMsgs.ContractRequests").Find(&transactions)
		}
	} else {
		if len(request.Type) > 0 {
			result = api.db.Where("addresses LIKE ? AND type LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%%s%%", request.Type)).Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Preload("UTXOMsgs.Inputs").Preload("UTXOMsgs.Outputs").Preload("UTXOMsgs.ContractRequests").Find(&transactions)
		} else {
			result = api.db.Where("addresses LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name"))).Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Preload("UTXOMsgs.Inputs").Preload("UTXOMsgs.Outputs").Preload("UTXOMsgs.ContractRequests").Find(&transactions)
		}
	}
	if result.Error != nil {
		api.logger.Error("GetAddressTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	var blockchain model.BlockChain
	if len(transactions) > 0 {
		if result := api.db.Last(&blockchain); result.Error != nil {
			api.logger.Error("GetAddressTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
	}

	var total int64
	if len(request.Token) > 0 {
		if len(request.Type) > 0 {
			api.db.Model(&model.Transaction{}).Where("addresses LIKE ? AND assets LIKE ? AND type LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%,%s,%%", request.Token), fmt.Sprintf("%%%s%%", request.Type)).Count(&total)
		} else {
			api.db.Model(&model.Transaction{}).Where("addresses LIKE ? AND assets LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%,%s,%%", request.Token)).Count(&total)
		}
	} else {
		if len(request.Type) > 0 {
			api.db.Model(&model.Transaction{}).Where("addresses LIKE ? AND type LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name")), fmt.Sprintf("%%%s%%", request.Type)).Count(&total)
		} else {
			api.db.Model(&model.Transaction{}).Where("addresses LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name"))).Count(&total)
		}
	}
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &TransactionsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		TransactionList: model.FillConfirmed(transactions, blockchain.BlockNum, c.Param("name"), request.Token),
	}
	c.JSON(http.StatusOK, response)
}
