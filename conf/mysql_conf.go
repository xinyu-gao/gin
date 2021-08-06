package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	sqlConf := GetMysqlConf()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		sqlConf.username,
		sqlConf.password,
		sqlConf.host,
		sqlConf.port,
		sqlConf.dbname,
		sqlConf.timeout)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100) //最大连接数
	sqlDB.SetMaxIdleConns(20)  //最大允许的空闲连接数
}

func GetDB() *gorm.DB {
	return db
}