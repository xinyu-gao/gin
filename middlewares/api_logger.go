package middlewares

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}
func LoggerMiddleware() gin.HandlerFunc {
	logger := utils.GetLogger()
	return func(c *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		startTime := time.Now()               // 开始时间
		c.Next()                              // 处理请求
		endTime := time.Now()                 // 结束时间
		latencyTime := endTime.Sub(startTime) // 执行时间
		reqMethod := c.Request.Method         // 请求方式
		reqUri := c.Request.RequestURI        // 请求路由
		clientIP := c.Request.Host            // 请求IP
		result := blw.body.String()			  // 响应内容
		logger.WithFields(logrus.Fields{ // 日志格式
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"res_data":     result,
			"type":         "接口调用",
		}).Info()
	}
}
