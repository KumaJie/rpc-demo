package main

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/log"
)

var videoClient video.VideoServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.Video.Name)
	videoClient = video.NewVideoServiceClient(conn)
}

type FavoriteServiceImpl struct {
	favorite.UnimplementedFavoriteServiceServer
}

func (f FavoriteServiceImpl) FavoriteAction(ctx context.Context, request *favorite.FavoriteActionRequest) (*emptypb.Empty, error) {
	switch request.GetActionType() {
	case 1:
		entity := model.Favorite{
			VideoID: request.GetVideoId(),
			UserID:  request.GetUserId(),
		}
		dbRet := db.DBClient.Create(&entity)
		if dbRet.Error != nil {
			log.Error("FavoriteService: create favorite entity failed", zap.Int64("user_id", request.GetUserId()), zap.Int64("video_id", request.GetVideoId()), zap.Error(dbRet.Error))
			return &emptypb.Empty{}, dbRet.Error
		}
	case 2:
		dbRet := db.DBClient.Delete(&model.Favorite{}, "video_id = ? and user_id = ?", request.GetVideoId(), request.GetUserId())
		if dbRet.Error != nil {
			log.Error("FavoriteService: delete favorite entity failed,", zap.Int64("user_id", request.GetUserId()), zap.Int64("video_id", request.GetVideoId()), zap.Error(dbRet.Error))
			return &emptypb.Empty{}, dbRet.Error
		}
	}
	return &emptypb.Empty{}, nil
}

func (f FavoriteServiceImpl) FavoriteList(ctx context.Context, request *favorite.FavoriteListRequest) (*favorite.FavoriteListResponse, error) {
	favorites := make([]model.Favorite, 0)
	dbRet := db.DBClient.Where("user_id = ?", request.GetUserId()).Find(&favorites)
	if dbRet.Error != nil {
		log.Error("FavoriteService: get favorite list failed", zap.Int64("user_id", request.GetUserId()), zap.Error(dbRet.Error))
		return &favorite.FavoriteListResponse{}, dbRet.Error
	}
	log.Info("FavoriteService: get favorite list succeed", zap.Int64("user_id", request.GetUserId()), zap.Int64("num", dbRet.RowsAffected))
	favoriteVideos := make([]*video.Video, 0)
	for _, entity := range favorites {
		vcRet, err := videoClient.GetVideo(ctx, &video.GetVideoRequest{VideoId: entity.VideoID, UserId: request.GetUserId()})
		if err != nil {
			return &favorite.FavoriteListResponse{}, err
		}
		favoriteVideos = append(favoriteVideos, vcRet.GetVideo())
	}
	return &favorite.FavoriteListResponse{VideoList: favoriteVideos}, nil
}

func (f FavoriteServiceImpl) UserFavoriteCount(ctx context.Context, request *favorite.UserFavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	var count int64
	dbRet := db.DBClient.Model(&model.Favorite{}).Where("user_id = ?", request.GetUserId()).Count(&count)
	if dbRet.Error != nil {
		log.Error("FavoriteService: get user favorite count failed", zap.Int64("user_id", request.GetUserId()), zap.Error(dbRet.Error))
		return &favorite.FavoriteCountResponse{}, dbRet.Error
	}
	return &favorite.FavoriteCountResponse{FavoriteCount: count}, nil
}

func (f FavoriteServiceImpl) UserTotalFavorite(ctx context.Context, request *favorite.UserFavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	vcRet, err := videoClient.GetPublishId(ctx, &video.PublishListRequest{UserId: request.GetUserId()})
	if err != nil {
		return &favorite.FavoriteCountResponse{}, err
	}
	publishIDs := vcRet.GetVideoId()
	var count int64 = 0
	for _, id := range publishIDs {
		t, err := GetFavoriteInVideo(id)
		if err != nil {
			return &favorite.FavoriteCountResponse{}, err
		}
		log.Info("FavoriteService: get video favorite count succeed", zap.Int64("video_id", id), zap.Int64("count", t))
		count += t
	}
	return &favorite.FavoriteCountResponse{FavoriteCount: count}, nil
}

func (f FavoriteServiceImpl) VideoFavoriteCount(ctx context.Context, request *favorite.VideoFavoriteCountRequest) (*favorite.FavoriteCountResponse, error) {
	count, err := GetFavoriteInVideo(request.GetVideoId())
	if err != nil {
		return &favorite.FavoriteCountResponse{}, err
	}
	return &favorite.FavoriteCountResponse{FavoriteCount: count}, err
}

func (f FavoriteServiceImpl) IsFavorite(ctx context.Context, request *favorite.IsFavoriteRequest) (*favorite.IsFavoriteResponse, error) {
	var count int64
	dbRet := db.DBClient.Model(&model.Favorite{}).Where("user_id = ? and video_id = ?", request.GetUserId(), request.GetVideoId()).Count(&count)
	if dbRet.Error != nil {
		log.Error("FavoriteService: get favorite failed", zap.Int64("user_id", request.GetUserId()), zap.Int64("video_id", request.GetVideoId()))
		return &favorite.IsFavoriteResponse{}, dbRet.Error
	}
	return &favorite.IsFavoriteResponse{Favorite: count > 0}, nil
}

func GetFavoriteInVideo(videoID int64) (count int64, e error) {
	dbRet := db.DBClient.Model(&model.Favorite{}).Where("video_id = ?", videoID).Count(&count)
	if dbRet.Error != nil {
		e = dbRet.Error
		log.Error("FavoriteService: get video favorite count failed", zap.Int64("video_id", videoID), zap.Error(dbRet.Error))
		return
	}
	return
}
