package main

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"os/exec"
	"path"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/fileWrapper"
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
		return &emptypb.Empty{}, err
	}

	go func() {
		coverPath := path.Join(config.Cfg.File.Dir, request.GetTitle()) + ".jpg"
		cmd := exec.Command("ffmpeg",
			"-i", filePath,
			"-vframes", "1",
			"-update", "true",
			"-y",
			"-f", "image2",
			coverPath)
		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}()

	videoInfo := model.Video{
		UserID:     request.GetUserId(),
		PlayURL:    request.GetTitle() + ".mp4",
		CoverURL:   request.GetTitle() + ".jpg",
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
		return &video.PublishListResponse{}, dbRet.Error
	}

	videoList := make([]*video.Video, 0)
	for _, videoIter := range rawVideoList {
		userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: videoIter.UserID})
		if err != nil {
			return &video.PublishListResponse{}, err
		}
		videoList = append(videoList, &video.Video{
			Id:            videoIter.ID,
			Author:        userInfoResp.GetUser(),
			PlayUrl:       fileWrapper.GetFullPath(videoIter.PlayURL),
			CoverUrl:      fileWrapper.GetFullPath(videoIter.CoverURL),
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         videoIter.Title,
		})
	}
	return &video.PublishListResponse{VideoList: videoList}, nil
}

func (v VideoServiceImpl) Feed(ctx context.Context, request *video.FeedRequest) (*video.FeedResponse, error) {
	rawVideos := make([]model.Video, 0)
	dbRet := db.DBClient.Order("create_time desc").Where("create_time < ?", time.UnixMilli(request.GetLatestTime())).Limit(3).Find(&rawVideos)
	if dbRet.Error != nil {
		return &video.FeedResponse{}, dbRet.Error
	}
	feedVideos := make([]*video.Video, 0)
	for _, rawVideo := range rawVideos {
		userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: rawVideo.UserID})
		if err != nil {
			return &video.FeedResponse{}, err
		}
		feedVideos = append(feedVideos, &video.Video{
			Id:            rawVideo.ID,
			Author:        userInfoResp.GetUser(),
			PlayUrl:       fileWrapper.GetFullPath(rawVideo.PlayURL),
			CoverUrl:      fileWrapper.GetFullPath(rawVideo.CoverURL),
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
			Title:         rawVideo.Title,
		})
	}
	// 返回的这一批视频中，发布最早的时间
	nextTime := rawVideos[len(rawVideos)-1].CreateTime.UnixMilli()
	return &video.FeedResponse{
		VideoList: feedVideos,
		NextTime:  &nextTime,
	}, nil
}
