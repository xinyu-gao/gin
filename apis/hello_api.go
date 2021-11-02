package apis

import (
	"gin/confs"
	"gin/models"
	http "gin/utils"
	"github.com/gin-gonic/gin"
)

var (
	db         = confs.GetDB()
	redis, ctx = confs.GetRedis()
)

func HelloHandler(c *gin.Context) {
	//result := redis.Get(ctx, "id").Val()
	users := models.SqlUser{
		Username: "xinyu",
	}
	_ = db.First(&users)
	http.Ok(c, users)
}

func PostHandler(c *gin.Context) {
	user := models.SqlUser{
		Username: "xinyu",
	}
	_ = db.Create(&user)

	http.Ok(c, user.Uid)
}
