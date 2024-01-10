package main

import (
	"context"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/storage/db"
)

type UserServiceImpl struct {
	user.UnimplementedUserServiceServer
}

func (u UserServiceImpl) GetUserInfo(ctx context.Context, request *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	//TODO implement me
	userID := request.GetUserId()
	var res model.User
	err := db.DBClient.Where("user_id = ?", userID).First(&res).Error
	return &user.UserInfoResponse{
		User: &user.User{
			Id:       res.ID,
			Name:     res.Name,
			IsFollow: false,
		},
	}, err
}

func (u UserServiceImpl) UserLogin(ctx context.Context, request *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	//TODO implement me
	var loginUser model.User
	err := db.DBClient.Where("username = ? and password = ?", request.GetUsername(), request.GetPassword()).First(&loginUser).Error
	return &user.UserLoginResponse{
		UserId: loginUser.ID,
	}, err
}

func (u UserServiceImpl) UserRegister(ctx context.Context, request *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	//TODO implement me
	registerUser := model.User{
		Name:     request.GetUsername(),
		Password: request.GetPassword(),
	}
	err := db.DBClient.Create(&registerUser).Error
	return &user.UserRegisterResponse{
		UserId: registerUser.ID,
	}, err
}

func (u UserServiceImpl) mustEmbedUnimplementedUserServiceServer() {}
