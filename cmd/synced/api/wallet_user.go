package api

import (
	"fmt"
	"net/http"

	"github.com/alexandrevicenzi/unchained"
	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
)

// @用户注册
// @Summary 注册新用户
// @Description
// @Tags user
// @Accept  json
// @Produce json
// @Param user body UserRegisterRequest true "用户信息"
// @Success 200 {object} Response
// @Router /user/register [post]
func (api *API) UserRegister(c *gin.Context) {
	var request UserRegisterRequest
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		response.Code = RequestCode
		response.Msg = ERROR_REQ
		response.Detail = err.Error()
		api.logger.Error(c.Request.URL.Path, "error", err.Error())
		c.JSON(http.StatusOK, response)
		return
	}

	if !captcha.VerifyString(request.CaptchaId, request.Captcha) {
		fmt.Println(request.CaptchaId, request.Captcha)
		response.Code = ExecuteCode
		response.Msg = ERROR_CAPTCHA
		response.Detail = response.Msg
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	password, err := unchained.MakePassword(request.Password, unchained.GetRandomString(12), "default")
	if err != nil {
		panic(err)
	}
	usr := &model.User{
		Name:     request.Name,
		Password: password,
		Nick:     request.Nick,
		Mobile:   request.Mobile,
		Email:    request.Email,
	}
	if result := api.db.Create(usr); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_USER_EXIST
		response.Detail = result.Error.Error()
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Data = &UserResponse{
		Name:   usr.Name,
		Nick:   usr.Nick,
		Email:  usr.Email,
		Mobile: usr.Mobile,
	}
	auth := newAuthorizeToken(usr.ID, 0)
	token, err := auth.toJWT()
	if err != nil {
		panic(err)
	}
	c.Header(headerTokenKey, token)
	c.JSON(http.StatusOK, response)
}

// @用户登陆
// @Summary 用户登陆
// @Description
// @Tags user
// @Accept  json
// @Produce json
// @Param user body UserLoginRequest true "登陆信息"
// @Success 200 {object} Response
// @Router /user/login [post]
func (api *API) UserLogin(c *gin.Context) {
	request := UserLoginRequest{}
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

	if !captcha.VerifyString(request.CaptchaId, request.Captcha) {
		response.Code = RequestCode
		response.Msg = ERROR_CAPTCHA
		response.Detail = response.Msg
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	var user model.User
	if result := api.db.Where(&model.User{
		Name: request.Name,
	}).First(&user); result.Error != nil {
		response.Code = ExecuteCode
		response.Msg = ERROR_USER_NO
		response.Detail = fmt.Sprintf("user %s not exist", request.Name)
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	if valid, err := unchained.CheckPassword(request.Password, user.Password); err != nil || !valid {
		response.Code = RequestCode
		response.Msg = ERROR_PASSWORD
		response.Detail = response.Msg
		api.logger.Error(c.Request.URL.Path, "error", response.Detail)
		c.JSON(http.StatusOK, response)
		return
	}

	response.Data = &UserResponse{
		Name:   user.Name,
		Nick:   user.Nick,
		Email:  user.Email,
		Mobile: user.Mobile,
	}
	auth := newAuthorizeToken(user.ID, request.ExpDuration)
	token, err := auth.toJWT()
	if err != nil {
		panic(err)
	}
	c.Header(headerTokenKey, token)
	c.JSON(http.StatusOK, response)
}

// @用户退出
// @Summary 用户退出
// @Description
// @Tags user
// @Accept  json
// @Produce json
// @Success 200 {object} Response
// @Security ApiKeyAuth
// @Router /user/logout [post]
func (api *API) UserLogout(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}

	c.Header(headerTokenKey, "")
	c.JSON(http.StatusOK, response)
}
