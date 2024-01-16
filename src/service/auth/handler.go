package main

import (
	"context"
	"go.uber.org/zap"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/util"
	"rpc-douyin/src/util/log"
)

type AuthServiceImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (a AuthServiceImpl) AuthGen(ctx context.Context, request *auth.AuthGenRequest) (*auth.AuthGenResponse, error) {
	userID := request.GetUserId()
	token, err := util.GenerateToken(userID)
	log.Info("AuthService: generate token", zap.Int64("user_id", userID), zap.String("token", token))
	return &auth.AuthGenResponse{
		Token: token,
	}, err

}

func (a AuthServiceImpl) Authenticate(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {
	token := request.GetToken()
	claim, err := util.VerifyToken(token)
	if err != nil {
		log.Error("AuthService: parse token failed", zap.String("token", token), zap.Error(err))
		return &auth.AuthResponse{}, err
	}
	return &auth.AuthResponse{UserId: claim.UserId}, nil
}

func (a AuthServiceImpl) mustEmbedUnimplementedAuthServiceServer() {}
