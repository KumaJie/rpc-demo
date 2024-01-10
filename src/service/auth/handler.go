package main

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/ptypes/empty"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/util"
)

type AuthServiceImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (a AuthServiceImpl) AuthGen(ctx context.Context, request *auth.AuthGenRequest) (*auth.AuthGenResponse, error) {
	userID := request.GetUserId()
	token, err := util.GenerateToken(userID)
	return &auth.AuthGenResponse{
		Token: token,
	}, err

}

func (a AuthServiceImpl) Authenticate(ctx context.Context, request *auth.AuthRequest) (*empty.Empty, error) {
	//TODO implement me
	token := request.GetToken()
	claim, err := util.VerifyToken(token)
	fmt.Println(claim.UserId)
	return &empty.Empty{}, err
}

func (a AuthServiceImpl) mustEmbedUnimplementedAuthServiceServer() {}
