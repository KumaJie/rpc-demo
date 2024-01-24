package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"os"
	"path"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/comment"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/fileWrapper"
	"rpc-douyin/src/util/log"
	"time"
)

var userClient user.UserServiceClient
var favoriteClient favorite.FavoriteServiceClient
var commentClient comment.CommentServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.User.Name)
	userClient = user.NewUserServiceClient(conn)
	conn = connectWrapper.Connect(config.Cfg.Server.Favorite.Name)
	favoriteClient = favorite.NewFavoriteServiceClient(conn)
	conn = connectWrapper.Connect(config.Cfg.Server.Comment.Name)
	commentClient = comment.NewCommentServiceClient(conn)
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
	videoInfo := model.Video{
		UserID:     request.GetUserId(),
		PlayURL:    request.GetTitle() + ".mp4",
		CoverURL:   request.GetTitle() + ".jpg",
		Title:      request.GetTitle(),
		CreateTime: time.Now(),
	}
	err = db.DBClient.Create(&videoInfo).Error
	//producer, _ := mq.NewSyncProducer()
	//msg, _ := mq.NewMessage("publish", request.GetTitle())
	//producer.SendMessage(msg)
	return &emptypb.Empty{}, err
}

func (v VideoServiceImpl) GetPublishList(ctx context.Context, request *video.PublishListRequest) (*video.PublishListResponse, error) {
	userID := request.GetUserId()
	rawVideoList := make([]model.Video, 0)
	dbRet := db.DBClient.Where("user_id = ?", userID).Find(&rawVideoList)
	if dbRet.Error != nil {
		log.Error("VideoService: get VideoList with user_id failed in db", zap.Int64("user_id", userID))
		return &video.PublishListResponse{}, dbRet.Error
	}
	log.Info("VideoService: get VideoList succeed", zap.Int64("user_id", userID), zap.Int64("num", dbRet.RowsAffected))
	videoList := make([]*video.Video, 0)
	userInfoResp, err := userClient.GetUserInfo(ctx, &user.UserInfoRequest{UserId: userID})
	if err != nil {
		log.Error("VideoService: get userinfo failed", zap.Int64("user_id", userID))
		return &video.PublishListResponse{}, err
	}
	author := userInfoResp.GetUser()
	for _, videoIter := range rawVideoList {
		fcRet, err := favoriteClient.VideoFavoriteCount(ctx, &favorite.VideoFavoriteCountRequest{VideoId: videoIter.ID})
		if err != nil {
			return &video.PublishListResponse{}, err
		}
		ifRet, err := favoriteClient.IsFavorite(ctx, &favorite.IsFavoriteRequest{
			UserId:  author.Id,
			VideoId: videoIter.ID,
		})
		if err != nil {
			return &video.PublishListResponse{}, err
		}

		ccRet, err := commentClient.CommentCount(ctx, &comment.CommentCountRequest{VideoId: videoIter.ID})
		if err != nil {
			return &video.PublishListResponse{}, err
		}

		videoList = append(videoList, &video.Video{
			Id:            videoIter.ID,
			Author:        author,
			PlayUrl:       fileWrapper.GetFullPath(videoIter.PlayURL),
			CoverUrl:      fileWrapper.GetFullPath(videoIter.CoverURL),
			FavoriteCount: fcRet.GetFavoriteCount(),
			CommentCount:  ccRet.GetCount(),
			IsFavorite:    ifRet.GetFavorite(),
			Title:         videoIter.Title,
		})
	}
	return &video.PublishListResponse{VideoList: videoList}, nil
}

func (v VideoServiceImpl) GetPublishId(ctx context.Context, request *video.PublishListRequest) (*video.PublishIdListResponse, error) {
	userID := request.GetUserId()
	rawVideoList := make([]model.Video, 0)
	dbRet := db.DBClient.Where("user_id = ?", userID).Find(&rawVideoList)
	if dbRet.Error != nil {
		log.Error("VideoService: get VideoList with user_id failed in db", zap.Int64("user_id", userID))
		return &video.PublishIdListResponse{}, dbRet.Error
	}
	log.Info("VideoService: get VideoList succeed", zap.Int64("user_id", userID), zap.Int64("num", dbRet.RowsAffected))
	IDs := make([]int64, 0)
	for _, rawVideo := range rawVideoList {
		IDs = append(IDs, rawVideo.ID)
	}
	return &video.PublishIdListResponse{VideoId: IDs}, nil
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

		fcRet, err := favoriteClient.VideoFavoriteCount(ctx, &favorite.VideoFavoriteCountRequest{VideoId: rawVideo.ID})
		if err != nil {
			return &video.FeedResponse{}, err
		}

		ifLike := false
		// 对于未登录的用户不需要判断是否喜爱
		if request.UserId != nil {
			ifRet, err := favoriteClient.IsFavorite(ctx, &favorite.IsFavoriteRequest{
				UserId:  rawVideo.UserID,
				VideoId: rawVideo.ID,
			})
			if err != nil {
				return &video.FeedResponse{}, err
			}
			ifLike = ifRet.GetFavorite()
		}

		ccRet, err := commentClient.CommentCount(ctx, &comment.CommentCountRequest{VideoId: rawVideo.ID})
		if err != nil {
			return &video.FeedResponse{}, err
		}

		feedVideos = append(feedVideos, &video.Video{
			Id:            rawVideo.ID,
			Author:        userInfoResp.GetUser(),
			PlayUrl:       fileWrapper.GetFullPath(rawVideo.PlayURL),
			CoverUrl:      fileWrapper.GetFullPath(rawVideo.CoverURL),
			FavoriteCount: fcRet.GetFavoriteCount(),
			CommentCount:  ccRet.GetCount(),
			IsFavorite:    ifLike,
			Title:         rawVideo.Title,
		})
	}
	if len(rawVideos) == 0 {
		return &video.FeedResponse{}, nil
	}
	// 返回的这一批视频中，发布最早的时间
	var nextTime int64
	nextTime = rawVideos[len(rawVideos)-1].CreateTime.UnixMilli()

	return &video.FeedResponse{
		VideoList: feedVideos,
		NextTime:  &nextTime,
	}, nil
}

func (v VideoServiceImpl) PublishCount(ctx context.Context, request *video.PublishCountRequest) (*video.PublishCountResponse, error) {
	var count int64
	dbRet := db.DBClient.Model(&model.Video{}).Where("user_id = ?", request.GetUserId()).Count(&count)
	if dbRet.Error != nil {
		log.Error("VideoService: get publish count failed", zap.Int64("user_id", request.GetUserId()), zap.Error(dbRet.Error))
		return &video.PublishCountResponse{}, dbRet.Error
	}
	log.Info("VideoService: get publish count succeed", zap.Int64("user_id", request.GetUserId()), zap.Int64("num", count))
	return &video.PublishCountResponse{Count: count}, nil
}

func (v VideoServiceImpl) GetVideo(ctx context.Context, request *video.GetVideoRequest) (*video.GetVideoResponse, error) {
	rawVideo := model.Video{ID: request.GetVideoId()}
	dbRet := db.DBClient.Find(&rawVideo)
	if dbRet.Error != nil {
		log.Error("VideoService: get video by video_id failed", zap.Int64("video_id", request.GetVideoId()), zap.Error(dbRet.Error))
	}

	ucRet, err := userClient.GetUserInfo(ctx, &user.UserInfoRequest{UserId: rawVideo.UserID})
	if err != nil {
		return &video.GetVideoResponse{}, err
	}

	fcRet, err := favoriteClient.VideoFavoriteCount(ctx, &favorite.VideoFavoriteCountRequest{VideoId: rawVideo.ID})
	if err != nil {
		return &video.GetVideoResponse{}, err
	}

	ifRet, err := favoriteClient.IsFavorite(ctx, &favorite.IsFavoriteRequest{
		UserId:  request.GetUserId(),
		VideoId: rawVideo.ID,
	})
	if err != nil {
		return &video.GetVideoResponse{}, err
	}

	ccRet, err := commentClient.CommentCount(ctx, &comment.CommentCountRequest{VideoId: rawVideo.ID})
	if err != nil {
		return &video.GetVideoResponse{}, err
	}

	out := video.Video{
		Id:            rawVideo.ID,
		Author:        ucRet.GetUser(),
		PlayUrl:       fileWrapper.GetFullPath(rawVideo.PlayURL),
		CoverUrl:      fileWrapper.GetFullPath(rawVideo.CoverURL),
		FavoriteCount: fcRet.GetFavoriteCount(),
		CommentCount:  ccRet.GetCount(),
		IsFavorite:    ifRet.GetFavorite(),
		Title:         rawVideo.Title,
	}

	return &video.GetVideoResponse{Video: &out}, nil
}
