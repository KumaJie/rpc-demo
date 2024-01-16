package main

import (
	"context"
	"go.uber.org/zap"
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
	"rpc-douyin/src/util/log"
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
		log.Error("VideoService: save file failed", zap.String("file", filePath))
		return &emptypb.Empty{}, err
	}
	log.Info("VideoService: save file succeed", zap.String("file", filePath))
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
			log.Error("VideoService: generate cover failed", zap.String("cover", coverPath))
			return
		}
		log.Info("VideoService: generate cover succeed", zap.String("cover", coverPath))
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
		log.Error("VideoService: get VideoList with user_id failed in db", zap.Int64("user_id", userID))
		return &video.PublishListResponse{}, dbRet.Error
	}
	log.Info("VideoService: get VideoList succeed", zap.Int64("user_id", userID), zap.Int64("num", dbRet.RowsAffected))
	videoList := make([]*video.Video, 0)
	userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: userID})
	if err != nil {
		log.Error("VideoService: get userinfo failed", zap.Int64("user_id", userID))
		return &video.PublishListResponse{}, err
	}
	for _, videoIter := range rawVideoList {
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
		log.Error("VideoService: get feed failed in db ")
		return &video.FeedResponse{}, dbRet.Error
	}
	log.Info("VideoService: get feed succeed", zap.Int64("latest_time", request.GetLatestTime()), zap.Int64("num", dbRet.RowsAffected))
	feedVideos := make([]*video.Video, 0)
	for _, rawVideo := range rawVideos {
		userInfoResp, err := userClient.GetUserInfo(context.Background(), &user.UserInfoRequest{UserId: rawVideo.UserID})
		if err != nil {
			log.Error("VideoService: get userinfo failed", zap.Int64("user_id", rawVideo.UserID))
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
