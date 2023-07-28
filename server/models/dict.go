package models

type Dict struct {
	Id          int64  `gorm:"primaryKey"`
	Email       string `gorm:"email"`
	Password    string `gorm:"password"`
	Status      int    `gorm:"status"`
	Created     int64  `gorm:"created"`
	Updated     int64  `gorm:"updated"`
	Permissions string `gorm:"permissions"`
	Role        string `gorm:"role"`
	Address     string `gorm:"address"`
	Username    string `gorm:"username"`
	Realname    string `gorm:"realname"`
}