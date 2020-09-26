package xutil

import (
	gjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetPayloadFromGin(c *gin.Context) (userId uint,appId uint,t int64) {

	claims := gjwt.ExtractClaims(c)
	if len(claims) == 0 {
		return 0,0,0
	}

	//fmt.Printf("%v",claims)
	userId = uint(claims["user_id"].(float64))
	appId = uint(claims["app_id"].(float64))
	t = int64(claims["t"].(float64))

	return userId,appId,t
}

func GetPayloadFromToken(token *jwt.Token) (userId uint,appId uint,t int64) {

	claims := gjwt.ExtractClaimsFromToken(token)
	userId = uint(claims["user_id"].(float64))
	appId = uint(claims["app_id"].(float64))
	t = int64(claims["t"].(float64))

	return userId,appId,t
}

func GetToken(c *gin.Context) string {
	token := gjwt.GetToken(c)
	return token
}