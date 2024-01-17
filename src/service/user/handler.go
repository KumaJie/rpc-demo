package main

import (
	"context"
	"go.uber.org/zap"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/log"
	"sync"
)

var favoriteClient favorite.FavoriteServiceClient
var videoClient video.VideoServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.Favorite.Name)
	favoriteClient = favorite.NewFavoriteServiceClient(conn)
	conn = connectWrapper.Connect(config.Cfg.Server.Video.Name)
	videoClient = video.NewVideoServiceClient(conn)
}

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
	g := sync.WaitGroup{}
	g.Add(3)
	var totalFavorited int64
	var workCount int64
	var favoriteCount int64
	go func() {
		defer g.Done()
		fcRet, _ := favoriteClient.UserFavoriteCount(ctx, &favorite.UserFavoriteCountRequest{UserId: userID})
		favoriteCount = fcRet.GetFavoriteCount()
	}()

	go func() {
		defer g.Done()
		fcRet, _ := favoriteClient.UserTotalFavorite(ctx, &favorite.UserFavoriteCountRequest{UserId: userID})
		totalFavorited = fcRet.GetFavoriteCount()
	}()

	go func() {
		defer g.Done()
		vcRet, _ := videoClient.PublishCount(ctx, &video.PublishCountRequest{UserId: userID})
		workCount = vcRet.Count
	}()

	g.Wait()
	return &user.UserInfoResponse{
		User: &user.User{
			Id:              res.ID,
			Name:            res.Name,
			FollowCount:     nil,
			FollowerCount:   nil,
			IsFollow:        false,
			Avatar:          nil,
			BackgroundImage: nil,
			Signature:       nil,
			TotalFavorited:  &totalFavorited,
			WorkCount:       &workCount,
			FavoriteCount:   &favoriteCount,
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
