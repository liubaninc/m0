package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

func (api *API) GetBlockChain(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	var blockchain model.BlockChain
	if result := api.db.First(&blockchain); result.Error != nil {
		api.logger.Error("GetBlockChain", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
	} else {
		response.Data = blockchain
	}
	c.JSON(http.StatusOK, response)
}
