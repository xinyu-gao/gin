package routers

import (
	"gin/handler"
	"github.com/gin-gonic/gin"
)

func loadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/hello", handler.HelloHandler)
		userGroup.POST("/post", handler.PostHandler)
	}
}
