package main

import (
	"gin/conf"
	m "gin/middlewares"
	"gin/routers"
	"github.com/gin-gonic/gin"
)

var mode = gin.DebugMode // 开发模式

func main() {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(m.LoggerMiddleware(),
		gin.Recovery(),
		m.Cors(),
	)
	routers.LoadRouters(r)
	_ = r.Run(conf.GetServePort())
}
