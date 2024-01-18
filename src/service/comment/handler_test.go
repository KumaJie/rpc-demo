package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/comment"
	"rpc-douyin/src/util/connectWrapper"
	"testing"
	"time"
)

func TestCommentServiceImpl_CommentPost(t *testing.T) {
	conn := connectWrapper.Connect(config.Cfg.Server.Comment.Name)
	client := comment.NewCommentServiceClient(conn)
	t.Run("新增评论", func(t *testing.T) {
		userID := int64(1)
		videoID := int64(1)
		content := "新的评论"
		ret, err := client.CommentPost(context.Background(), &comment.CommentPostRequest{
			UserId:      userID,
			VideoId:     videoID,
			CommentText: content,
		})
		assert.NoError(t, err)
		assert.Equal(t, userID, ret.GetComment().GetUser().GetId())
		assert.Equal(t, time.Now().Format("01-02"), ret.GetComment().CreateData)
		assert.Equal(t, content, ret.GetComment().Content)
	})
}

func TestCommentServiceImpl_CommentDel(t *testing.T) {
	conn := connectWrapper.Connect(config.Cfg.Server.Comment.Name)
	client := comment.NewCommentServiceClient(conn)
	t.Run("删除评论", func(t *testing.T) {
		ret, err := client.CommentDel(context.Background(), &comment.CommentDelRequest{CommentId: 1})
		assert.NoError(t, err)
		assert.Nil(t, ret.Comment)
	})
}
