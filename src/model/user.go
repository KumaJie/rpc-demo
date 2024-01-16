package model

import "rpc-douyin/src/proto/user"

type User struct {
	ID              int64  `gorm:"primary_key;column:user_id"`
	Name            string `gorm:"column:username"`
	Password        string `gorm:"column:password"`
	Avatar          string `gorm:"column:avatar"`
	BackGroundImage string `gorm:"column:background_image"`
	Signature       string `gorm:"column:signature"`
}

func (User) TableName() string {
	return "user"
}

type UserResponse struct {
	Response
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserInfoResponse struct {
	Response
	User user.User `json:"user"`
}
