package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

type PeersResponse struct {
	PageResponse
	PeerList []*model.Peer `json:"peers"`
}

func (api *API) GetPeers(c *gin.Context) {
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
	var peers []*model.Peer
	if result := api.db.Order("Time desc").Offset(offset).Limit(request.PageSize).Find(&peers); result.Error != nil {
		api.logger.Error("GetPeers", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	var blockchain model.BlockChain
	if result := api.db.Last(&blockchain); result.Error != nil {
		api.logger.Error("GetPeers", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}

	total := blockchain.PeerNum
	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &PeersResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		PeerList: peers,
	}
	c.JSON(http.StatusOK, response)
}
