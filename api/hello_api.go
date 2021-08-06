package api

import (
	http "gin/utils"
	log "gin/utils"
	"github.com/gin-gonic/gin"
)

func HelloHandler(c *gin.Context) {
	log.Errors(c, "hello")
	http.Ok(c, "hello")
}

func PostHandler(c *gin.Context) {
	http.Ok(c, "s")
}