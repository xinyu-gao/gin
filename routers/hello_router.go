package routers

import (
	"gin/api"
	"github.com/gin-gonic/gin"
)

func loadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/hello", api.HelloHandler)
		userGroup.POST("/post", api.PostHandler)
	}
}
