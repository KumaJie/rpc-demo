package model

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
