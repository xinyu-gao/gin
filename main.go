package main

import (
	"gin/conf"
	m "gin/middlewares"
	"github.com/gin-gonic/gin"
)

var mode = gin.DebugMode // 开发模式

func main() {
	gin.SetMode(mode)
	r := gin.New()
	r.Use(
		gin.Recovery(),
		m.LoggerMiddleware(),
		m.Cors(),
	)
	conf.LoadRouters(r)
	_ = r.Run(conf.GetServePort())

}
