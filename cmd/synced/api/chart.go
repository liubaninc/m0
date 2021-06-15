package api

import (
	"encoding/hex"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (api *API) GetSearch(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	content := c.Param("content")
	if _, err := strconv.ParseInt(content, 10, 64); err == nil {
		response.Data = "blocks/" + content
		c.JSON(http.StatusOK, response)
		return
	}
	if _, err := sdk.AccAddressFromBech32(content); err == nil {
		response.Data = "addresses/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	if _, err := hex.DecodeString(content); err == nil {
		var block model.Block
		if result := api.db.Find(&block, "hash = ?", content); result.RowsAffected == 1 {
			response.Data = "blocks/" + content
			c.JSON(http.StatusOK, response)
			return
		}

		var tx model.Transaction
		if result := api.db.Find(&tx, "hash = ?", content); result.RowsAffected == 1 {
			response.Data = "transactions/" + content
			c.JSON(http.StatusOK, response)
			return
		}
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if result := api.db.Last(&blockchain); result.Error != nil {
		api.logger.Error("GetSearch", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	cc := api.client.WithHeight(blockchain.BlockNum)
	if has, _ := cc.GetContract(content); has != nil {
		response.Data = "contracts/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	if has, _ := cc.GetToken(content); has != nil {
		response.Data = "assets/" + content
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

type LineChart struct {
	Time   string `json:"time"`
	Number int64  `json:"number"`
}

type ChartsResponse struct {
	PageResponse
	BlockList       []*LineChart `json:"blocks"`
	TransactionList []*LineChart `json:"txs"`
	MessageList     []*LineChart `json:"msgs"`
	AssetList       []*LineChart `json:"assets"`
	ContractList    []*LineChart `json:"contracts"`
}

func (api *API) GetCharts(c *gin.Context) {
	var request PageRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.BindQuery(&request); err != nil {
		response.Code = RequestCode
		response.Msg = err.Error()
	} else {
		if request.PageNum < 1 {
			request.PageNum = 1
		}
		if request.PageSize < 1 {
			request.PageSize = 10
		}
		offset := (request.PageNum - 1) * request.PageSize

		var blockChainChart model.BlockChainChart
		if result := api.db.Last(&blockChainChart); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		var blocks []*LineChart
		if result := api.db.Model(&model.BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "block_num as number").Find(&blocks); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
		var transactions []*LineChart
		if result := api.db.Model(&model.BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "tx_num as number").Find(&transactions); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
		var messages []*LineChart
		if result := api.db.Model(&model.BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "msg_num as number").Find(&messages); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
		var assets []*LineChart
		if result := api.db.Model(&model.BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "asset_num as number").Find(&assets); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}
		var contracts []*LineChart
		if result := api.db.Model(&model.BlockChainChart{}).Order("ID desc").Offset(offset).Limit(request.PageSize).Select("time", "contract_num as number").Find(&contracts); result.Error != nil {
			api.logger.Error("GetCharts", "error", result.Error)
			response.Code = ExecuteCode
			response.Msg = result.Error.Error()
			c.JSON(http.StatusOK, response)
			return
		}

		total := int64(blockChainChart.ID)
		pageTotal := total / int64(request.PageSize)
		if int64(blockChainChart.ID)%int64(request.PageSize) != 0 {
			pageTotal += 1
		}
		response.Data = &ChartsResponse{
			PageResponse: PageResponse{
				Total:     total,
				PageSize:  request.PageSize,
				PageNum:   request.PageNum,
				PageTotal: pageTotal,
			},
			BlockList:       blocks,
			TransactionList: transactions,
			MessageList:     messages,
			AssetList:       assets,
			ContractList:    contracts,
		}
	}
	c.JSON(http.StatusOK, response)
}

type TPSChart struct {
	Time   string `json:"time"`
	Number int64  `json:"number"`
}

type TPSResponse struct {
	Transaction *TPSChart `json:"tx"`
	Message     *TPSChart `json:"msg"`
}

func (api *API) GetMaxTPS(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	var tx TPSChart
	if result := api.db.Model(&model.BlockChainTPSChart{}).Select("time, max(tx_num) as number").Find(&tx); result.Error != nil {
		api.logger.Error("GetMaxTPS", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var msg TPSChart
	if result := api.db.Model(&model.BlockChainTPSChart{}).Select("time, max(msg_num) as number").Find(&msg); result.Error != nil {
		api.logger.Error("GetMaxTPS", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	response.Data = &TPSResponse{
		Transaction: &tx,
		Message:     &msg,
	}
	c.JSON(http.StatusOK, response)
}
