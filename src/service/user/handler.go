package main

import (
	"context"
	"fmt"
	"net/http"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/storage/db"
)

type UserServiceImpl struct {
	user.UserServiceServer
}

func (s UserServiceImpl) GetUserExist(ctx context.Context, request *user.UserRequest) (*user.UserResponse, error) {
	userID := request.UserId
	var u model.User
	err := db.DBClient.Where("user_id = ?", userID).First(&u).Error
	fmt.Println(u)
	return &user.UserResponse{StatusCode: http.StatusOK}, err
}
