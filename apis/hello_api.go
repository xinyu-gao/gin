package apis

import (
	"gin/confs"
	"gin/models"
	http "gin/utils"
	"github.com/gin-gonic/gin"
)

var (
	db = confs.GetDB()
	redis = confs.GetRedis()
	ctx = confs.GetCtx()
)


func HelloHandler(c *gin.Context) {
	result := redis.Get(ctx, "id").Val()
	http.Ok(c, interface{}(result))
}

func PostHandler(c *gin.Context) {
	user := models.User{
		Name:"helloworld",
	}
	_ = db.Create(&user)

	http.Ok(c, user.ID)
}

