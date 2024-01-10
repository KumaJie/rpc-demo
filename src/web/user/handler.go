package user

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/util/connectWrapper"
)

var userClient user.UserServiceClient
var authClient auth.AuthServiceClient

func init() {
	userConn := connectWrapper.Connect(8000)
	userClient = user.NewUserServiceClient(userConn)
	authConn := connectWrapper.Connect(8001)
	authClient = auth.NewAuthServiceClient(authConn)
}

func LoginHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ret, err := userClient.UserLogin(context.Background(), &user.UserLoginRequest{
		Username: username,
		Password: password,
	})
	if err != nil {

	}
	userID := ret.GetUserId()
	res, err := authClient.AuthGen(context.Background(), &auth.AuthGenRequest{UserId: ret.GetUserId()})
	if err != nil {

	}
	token := res.GetToken()
	c.JSON(http.StatusOK, map[string]interface{}{
		"status_code": 0,
		"status_msg":  "登陆成功",
		"user_id":     userID,
		"token":       token,
	})

}
