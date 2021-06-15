package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"gorm.io/gorm"
)

type AssetsResponse struct {
	PageResponse
	AssetList []*model.Asset `json:"assets,omitempty"`
}

func (api *API) GetAssets(c *gin.Context) {
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
	var assets []*model.Asset
	if result := api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&assets); result.Error != nil {
		api.logger.Error("GetAssets", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if result := api.db.First(&blockchain); result.Error != nil {
		api.logger.Error("GetAssets", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.AssetNum
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &AssetsResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		AssetList: assets,
	}
	c.JSON(http.StatusOK, response)
}

// @资产详情
// @Summary 资产详情
// @Description
// @Tags asset
// @Accept  json
// @Produce json
// @Param name path string true "资产名"
// @Success 200 {object} Response
// @Router /assets/{name} [get]
func (api *API) GetAsset(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["name"] = c.Param("name")
	var asset model.Asset
	if result := api.db.Find(&asset, cond); result.Error != nil {
		api.logger.Error("GetAsset", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = asset
	}
	c.JSON(http.StatusOK, response)
}

type AssetTransactionsRequest struct {
	PageRequest
}

func (api *API) GetAssetTransactions(c *gin.Context) {
	var request AssetTransactionsRequest
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

	var result *gorm.DB
	offset := (request.PageNum - 1) * request.PageSize
	var transactions []*model.Transaction

	result = api.db.Order("ID desc").Where("assets LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name"))).Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Find(&transactions)
	if result.Error != nil {
		api.logger.Error("GetAssetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if len(transactions) > 0 {
		if result := api.db.Last(&blockchain); result.Error != nil {
			api.logger.Error("GetAssetTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
	}

	var total int64
	api.db.Model(&model.Transaction{}).Where("assets LIKE ?", fmt.Sprintf("%%,%s,%%", c.Param("name"))).Count(&total)
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
		TransactionList: model.FillConfirmed(transactions, blockchain.BlockNum, ""),
	}
	c.JSON(http.StatusOK, response)
}
