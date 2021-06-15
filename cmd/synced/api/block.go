package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

type BlocksResponse struct {
	PageResponse
	BlockList []model.Block `json:"blocks,omitempty"`
}

func (api *API) GetBlocks(c *gin.Context) {
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
	var blocks []model.Block
	if result := api.db.Order("ID desc").Offset(offset).Limit(request.PageSize).Preload("TxList").Find(&blocks); result.Error != nil {
		api.logger.Error("GetBlocks", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var block model.Block
	if result := api.db.Last(&block); result.Error != nil {
		api.logger.Error("GetBlocks", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := block.Height
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &BlocksResponse{
		PageResponse: PageResponse{
			PageNum:   request.PageNum,
			PageSize:  request.PageSize,
			PageTotal: pageTotal,
			Total:     total,
		},
		BlockList: blocks,
	}
	c.JSON(http.StatusOK, response)
}

func (api *API) GetBlock(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	var block model.Block
	cond := map[string]interface{}{}
	id := c.Param("id")
	if height, err := strconv.ParseInt(id, 10, 64); err != nil {
		cond["hash"] = id
	} else {
		cond["height"] = height
	}
	if result := api.db.Preload("TxList").Find(&block, cond); result.Error != nil {
		api.logger.Error("GetBlock", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else if result.RowsAffected > 0 {
		response.Data = &block
	}
	c.JSON(http.StatusOK, response)
}
