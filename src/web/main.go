package main

import (
	"github.com/gin-gonic/gin"
	"rpc-douyin/src/web/middleware"
	"rpc-douyin/src/web/user"
)

func main() {
	r := gin.Default()
	group := r.Group("/douyin")
	{
		userGroup := group.Group("/user")
		{
			userGroup.POST("/login/", user.LoginHandler)
			userGroup.POST("/register/", user.RegisterHandler)
			userGroup.GET("/", middleware.AuthMiddleware, user.UserInfoHandler)
		}
	}
	r.Run()
}
