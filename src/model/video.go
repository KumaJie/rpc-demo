package model

import (
	"rpc-douyin/src/proto/video"
	"time"
)

type Video struct {
	ID         int64     `gorm:"column:video_id"`
	UserID     int64     `gorm:"column:user_id"`
	PlayURL    string    `gorm:"column:play_url"`
	CoverURL   string    `gorm:"column:cover_url"`
	Title      string    `gorm:"column:title"`
	CreateTime time.Time `gorm:"column:create_time"`
}

func (Video) TableName() string {
	return "video"
}

type VideoResp struct {
	Response
	VideoList []*video.Video `json:"video_list"`
}

type FeedResp struct {
	Response
	VideoList []*video.Video `json:"video_list"`
	NextTime  int64          `json:"next_time"`
}
