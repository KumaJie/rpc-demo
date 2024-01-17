package main

import (
	"context"
	"go.uber.org/zap"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/comment"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/storage/db"
	"rpc-douyin/src/util/connectWrapper"
	"rpc-douyin/src/util/log"
	"time"
)

var userClient user.UserServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.User.Name)
	userClient = user.NewUserServiceClient(conn)
}

type CommentServiceImpl struct {
	comment.UnimplementedCommentServiceServer
}

func (c CommentServiceImpl) CommentPost(ctx context.Context, request *comment.CommentPostRequest) (*comment.CommentActionResponse, error) {
	entity := model.Comment{
		VideoID:    request.GetVideoId(),
		UserID:     request.GetUserId(),
		Content:    request.GetCommentText(),
		CreateTime: time.Now(),
	}
	dbRet := db.DBClient.Create(&entity)
	if dbRet.Error != nil {
		log.Error("CommentService: post comment failed", zap.Int64("user_id", request.GetUserId()), zap.Int64("video_id", request.GetVideoId()), zap.String("content", request.GetCommentText()), zap.Error(dbRet.Error))
		return &comment.CommentActionResponse{}, nil
	}
	ucRet, err := userClient.GetUserInfo(ctx, &user.UserInfoRequest{UserId: request.GetUserId()})
	if err != nil {
		return &comment.CommentActionResponse{}, nil
	}
	log.Info("CommentService: post comment succeed", zap.Int64("user_id", request.GetUserId()), zap.Int64("video_id", request.GetVideoId()), zap.String("content", request.GetCommentText()))
	return &comment.CommentActionResponse{
		Comment: &comment.Comment{
			Id:         entity.ID,
			User:       ucRet.GetUser(),
			Content:    entity.Content,
			CreateData: entity.CreateTime.Format("01-02"),
		},
	}, nil

}

func (c CommentServiceImpl) CommentDel(ctx context.Context, request *comment.CommentDelRequest) (*comment.CommentActionResponse, error) {
	dbRet := db.DBClient.Delete(&model.Comment{}, request.GetCommentId())
	if dbRet.Error != nil {
		log.Error("CommentService: delete comment failed", zap.Int64("comment_id", request.GetCommentId()))
		return &comment.CommentActionResponse{}, dbRet.Error
	}
	log.Info("CommentService: delete comment succeed", zap.Int64("comment_id", request.GetCommentId()))
	return &comment.CommentActionResponse{}, nil
}

func (c CommentServiceImpl) CommentList(ctx context.Context, request *comment.CommentListRequest) (*comment.CommentListResponse, error) {
	comments := make([]model.Comment, 0)
	dbRet := db.DBClient.Where("video_id = ?", request.GetVideoId()).Find(&comments)
	if dbRet.Error != nil {
		log.Error("CommentService: get comment list failed", zap.Int64("video_id", request.GetVideoId()), zap.Error(dbRet.Error))
		return &comment.CommentListResponse{}, dbRet.Error
	}
	log.Info("CommentService: get comment list succeed", zap.Int64("video_id", request.GetVideoId()), zap.Int64("num", dbRet.RowsAffected))
	out := make([]*comment.Comment, 0)
	for _, entity := range comments {
		ucRet, err := userClient.GetUserInfo(ctx, &user.UserInfoRequest{UserId: entity.UserID})
		if err != nil {
			return &comment.CommentListResponse{}, err
		}
		out = append(out, &comment.Comment{
			Id:         entity.ID,
			User:       ucRet.GetUser(),
			Content:    entity.Content,
			CreateData: entity.CreateTime.Format("01-02"),
		})
	}
	return &comment.CommentListResponse{CommentList: out}, nil
}

func (c CommentServiceImpl) CommentCount(ctx context.Context, request *comment.CommentCountRequest) (*comment.CommentCountResponse, error) {
	var count int64
	dbRet := db.DBClient.Model(&model.Comment{}).Where("video_id = ?", request.GetVideoId()).Count(&count)
	if dbRet.Error != nil {
		log.Error("CommentService: get comment count failed", zap.Int64("video_id", request.GetVideoId()), zap.Error(dbRet.Error))
		return &comment.CommentCountResponse{}, nil
	}
	log.Info("CommentService: get comment count succeed", zap.Int64("video_id", request.GetVideoId()), zap.Int64("num", count))
	return &comment.CommentCountResponse{Count: count}, nil
}
