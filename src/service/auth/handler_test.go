package main

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/util"
	"rpc-douyin/src/util/connectWrapper"
	"testing"
)

var authServiceName = config.Cfg.Server.Auth.Name

func TestAuthServiceImpl_AuthGen(t *testing.T) {
	conn := connectWrapper.Connect(authServiceName)
	defer conn.Close()
	client := auth.NewAuthServiceClient(conn)
	ret, err := client.AuthGen(context.Background(), &auth.AuthGenRequest{UserId: 10})
	assert.NoError(t, err)
	fmt.Println(ret)

}

func TestAuthServiceImpl_Authenticate(t *testing.T) {
	conn := connectWrapper.Connect(authServiceName)
	defer conn.Close()
	client := auth.NewAuthServiceClient(conn)
	token, _ := util.GenerateToken(10)
	ret, err := client.Authenticate(context.Background(), &auth.AuthRequest{Token: token})
	assert.NoError(t, err)
	fmt.Println(ret)
}
