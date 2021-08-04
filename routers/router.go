package routers

import (
	"github.com/gin-gonic/gin"
)

func LoadRouters(e *gin.Engine) {
	handleErrorRoutes(e)
	loadUser(e)
}

