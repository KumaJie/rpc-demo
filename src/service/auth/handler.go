package main

import (
	"context"
	"fmt"
	"net/http"
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
		StatusCode: http.StatusOK,
		StatusMsg:  "",
		Token:      token,
	}, err

}

func (a AuthServiceImpl) Authenticate(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {
	//TODO implement me
	token := request.GetToken()
	claim, err := util.VerifyToken(token)
	fmt.Println(claim.UserId)
	return &auth.AuthResponse{
		StatusCode: http.StatusOK,
		StatusMsg:  "",
	}, err
}

func (a AuthServiceImpl) mustEmbedUnimplementedAuthServiceServer() {}
