package user

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/tracer"
	"rpc-douyin/src/web/middleware"
	"strconv"
)

var userClient user.UserServiceClient
var authClient auth.AuthServiceClient

func init() {
	userConn := connectWrapper.Connect(config.Cfg.Server.User.Name)
	userClient = user.NewUserServiceClient(userConn)
	authConn := connectWrapper.Connect(config.Cfg.Server.Auth.Name)
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
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "登录成功",
		},
		UserID: userID,
		Token:  token,
	})
}

func RegisterHandler(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	ret, err := userClient.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		return
	}
	userID := ret.GetUserId()
	res, err := authClient.AuthGen(context.Background(), &auth.AuthGenRequest{UserId: userID})
	if err != nil {
		return
	}
	token := res.GetToken()
	c.JSON(http.StatusOK, model.UserResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "注册成功",
		},
		UserID: userID,
		Token:  token,
	})
}

func UserInfoHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")
	spanCtx, _ := c.Get(middleware.SpanCtx)
	span, _ := tracer.NewSpanFromContext(spanCtx.(context.Context), "UserInfoHandler")
	defer span.Finish()

	ctx := tracer.InjectGRPC(span.Context())

	userID, _ := strconv.ParseInt(userIDStr, 10, 64)
	ret, err := userClient.GetUserInfo(ctx, &user.UserInfoRequest{UserId: userID})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, model.UserInfoResponse{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		User: *ret.GetUser(),
	})
}
