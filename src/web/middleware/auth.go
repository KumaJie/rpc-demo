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
	token := c.Query("token")
	_, err := authClient.Authenticate(context.Background(), &auth.AuthRequest{Token: token})
	if err != nil {
		return
	}
	c.Next()
}
