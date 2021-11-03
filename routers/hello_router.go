package routers

import (
	"gin/apis"
	"github.com/gin-gonic/gin"
)

func LoadHello(e *gin.Engine) {
	userGroup := e.Group("/hello")
	{
		userGroup.GET("/hello", apis.HelloHandler)
		userGroup.POST("/post", apis.PostHandler)
	}
}

