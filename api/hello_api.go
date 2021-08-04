package api

import (
	http "gin/utils"
	"gin/utils/log"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	log.Error(c, "hello")
	http.Ok(c, "hello")
}

func PostHandler(c *gin.Context) {
	http.Ok(c, "s")
}