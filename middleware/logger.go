package middleware

import (
	"time"

	"github.com/CSystem/gcuu/mlog"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Logger gin 日志输出到文件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()               // 开始时间
		c.Next()                              // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqURI := c.Request.RequestURI        // 请求路由
		statusCode := c.Writer.Status()       // 状态码
		clientIP := c.ClientIP()              // 请求IP

		// 日志格式
		entry := mlog.Logger.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime.String(),
			"client_ip":    clientIP,
			"method":       reqMethod,
			"path":         reqURI,
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String()) // Append error field if this is an erroneous request.
		} else {
			entry.Info()
		}
	}
}
