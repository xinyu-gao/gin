package routers

import (
	"gin/apis"
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/find", apis.GetUserByIDHandler)
		//userGroup.POST("/post", apis.PostHandler)
	}
}


