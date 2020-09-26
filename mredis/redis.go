package mredis

import (
	"github.com/CSystem/gcuu/initializer/config"
	"github.com/CSystem/gcuu/initializer/redigo"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

var (
	Pool    *meRedis
	RedLock *redsync.Redsync
)

type meRedis struct {
	*redis.Pool
}

func (cr *meRedis) WithConn(fn func(c redis.Conn) error) error {
	c := Pool.Get()
	defer c.Close()

	return fn(c)
}

func Setup(conf *config.Configuration) {
	Pool = &meRedis{redigo.InitConnection(conf)} // 初始化 RedisPool 连接池
	RedLock = redigo.InitRedSyncConnection(Pool) // 初始化 RedSync 实例
}
