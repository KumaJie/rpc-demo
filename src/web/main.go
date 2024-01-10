package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/web/user"
)

func main() {
	r := gin.Default()
	r.POST("/douyin/user/login/", user.LoginHandler)
	r.GET("/douyin/user/", func(context *gin.Context) {
		context.JSON(http.StatusOK, map[string]interface{}{
			"status_code": 0,
			"status_msg":  "",
			"user": map[string]interface{}{
				"id":        1,
				"name":      "测试",
				"is_follow": false,
			},
		})
	})
	r.Run()
}
