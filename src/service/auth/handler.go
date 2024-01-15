package main

import (
	"context"
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

func (a AuthServiceImpl) Authenticate(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {
	token := request.GetToken()
	claim, err := util.VerifyToken(token)
	return &auth.AuthResponse{UserId: claim.UserId}, err
}

func (a AuthServiceImpl) mustEmbedUnimplementedAuthServiceServer() {}
