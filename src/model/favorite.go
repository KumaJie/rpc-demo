package model

type Favorite struct {
	ID      int64 `gorm:"column:favorite_id"`
	VideoID int64 `gorm:"column:video_id"`
	UserID  int64 `gorm:"column:user_id"`
}

func (Favorite) TableName() string {
	return "favorite"
}
