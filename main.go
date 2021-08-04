package main

import (
	"gin/conf"
	"gin/middlewares"
	"gin/routers"
	"github.com/gin-gonic/gin"
)

var mode = gin.DebugMode	// 开发模式
var config = conf.InitConfigure()

func main() {
	gin.SetMode(mode)
	r := gin.Default()
	r.Use(middlewares.LoggerMiddleware())
	//r.Use(gin.Logger(), gin.Recovery())
	routers.LoadRouters(r)
	_ = r.Run(config.GetServePort())
}
