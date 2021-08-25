package conf

import (
	"gin/routers"
	"github.com/gin-gonic/gin"
)

func LoadRouters(e *gin.Engine) {
	routers.HandleErrorRoutes(e)
	routers.LoadUser(e)
}

