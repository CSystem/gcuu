package initializer

import (
	"github.com/CSystem/gcuu/initializer/config"
	"github.com/CSystem/gcuu/initializer/response"
	gjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func InitAuthMiddleware(conf *config.Configuration) *gjwt.GinJWTMiddleware {
	// the jwt middleware
	authMiddleware, err := gjwt.New(&gjwt.GinJWTMiddleware{
		Realm:         conf.JWT.Realm,       // 必要项，显示给用户看的域
		Key:           []byte(conf.JWT.Key), // 用来进行签名的密钥，就是加盐用的
		Timeout:       time.Hour * 24 * 30,  // JWT 的有效时间
		MaxRefresh:    time.Hour * 24 * 30,  // 最长的刷新时间，用来给客户端自己刷新 token 用的
		PayloadFunc:   payloadFunc,          // 往 Payload 中写入数据
		Unauthorized:  unAuthFunc,           // 自定义鉴权失败返回数据结构
		Authorizator:  authorizator,		 // 自定义验证token
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

func payloadFunc(data interface{}) gjwt.MapClaims {
	if v, ok := data.(*response.SignIn); ok {
		return gjwt.MapClaims{
			"user_id": v.ID,
			"app_id" : v.AppId,
			"t" : v.T,
		}
	}

	return gjwt.MapClaims{}
}

// unAuthFunc - 认证不成功的处理函数
func unAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"status": "fail",
		"code":  http.StatusUnauthorized,
		"data": gin.H{"error": message},
	})
}

func authorizator(data interface{}, c *gin.Context) bool {
	return true
}
