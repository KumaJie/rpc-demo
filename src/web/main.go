package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/web/comment"
	"rpc-douyin/src/web/favorite"
	"rpc-douyin/src/web/middleware"
	"rpc-douyin/src/web/user"
	"rpc-douyin/src/web/video"
)

func main() {
	r := gin.Default()
	r.StaticFS("/feed", http.Dir("feed"))
	group := r.Group("/douyin")
	{
		userGroup := group.Group("/user")
		{
			userGroup.POST("/login/", user.LoginHandler)
			userGroup.POST("/register/", user.RegisterHandler)
			userGroup.GET("/", middleware.AuthMiddleware, user.UserInfoHandler)
		}
		group.GET("/feed", video.FeedHandler)
		publishGroup := group.Group("/publish")
		{
			publishGroup.POST("/action/", middleware.AuthMiddleware, video.VideoPublishHandler)
			publishGroup.GET("/list/", middleware.AuthMiddleware, video.PublishListHandler)
		}
		favoriteGroup := group.Group("/favorite")
		{
			favoriteGroup.POST("/action/", middleware.AuthMiddleware, favorite.FavoriteActionHanler)
			favoriteGroup.GET("/list/", middleware.AuthMiddleware, favorite.FavoriteListHandler)
		}
		commentGroup := group.Group("/comment")
		{
			commentGroup.POST("/action/", middleware.AuthMiddleware, comment.CommentActionHandler)
			commentGroup.GET("/list/", middleware.AuthMiddleware, comment.CommentListHandler)
		}
	}
	r.Run()
}
