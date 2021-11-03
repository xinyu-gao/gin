package daos

import (
	"gin/confs"
	"gin/models"
)

var (
	db = confs.GetDB()
)

func FindUserByID(id int) *models.User{
	user := &models.User{
		ID: id,
	}
	db.First(&user)
	return user
}

func SaveUser(user models.User) bool{
	result := db.Create(&user)
	return result.RowsAffected == 1
}
