package apis

import (
	"gin/daos"
	http "gin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetUserByIDHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		http.ParamError(c, "id格式错误")
	}
	user := daos.FindUserByID(id)
	if user.ID == 0 {
		http.Ok(c, user)
	}
	http.Ok(c, user)
}

