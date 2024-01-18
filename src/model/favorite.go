package model

import "rpc-douyin/src/proto/video"

type Favorite struct {
	ID      int64 `gorm:"column:favorite_id"`
	VideoID int64 `gorm:"column:video_id"`
	UserID  int64 `gorm:"column:user_id"`
}

func (Favorite) TableName() string {
	return "favorite"
}

type FavoriteListResp struct {
	Response
	VideoList []*video.Video `json:"video_list"`
}
