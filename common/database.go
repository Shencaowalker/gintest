package common

import (
	"fmt"
	"gintest1/model"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

//初始化数据库链接
func InitDB() *gorm.DB {
	driverName := "mysql"
	host := "116.62.114.212"
	post := "3306"
	database := "gintest1"
	username := "root"
	password := "123456"
	charset := "utf8mb4"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		post,
		database,
		charset)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err: " + err.Error())
	}
	//创建数据库迁移  这里主要是user表 
	db.AutoMigrate(&model.User{})
	DB = db
	return db
}
//得到数据库链接
func GetDB() *gorm.DB {
	return DB
}
