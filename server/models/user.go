package models

type User struct {
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

type UserCreateParam struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	Password string `json:"password" binding:"required"`
}

type UserDeleteParam struct {
	Id    int64  `json:"id,omitempty" binding:"-"`
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=6"`
}

type UserLoginParam struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserVerifyCodeParam struct {
	Email string `form:"email" binding:"required,email"`
}

type UserPassParam struct {
	Email    string `json:"email" binding:"required,email"`
	Code     string `json:"code" binding:"required,len=6"`
	Password string `json:"password" binding:"required"`
}

type UserInfo struct {
	Uid         int64  `json:"uid"`
	Role        string `json:"role"`
	Username    string `json:"username"`
	Permissions string `json:"permissions"`
	Token       string `json:"token"`
}

type UserPersonInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Version int    `json:"version"`
}
