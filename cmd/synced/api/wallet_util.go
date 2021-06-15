package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/liubaninc/m0/cmd/synced/model"
	"gorm.io/gorm"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	TokenKey       = "user secret"
	headerTokenKey = "Authorization"
	userIDKey      = "user_id"
	userNameKey    = "user_name"
	expDuration    = 5 * time.Minute
)

func createToken(key string, m map[string]interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	for index, val := range m {
		claims[index] = val
	}
	token.Claims = claims
	return token.SignedString([]byte(key))
}

func parseToken(tokenString string, key string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//c, _ := token.Claims.(jwt.MapClaims)
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		m := make(map[string]interface{})
		for index, val := range claims {
			m[index] = val
		}
		return m, nil
	}
	return nil, fmt.Errorf("invalid token")
}

// AuthorizeToken login jwt
type authorizeToken struct {
	UserID uint  `json:"userID"`
	Exp    int64 `json:"exp"`
	Iat    int64 `json:"iat"`
}

func newAuthorizeToken(UserID uint, exp int64) *authorizeToken {
	token := &authorizeToken{}
	token.UserID = UserID
	now := time.Now()
	token.Iat = now.Unix()
	if exp == 0 {
		token.Exp = now.Add(expDuration).Unix()
	} else {
		token.Exp = now.Add(time.Duration(exp) * time.Minute).Unix()
	}
	return token
}

func newFromJWT(jwttoken string) (*authorizeToken, error) {
	info, err := parseToken(jwttoken, TokenKey)
	if err != nil {
		return nil, err
	}
	bts, err := json.Marshal(info)
	if err != nil {
		return nil, err
	}
	token := &authorizeToken{}
	err = json.Unmarshal(bts, token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (a *authorizeToken) toJWT() (string, error) {
	bts, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	info := make(map[string]interface{})
	err = json.Unmarshal(bts, &info)
	if err != nil {
		panic(err)
	}
	return createToken(TokenKey, info)
}

func (api *API) authorize(c *gin.Context) {
	response := &Response{
		Code: OKCode,
		Msg:  OKMsg,
	}
	token := c.GetHeader(headerTokenKey)
	auth, err := newFromJWT(token)
	if err != nil {
		response.Code = AuthCode
		response.Msg = "认证失败"
		api.logger.Error("auth failed", "error", err)
		c.Header(headerTokenKey, "")
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}

	var usr model.User
	if result := api.db.Where(model.User{
		Model: gorm.Model{
			ID: auth.UserID,
		},
	}).First(&usr); result.Error != nil {
		response.Code = AuthCode
		response.Msg = "认证失败"
		api.logger.Error("auth failed", "error", fmt.Sprintf("user %d not exist", auth.UserID))
		c.Header(headerTokenKey, "")
		c.JSON(http.StatusOK, response)
		c.Abort()
		return
	}
	session := sessions.Default(c)
	session.Set(userIDKey, usr.ID)
	session.Set(userNameKey, usr.Name)
	session.Save()

	now := time.Now()
	auth.Exp = now.Add(time.Duration(auth.Exp-auth.Iat) * time.Second).Unix()
	auth.Iat = now.Unix()
	if token, err := auth.toJWT(); err == nil {
		c.Header(headerTokenKey, token)
	}
	c.Next()
}

type CaptchaResponse struct {
	CaptchaId string `json:"captchaId"` //验证码Id
	ImageUrl  string `json:"imageUrl"`  //验证码图片url
}

// @生成验证码
// @Summary 生成验证码
// @Description
// @Tags user
// @Accept  json
// @Produce json
// @Success 200 {string} json
// @Router /captcha [get]
func GetCaptcha(c *gin.Context) {
	length := captcha.DefaultLen
	captchaId := captcha.NewLen(length)
	var captcha CaptchaResponse
	captcha.CaptchaId = captchaId
	captcha.ImageUrl = "/captcha/" + captchaId + ".png"
	c.JSON(http.StatusOK, captcha)
}

// @获取验证码
// @Summary 获取验证码
// @Description
// @Tags user
// @Accept  json
// @Produce json
// @Param captchaId path string true "验证码ID"
// @Param reload query string false "重新生成验证码"
// @Success 200 {string} json
// @Router /captcha/{captchaId} [get]
func GetCaptchaPNG(c *gin.Context) {
	ServeHTTP(c.Writer, c.Request)
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dir, file := path.Split(r.URL.Path)
	ext := path.Ext(file)
	id := file[:len(file)-len(ext)]
	if ext == "" || id == "" {
		http.NotFound(w, r)
		return
	}
	if r.FormValue("reload") != "" {
		captcha.Reload(id)
	}
	lang := strings.ToLower(r.FormValue("lang"))
	download := path.Base(dir) == "download"
	if Serve(w, r, id, ext, lang, download, captcha.StdWidth, captcha.StdHeight) == captcha.ErrNotFound {
		http.NotFound(w, r)
	}
}

func Serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}
