package log

import (
	"gin/conf"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"time"
)

var logger = logrus.New()
func init() {
	filePath := conf.InitConfigure().GetLogFilePath()
	logger.SetLevel(logrus.DebugLevel)           //设置日志级别
	logger.SetFormatter(&logrus.JSONFormatter{}) //设置日志格式

	logWriterInfo, _ := rotatelogs.New(
		filePath+"info-%Y%m%d.log",                // 分割后的文件名称
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)

	logWriterError, _ := rotatelogs.New(
		filePath+"error-%Y%m%d.log",               // 分割后的文件名称
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 设置最大保存时间(7天)
		rotatelogs.WithRotationTime(24*time.Hour), // 设置日志切割时间间隔(1天)
	)

	writeMap := lfshook.WriterMap{
		logrus.TraceLevel: logWriterInfo,
		logrus.DebugLevel: logWriterInfo,
		logrus.InfoLevel:  logWriterInfo,
		logrus.WarnLevel:  logWriterError,
		logrus.ErrorLevel: logWriterError,
		logrus.FatalLevel: logWriterError,
		logrus.PanicLevel: logWriterError,
	}

	logger.AddHook(lfshook.NewHook(writeMap, &logrus.JSONFormatter{
		TimestampFormat: "2000-01-01 12:0:00",
	}))

}
func GetLogger() *logrus.Logger{
	return logger;
}
func field(c *gin.Context) *logrus.Entry {
	reqMethod := c.Request.Method  // 请求方式
	reqUri := c.Request.RequestURI // 请求路由
	clientIP := c.Request.Host     // 请求IP
	return logger.WithFields(logrus.Fields{ // 日志格式
		"client_ip":  clientIP,
		"req_method": reqMethod,
		"req_uri":    reqUri,
		"type": "业务日志",
	})
}
func Info(c *gin.Context, msg string){
	field(c).Info(msg)
}

func Error(c *gin.Context, msg string){
	field(c).Error(msg)
}
func Warn(c *gin.Context, msg string){
	field(c).Warn(msg)
}
func Debug(c *gin.Context, msg string){
	field(c).Debug(msg)
}
