package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/util/connectWrapper"
)

var authClient auth.AuthServiceClient

func init() {
	authConn := connectWrapper.Connect(config.Cfg.Server.Auth.Name)
	authClient = auth.NewAuthServiceClient(authConn)
}

func AuthMiddleware(c *gin.Context) {
	var token string
	if c.Request.URL.Path == "/douyin/publish/action/" {
		token = c.PostForm("token")
	} else {
		token = c.Query("token")

	}
	authResp, err := authClient.Authenticate(context.Background(), &auth.AuthRequest{Token: token})
	if err != nil {
		return
	}
	if _, ok := c.Get("user_id"); !ok {
		c.Set("user_id", authResp.UserId)
	}
	c.Next()
}
