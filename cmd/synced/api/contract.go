package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

type ContractsResponse struct {
	PageResponse
	ContractList []*model.Contract `json:"contracts,omitempty"`
}

func (api *API) GetContracts(c *gin.Context) {
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
	var contracts []*model.Contract
	if result := api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&contracts); result.Error != nil {
		api.logger.Error("GetContracts", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if result := api.db.First(&blockchain); result.Error != nil {
		api.logger.Error("GetContracts", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.ContractNum
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

func (api *API) GetContract(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	cond := map[string]interface{}{}
	cond["name"] = c.Param("name")
	var contract model.Contract
	if result := api.db.Find(&contract, cond); result.Error != nil {
		api.logger.Error("GetContract", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = contract
	}
	c.JSON(http.StatusOK, response)
}

type ContractTransactionsRequest struct {
	PageRequest
	Invoke bool `form:"invoke" json:"invoke" xml:"invoke"`
}

func (api *API) GetContractTransactions(c *gin.Context) {
	var request ContractTransactionsRequest
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

	if request.Invoke {
		var upgrades []model.ContractUpgrade
		if result := api.db.Where("name = ?", c.Param("name")).Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&upgrades); result.Error != nil {
			api.logger.Error("GetContractTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		var contract model.Contract
		if result := api.db.Where("name = ?", c.Param("name")).Last(&contract); result.Error != nil {
			api.logger.Error("GetContractTransactions", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		total := contract.Deploy
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
			TransactionList: func() []*model.Transaction {
				var txs []*model.Transaction
				for _, upgrade := range upgrades {
					txs = append(txs, &model.Transaction{
						Hash:    upgrade.Hash,
						Time:    upgrade.Time,
						Version: upgrade.Version,
					})
				}
				return txs
			}(),
		}
		c.JSON(http.StatusOK, response)
		return
	}
	var invokes []model.ContractInvoke
	if result := api.db.Where("name = ?", c.Param("name")).Order("ID desc").Offset(offset).Limit(request.PageSize).Find(&invokes); result.Error != nil {
		api.logger.Error("GetContractTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var contract model.Contract
	if result := api.db.Where("name = ?", c.Param("name")).Last(&contract); result.Error != nil {
		api.logger.Error("GetContractTransactions", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := contract.Invoke
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
		TransactionList: func() []*model.Transaction {
			var txs []*model.Transaction
			for _, invoke := range invokes {
				txs = append(txs, &model.Transaction{
					Hash:    invoke.Hash,
					Time:    invoke.Time,
					Version: invoke.Version,
					Size:    invoke.Size,
					Height:  invoke.Height,
				})
			}
			return txs
		}(),
	}
	c.JSON(http.StatusOK, response)
	return
}
