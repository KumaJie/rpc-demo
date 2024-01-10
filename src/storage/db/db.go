package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DBClient *gorm.DB

func init() {
	var err error
	user := "root"
	pwd := "123456"
	host := "127.0.0.1"
	port := "3306"
	database := "douyin"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pwd, host, port, database)
	DBClient, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
