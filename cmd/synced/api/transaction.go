package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

type TransactionsResponse struct {
	PageResponse
	TransactionList []*model.Transaction `json:"txs,omitempty"`
}

func (api *API) GetTransactions(c *gin.Context) {
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
	var transactions []*model.Transaction
	if result := api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("UTXOMsgs").Find(&transactions); result.Error != nil {
		api.logger.Error("GetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if result := api.db.Last(&blockchain); result.Error != nil {
		api.logger.Error("GetTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	//var transaction model.Transaction
	//if result := api.db.Last(&transaction); result.Error != nil {
	//	api.logger.Error("GetTransactions", "error", result.Error)
	//	response.Code = ExecuteCode
	//	response.Msg = result.Error.Error()
	//	c.JSON(http.StatusOK, response)
	//	return
	//}

	total := blockchain.TxNum
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

// @交易详情
// @Summary 交易详情
// @Description
// @Tags transaction
// @Accept  json
// @Produce json
// @Param hash path string true "交易Hash"
// @Param address query string false "地址, 显示该地址的资产数量变化"
// @Param coin query string false "资产名, 显示该地址的该资产数量变化"
// @Success 200 {object} Response
// @Router /transactions/{hash} [get]
func (api *API) GetTransaction(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	var transaction model.Transaction
	cond := map[string]interface{}{}
	cond["hash"] = c.Param("hash")
	if result := api.db.Preload("UTXOMsgs").Preload("UTXOMsgs.Inputs").Preload("UTXOMsgs.Outputs").Preload("UTXOMsgs.ContractRequests").Find(&transaction, cond); result.Error != nil {
		api.logger.Error("GetTransaction", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		var blockchain model.BlockChain
		if result := api.db.Last(&blockchain); result.Error != nil {
			api.logger.Error("GetTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		transaction.FillConfirmed(blockchain.BlockNum, c.Query("address"), c.Query("coin"))
		response.Data = transaction
		c.JSON(http.StatusOK, response)
		return
	}

	var mtransaction model.MTransaction
	if result := api.db.Find(&mtransaction, cond); result.Error != nil {
		api.logger.Error("GetMTransaction", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		mtransaction.FillConfirmed(c.Query("address"), c.Query("coin"))
		response.Data = mtransaction
	}

	c.JSON(http.StatusOK, response)
}
