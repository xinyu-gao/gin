package routers

import (
	"gin/api"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/hello", api.HelloHandler)
		userGroup.POST("/post", api.PostHandler)
		userGroup.POST("/login", login)
	}
}

type User struct {
	Username string `json:"username"`
	Passwd   string `json:"passwd"`
}

func login(c *gin.Context) {
	f := c.Param("aaa")
	json := User{}
	if err := c.BindJSON(&json); err != nil {
		return
	}
	username := json.Username
	passwd := json.Passwd
	c.String(http.StatusOK, "Hello world "+username+"_"+passwd+f)
}
