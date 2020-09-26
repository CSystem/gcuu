package initializer

import (
	"time"
	"github.com/CSystem/gcuu/initializer/config"
	"github.com/CSystem/gcuu/initializer/database"
	"github.com/CSystem/gcuu/mlog"
	"github.com/CSystem/gcuu/mredis"
	gjwt "github.com/appleboy/gin-jwt/v2"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
)

type Application struct {
	Conf           *config.Configuration
	DB             *gorm.DB
	AuthMiddleware *gjwt.GinJWTMiddleware
	Cache          *cache.Cache
}

func InitApp() *Application {
	conf := config.Init()                                                                  // 初始化全局配置文件
	db := database.InitConnection(conf)                                                    // 初始化 MySQL 连接
	mlog.Setup(conf.Server.Mode)                                                           // 初始化日志
	mredis.Setup(conf)                                                                     // 初始化 Redis Pool 和 RedLock
	authMiddleware := InitAuthMiddleware(conf)                                             // 初始化 JWT Middleware for Gin framework
	c := cache.New(5*time.Minute, 10*time.Minute)                                          // 初始化缓存，参数1是默认过期时间，参数2是每隔多久清理过期缓存

	return &Application{
		Conf:           conf,
		DB:             db,
		AuthMiddleware: authMiddleware,
		Cache:          c,
	}
}

func (app *Application) Close() {
	mlog.Logger.Close()
	mredis.Pool.Close()
	app.DB.Close()
}
