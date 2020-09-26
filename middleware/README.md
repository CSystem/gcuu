# 各种自制 Gin 中间件

## logger.go

logrus 是结构化的，与 golang 标准库完全兼容的 golang 日志打印框架。支持 Filed 和 HOOK 机制，自带 text 和 json 两种日志输出格式。

logrus 自带两种格式 logrus.JSONFormatter{} 和 logrus.TextFormatter{}。

logrus 提供 6 档日志级别，分别是：

* PanicLevel
* FatalLevel - 网站挂了，或者极度不正常
* ErrorLevel - 跟遇到的用户说对不起，可能有bug
* WarnLevel - 记录一下，某事又发生了
* InfoLevel - 提示一切正常
* DebugLevel - 没问题，就看看堆栈

### middleware 的使用方法

```golang
r := gin.Default()                  // 默认启动方式，包含 Logger、Recovery 中间件
r.Use(middleware.LoggerToFile(app)) // 初始化日
```

```golang
log.Debug("Useful debugging information.")
log.Info("Something noteworthy happened!")
log.Warn("You should probably take a look at this.")
log.Error("Something failed but I'm not quitting.")
log.Fatal("Bye.") // 随后会触发os.Exit(1)
log.Panic("I'm bailing.") // 随后会触发panic()

```

```golang
[GIN] 2018/05/30 - 19:21:17 | 404 |     114.656µs |             ::1 |  GET     /comment/view/99999
[GIN] 2018/05/30 - 19:21:20 | 200 |     1.82468ms |             ::1 |  GET     /comment/view/00001
```

### 参考

* [How to write request log and error log in a separate manner](https://github.com/gin-gonic/gin/issues/1376)
* [gin-glog](https://github.com/szuecs/gin-glog/blob/master/ginglog.go)
* [os/example_test.go](https://golang.org/src/os/example_test.go)
* [Gin框架 - 使用 Logrus 进行日志记录](https://juejin.im/post/5d3932bde51d454f73356e2d)
* [gin-logrus](https://github.com/toorop/gin-logrus/blob/master/logger.go)
* [Golang之使用Logrus](https://o-my-chenjian.com/2017/09/19/Using-Logrus-With-Golang/)
* [golang logrus (1) 简单使用](https://juejin.im/post/5b20789f518825137478a697)
