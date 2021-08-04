package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
	"gin/utils/log"
)

func LoggerMiddleware() gin.HandlerFunc {
	logger := log.GetLogger()
	return func(c *gin.Context) {
		startTime := time.Now()               // 开始时间
		c.Next()                              // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqUri := c.Request.RequestURI        // 请求路由
		statusCode := c.Writer.Status()       // 状态码
		clientIP := c.Request.Host            // 请求IP
		logger.WithFields(logrus.Fields{      // 日志格式
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"type": "接口调用日志",
		}).Info()
	}
}
