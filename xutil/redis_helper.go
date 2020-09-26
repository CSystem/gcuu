package xutil

import (
	"fmt"
	"github.com/CSystem/gcuu/mredis"
	"github.com/gomodule/redigo/redis"
)

// 快捷获取今天的开始时间 https://github.com/jinzhu/now
// 比较两个时间点 https://juejin.im/post/5ae32a8651882567105f7dd3
// 秒、毫秒、纳秒时间戳输出 https://blog.csdn.net/mirage003/article/details/80822608

func GetTokenFromRedis(userId uint,appId uint) (token string, err error) {
	c := mredis.Pool.Get()
	defer c.Close()

	key := fmt.Sprintf("jwt-token-%d-%d", userId,appId)
	exists, err := redis.Bool(c.Do("EXISTS", key))
	if !exists {
		return "", err
	}
	token, err = redis.String(c.Do("GET", key))

	return token, nil
}

func SetTokenToRedis(userId uint,appId uint,token string) (err error)  {
	c := mredis.Pool.Get()
	defer c.Close()

	key := fmt.Sprintf("jwt-token-%d-%d", userId,appId)
	if _, err = c.Do("SET", key,token, "EX", 24*3600); err != nil {
		return err
	}

	return nil
}

func DeleteTokenInRedis(userId uint,appId uint) (err error)  {
	c := mredis.Pool.Get()
	defer c.Close()

	key := fmt.Sprintf("jwt-token-%d-%d", userId,appId)
	if _, err = c.Do("DEL", key); err != nil {
		return err
	}
	return nil
}