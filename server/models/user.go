package models

type User struct {
	Id       int64  `gorm:"primaryKey"`
	Email    string `gorm:"email"`
	Password string `gorm:"password"`
	RealName string `gorm:"realname"`
	UserName string `gorm:"username"`
	Status   int    `gorm:"status"`
	Created  int64  `gorm:"created"`
	Updated  int64  `gorm:"updated"`
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
	UserName string `json:"username" form:"username" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
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
	Uid   int64  `json:"uid"`
	UserName string `json:"username"`
	Role string `json:"role"`
	Token string `json:"token"`
	Permissions string `json:"permissions"`
}

type UserPersonInfo struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Version int    `json:"version"`
}
