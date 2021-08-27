package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadRouters(e *gin.Engine) {
	HandleErrorRoutes(e)
	LoadUser(e)
}

