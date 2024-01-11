package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/util/connectWrapper"
	"testing"
)

var userServiceName = config.Cfg.Server.User.Name

func TestUserServiceImpl_GetUserInfo(t *testing.T) {
	conn := connectWrapper.Connect(userServiceName)
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	ret, err := client.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: 10})
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestUserServiceImpl_UserLogin(t *testing.T) {
	conn := connectWrapper.Connect(userServiceName)
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	ret, err := client.UserLogin(context.Background(), &user.UserLoginRequest{
		Username: "xiaomin",
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Println(ret)
}

func TestUserServiceImpl_UserRegister(t *testing.T) {
	conn := connectWrapper.Connect(userServiceName)
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	ret, err := client.UserRegister(context.Background(), &user.UserRegisterRequest{
		Username: "k213",
		Password: "123456",
	})
	assert.NoError(t, err)
	fmt.Println(ret)
}
