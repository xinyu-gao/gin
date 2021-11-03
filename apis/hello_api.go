package apis

import (
	"gin/confs"
	"gin/daos"
	"gin/models"
	"github.com/gin-gonic/gin"
)

var (
	db         = confs.GetDB()
	redis, ctx = confs.GetRedis()
)

func HelloHandler(c *gin.Context) {
	//result := redis.Get(ctx, "id").Val()
	//users := models.SqlUser{
	//	Username: "xinyu",
	//}
	//_ = db.First(&users)
	//http.Ok(c, users)
	users := models.User{
			ID: 1,
			Name: "xinyu",
			Password: "Sdasds",
			Status: "1",
		}
		daos.SaveUser(users)
}

func PostHandler(c *gin.Context) {
	//user := models.SqlUser{
	//	Username: "xinyu",
	//}
	//_ = db.Create(&user)
	//
	//http.Ok(c, user.Uid)
}
