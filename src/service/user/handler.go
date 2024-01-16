package main

import (
	"context"
	"go.uber.org/zap"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/log"
)

type UserServiceImpl struct {
	user.UnimplementedUserServiceServer
}

func (u UserServiceImpl) GetUserInfo(ctx context.Context, request *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	userID := request.GetUserId()
	var res model.User
	err := db.DBClient.Where("user_id = ?", userID).First(&res).Error
	if err != nil {
		log.Info("UserService: can't find user detail", zap.Int64("user_id", userID))
		return &user.UserInfoResponse{}, err
	}
	return &user.UserInfoResponse{
		User: &user.User{
			Id:       res.ID,
			Name:     res.Name,
			IsFollow: false,
		},
	}, nil
}

func (u UserServiceImpl) UserLogin(ctx context.Context, request *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	var loginUser model.User
	err := db.DBClient.Where("username = ? and password = ?", request.GetUsername(), request.GetPassword()).First(&loginUser).Error
	if err != nil {
		log.Info("UserService: login failed", zap.String("username", request.GetUsername()), zap.String("password", request.GetPassword()))
		return &user.UserLoginResponse{}, err
	}
	log.Info("UserService: login success", zap.String("username", request.GetUsername()))
	return &user.UserLoginResponse{
		UserId: loginUser.ID,
	}, nil
}

func (u UserServiceImpl) UserRegister(ctx context.Context, request *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	registerUser := model.User{
		Name:     request.GetUsername(),
		Password: request.GetPassword(),
	}
	err := db.DBClient.Create(&registerUser).Error
	if err != nil {
		log.Info("UserService: register failed", zap.String("username", request.GetUsername()), zap.String("password", request.GetPassword()))
		return &user.UserRegisterResponse{}, err
	}
	return &user.UserRegisterResponse{
		UserId: registerUser.ID,
	}, nil
}

func (u UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {}
