package redigo

import (
	"fmt"
	"time"

	"github.com/CSystem/gcuu/initializer/config"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
)

// https://docs.objectrocket.com/redis_go_examples.html
// InitConnection - opening a redis connection.
func InitConnection(conf *config.Configuration) *redis.Pool {
	var err error

	redisURI := fmt.Sprintf("redis://%s@%s:%d", conf.Redis.Password, conf.Redis.Host, conf.Redis.Port)
	pool := &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.DialURL(redisURI) },
	}

	_, err = pool.Dial()

	if err != nil {
		fmt.Printf("redis connect error %v\n", err)
		panic("failed to connect redis")
	}

	return pool
}

func InitRedSyncConnection(pool redsync.Pool) *redsync.Redsync {
	return redsync.New([]redsync.Pool{pool})
}
