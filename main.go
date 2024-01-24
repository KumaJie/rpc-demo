package main

import (
	"fmt"
	"rpc-douyin/src/model"
	"rpc-douyin/src/storage/db"
)

func main() {
	var count int64
	ret := db.DBClient.Model(&model.User{}).Count(&count)
	if ret.Error != nil {
		fmt.Println(ret.Error)
		return
	}
	fmt.Println(count)
}
