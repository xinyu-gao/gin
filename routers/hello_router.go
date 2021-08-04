package routers

import (
	"gin/api"
	http "gin/utils"
	"github.com/gin-gonic/gin"
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
		http.Error(c, "json 解析错误")
		return
	}
	username := json.Username
	passwd := json.Passwd
	http.Ok(c, "Hello world "+username+"_"+passwd+f)
}
