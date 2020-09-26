# catredis

全局通用的 Redis 和 Redsync 变量。

## 用法

Redis 用法如下：

```golang
c := catredis.Pool.Get()
defer c.Close()

_, err := c.Do("ZADD", consts.GoldTopRedisKey, -amount, userID)

return err
```

Redsync 分布式锁的用法如下：

```golang
mutex := catredis.RedLock.NewMutex(fmt.Sprintf("%s%d", consts.RedLockCatMovePrefix, userID))
if err := mutex.Lock(); err != nil {
  return err
}
defer mutex.Unlock()
```
