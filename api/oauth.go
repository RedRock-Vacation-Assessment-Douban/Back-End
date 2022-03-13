package api

import (
	"douban/model"
	"douban/service"
	"douban/tool"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"time"
)

func Oauth(ctx *gin.Context) {
	var err error
	// 获取 code
	var code = ctx.Query("code")
	// 通过 code, 获取 token
	var tokenAuthUrl = service.GetTokenAuthUrl(code)
	var token *model.Token
	if token, err = service.GetToken(tokenAuthUrl); err != nil {
		tool.RespInternalError(ctx)
		return
	}

	// 通过token，获取用户信息
	var userInfo map[string]interface{}
	userInfo, err = service.GetUserInfo(token)
	if err != nil {
		tool.RespInternalError(ctx)
		return
	}
	user := userInfo["login"].(string)
	c := model.MyClaims{
		Username: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 6000000,
			Issuer:    "YuanXinHao",
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		tool.RespInternalError(ctx)
	}
	tool.RespSuccessfulWithTwoDate(ctx, user, s)
}
