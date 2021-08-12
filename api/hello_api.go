package api

import (
	"gin/conf"
	"gin/model"
	http "gin/utils"
	log "gin/utils"
	"github.com/gin-gonic/gin"
)

var db = conf.GetDB()

func HelloHandler(c *gin.Context) {
	log.Errors(c, "hello")
	http.Ok(c, "hello")
}

func PostHandler(c *gin.Context) {
	user := model.User{
		Name:"helloworld",
	}
	_ = db.Create(&user)

	http.Ok(c, user.ID)
}

