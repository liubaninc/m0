package api

import (
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"net/http"
)

type EventsResponse struct {
	PageResponse
	Items []*model.Events `json:"items"`
}

func (api *API) GetPeerEvents(c *gin.Context) {
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
	cond["route"] = "peer"
	if tp := c.Param("action"); len(tp) != 0 {
		cond["type"] = tp
	}

	var events []*model.Events
	if result := api.db.Order("height").Offset(offset).Limit(request.PageSize).Where(cond).Find(&events); result.Error != nil {
		api.logger.Error("GetEvents", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	total := int64(0)
	api.db.Model(&model.Events{}).Where(cond).Count(&total)

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &EventsResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		Items: events,
	}
	c.JSON(http.StatusOK, response)
}

func (api *API) GetValidatorEvents(c *gin.Context) {
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
	cond["route"] = "validator"
	if tp := c.Param("action"); len(tp) != 0 {
		cond["type"] = tp
	}

	var events []*model.Events
	if result := api.db.Order("height").Offset(offset).Limit(request.PageSize).Where(cond).Find(&events); result.Error != nil {
		api.logger.Error("GetEvents", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	total := int64(0)
	api.db.Model(&model.Events{}).Where(cond).Count(&total)

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &EventsResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		Items: events,
	}
	c.JSON(http.StatusOK, response)
}

func (api *API) GetCertEvents(c *gin.Context) {
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
	cond["route"] = "pki"
	if tp := c.Param("action"); len(tp) != 0 {
		cond["type"] = tp
	}

	var events []*model.Events
	if result := api.db.Order("height").Offset(offset).Limit(request.PageSize).Where(cond).Find(&events); result.Error != nil {
		api.logger.Error("GetEvents", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	total := int64(0)
	api.db.Model(&model.Events{}).Where(cond).Count(&total)

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &EventsResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		Items: events,
	}
	c.JSON(http.StatusOK, response)
}

func (api *API) GetAccountEvents(c *gin.Context) {
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
	cond["route"] = "permission"
	if tp := c.Param("action"); len(tp) != 0 {
		cond["type"] = tp
	}

	var events []*model.Events
	if result := api.db.Order("height").Offset(offset).Limit(request.PageSize).Where(cond).Find(&events); result.Error != nil {
		api.logger.Error("GetEvents", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	total := int64(0)
	api.db.Model(&model.Events{}).Where(cond).Count(&total)

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &EventsResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		Items: events,
	}
	c.JSON(http.StatusOK, response)
}

func (api *API) GetContractEvents(c *gin.Context) {
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
	cond["route"] = "wasm"
	if tp := c.Param("action"); len(tp) != 0 {
		cond["type"] = tp
	}

	var events []*model.Events
	if result := api.db.Order("height").Offset(offset).Limit(request.PageSize).Where(cond).Find(&events); result.Error != nil {
		api.logger.Error("GetEvents", "error", result.Error)
		response.Code = ExecuteCode
		response.Msg = result.Error.Error()
		c.JSON(http.StatusOK, response)
		return
	}
	total := int64(0)
	api.db.Model(&model.Events{}).Where(cond).Count(&total)

	pageTotal := total / int64(request.PageSize)
	if total%int64(request.PageSize) != 0 {
		pageTotal += 1
	}
	response.Data = &EventsResponse{
		PageResponse: PageResponse{
			Total:     total,
			PageSize:  request.PageSize,
			PageNum:   request.PageNum,
			PageTotal: pageTotal,
		},
		Items: events,
	}
	c.JSON(http.StatusOK, response)
}
