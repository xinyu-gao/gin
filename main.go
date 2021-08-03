package main

import (
	"gin/conf"
	"gin/routers"
	"github.com/gin-gonic/gin"
)

var mode = gin.DebugMode	// 开发模式

func main() {
	gin.SetMode(mode)
	var config = conf.InitConfigure()
	port := config.GetServePort()

	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())
	routers.LoadRouters(r)
	_ = r.Run(port)
}
