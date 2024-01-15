package main

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"path"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"time"
)

var userClient user.UserServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.User.Name)
	userClient = user.NewUserServiceClient(conn)
}

type VideoServiceImpl struct {
	video.UnimplementedVideoServiceServer
}

func (v VideoServiceImpl) VideoPublish(ctx context.Context, request *video.PublishRequest) (*emptypb.Empty, error) {
	filePath := path.Join(config.Cfg.File.Dir, request.GetTitle()) + ".mp4"
	err := os.WriteFile(filePath, request.GetData(), 755)
	if err != nil {
		return nil, err
	}
	videoInfo := model.Video{
		UserID:     request.GetUserId(),
		PlayURL:    filePath,
		CoverURL:   "",
		Title:      request.GetTitle(),
		CreateTime: time.Now(),
	}
	err = db.DBClient.Create(&videoInfo).Error
	return &emptypb.Empty{}, err
}

func (v VideoServiceImpl) GetPublistList(ctx context.Context, request *video.PublishListRequest) (*video.PublishListResponse, error) {
	userID := request.GetUserId()
	rawVideoList := make([]model.Video, 0)
	dbRet := db.DBClient.Where("user_id = ?", userID).Find(&rawVideoList)
	if dbRet.Error != nil {
		return nil, dbRet.Error
	}

	videoList := make([]*video.Video, 0)
	for _, videoIter := range rawVideoList {
		userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: videoIter.UserID})
		if err != nil {
			return nil, err
		}
		videoList = append(videoList, &video.Video{
			Id:            videoIter.ID,
			Author:        userInfoResp.GetUser(),
			PlayUrl:       videoIter.PlayURL,
			CoverUrl:      videoIter.CoverURL,
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         videoIter.Title,
		})
	}
	return &video.PublishListResponse{VideoList: videoList}, nil
}

func (v VideoServiceImpl) Feed(ctx context.Context, request *video.FeedRequest) (*video.FeedResponse, error) {
	rawVideo := model.Video{}
	dbRet := db.DBClient.Order("create_time desc").First(&rawVideo)
	if dbRet.Error != nil {
		return nil, dbRet.Error
	}
	userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: rawVideo.UserID})
	if err != nil {
		return nil, err
	}
	feedVideo := video.Video{
		Id:            rawVideo.ID,
		Author:        userInfoResp.GetUser(),
		PlayUrl:       rawVideo.PlayURL,
		CoverUrl:      rawVideo.CoverURL,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         rawVideo.Title,
	}
	t := rawVideo.CreateTime.Unix()
	return &video.FeedResponse{
		VideoList: &feedVideo,
		NextTime:  &t,
	}, nil
}
