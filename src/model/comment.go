package model

import (
	"rpc-douyin/src/proto/comment"
	"time"
)

type Comment struct {
	ID         int64     `gorm:"column:comment_id"`
	VideoID    int64     `gorm:"column:video_id"`
	UserID     int64     `gorm:"column:user_id"`
	Content    string    `gorm:"column:content"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Comment) TableName() string {
	return "comment"
}

type CommentActionResp struct {
	Response
	Comment *comment.Comment `json:"comment"`
}

type CommentListResp struct {
	Response
	CommentList []*comment.Comment `json:"comment_list"`
}
