package routers

import (
	http "gin/utils"
	"github.com/gin-gonic/gin"
)

// 自定义 404 / method not allowed 返回结果
func handleErrorRoutes(e *gin.Engine){
	e.HandleMethodNotAllowed = true
	e.NoMethod(func(c *gin.Context) {
		http.MethodNotAllowed(c)
	})
	e.NoRoute(func(c *gin.Context) {
		http.NotFound(c)
	})
}