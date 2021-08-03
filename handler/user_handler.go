package handler

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	http_result.Ok(c, "hello")
}

func PostHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
		"msg":    "r1.postHandler",
	})
}