package routers

import (
	"gin/apis"
	http "gin/utils"
	"github.com/gin-gonic/gin"
)

func LoadUser(e *gin.Engine) {
	userGroup := e.Group("/user")
	{
		userGroup.GET("/hello", apis.HelloHandler)
		userGroup.POST("/post", apis.PostHandler)
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
