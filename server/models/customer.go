package models

import (
	"time"
)

type Customer struct {
	Id          int64     `gorm:"primaryKey"`
	Name        string    `gorm:"name"`
	Source      string    `gorm:"source"`
	Mobilephone string    `gorm:"mobilephone"`
	Email       string    `gorm:"email"`
	Industry    string    `gorm:"industry"`
	Level       string    `gorm:"level"`
	Remarks     string    `gorm:"remarks"`
	Region      string    `gorm:"region"`
	Address     string    `gorm:"address"`
	Status      int       `gorm:"status"`
	Creator     int64     `gorm:"creator"`
	Created     time.Time `gorm:"created"`
	Updated     time.Time `gorm:"updated"`
	Domain      string    `gorm:"domain"`
	Contact     string    `gorm:"contact"`
	Latestsign  time.Time `gorm:"latestsign"`
	Importance  int       `gorm:"importance"`
}

type CustomerCreateParam struct {
	Name       string `json:"name" binding:"required"`
	Source     string `json:"source" binding:"-"`
	Phone      string `json:"phone" binding:"omitempty,len=11"`
	Email      string `json:"email" binding:"omitempty,email"`
	Industry   string `json:"industry" binding:"-"`
	Level      string `json:"level" binding:"-"`
	Remarks    string `json:"remarks" binding:"-"`
	Region     string `json:"region" binding:"-"`
	Address    string `json:"address" binding:"-"`
	Status     int    `json:"status" binding:"-"`
	Creator    int64  `json:"creator,omitempty" binding:"-"`
	Importance int    `json:"importance"`
}

type CustomerUpdateParam struct {
	Id          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"-"`
	Source      string `json:"source" binding:"-"`
	Mobilephone string `json:"mobilephone" binding:"omitempty,len=11"`
	Email       string `json:"email" binding:"omitempty,email"`
	Industry    string `json:"industry" binding:"-"`
	Level       string `json:"level" binding:"-"`
	Remarks     string `json:"remarks" binding:"-"`
	Region      string `json:"region" binding:"-"`
	Address     string `json:"address" binding:"-"`
	Status      int    `json:"status" binding:"-"`
	Contact     string `json:"contact" binding:"-"`
	Domain      string `json:"domain" binding:"-"`
	Latestsign  string `json:"latestsign" binding:"-"`
	Importance  int    `json:"importance" binding:"-"`
}

type CustomerSendMailParam struct {
	Uid        int64  `json:"uid" binding:"-"`
	Receiver   string `json:"receiver" binding:"required,email"`
	Subject    string `json:"subject" binding:"omitempty,gt=0"`
	Content    string `json:"content" binding:"required,gt=0"`
	Attachment string `json:"attachment" binding:"omitempty,gt=0"`
}

type CustomerDeleteParam struct {
	Ids []int64 `json:"ids" binding:"required"`
}

type CustomerQueryParam struct {
	Id       int64  `form:"id" binding:"omitempty,gt=0"`
	Name     string `form:"name" binding:"omitempty,gt=0"`
	Source   string `form:"source" binding:"omitempty,gt=0"`
	Industry string `form:"industry" binding:"omitempty,gt=0"`
	Level    string `form:"level" binding:"omitempty,gt=0"`
	Status   int    `form:"status" binding:"omitempty,oneof=1 2"`
	Creator  int64  `form:"creator,omitempty" binding:"-"`
	Page     Page
}

type CustomerList struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Source      string    `json:"source"`
	Mobilephone string    `json:"mobilephone"`
	Email       string    `json:"email"`
	Industry    string    `json:"industry"`
	Level       string    `json:"level"`
	Remarks     string    `json:"remarks"`
	Region      string    `json:"region"`
	Address     string    `json:"address"`
	Status      int       `json:"status"`
	Created     time.Time `json:"created"`
	Updated     time.Time `json:"updated"`
	Domain      string    `json:"domain"`
	Contact     string    `json:"contact"`
	Latestsign  time.Time `json:"latestsign"`
	Importance  int       `json:"importance"`
}

type CustomerInfo struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Source      string `json:"source"`
	Mobilephone string `json:"mobilephone"`
	Email       string `json:"email"`
	Industry    string `json:"industry"`
	Level       string `json:"level"`
	Remarks     string `json:"remarks"`
	Region      string `json:"region"`
	Address     string `json:"address"`
	Status      int    `json:"status"`
}

type CustomerOption struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type CustomerExcelRow struct {
	Name        string `json:"name"`
	Source      string `json:"source"`
	Mobilephone string `json:"mobilephone"`
	Email       string `json:"email"`
	Industry    string `json:"industry"`
	Level       string `json:"level"`
	Remarks     string `json:"remarks"`
	Region      string `json:"region"`
	Address     string `json:"address"`
	Status      string `json:"status"`
	Creator     int64  `json:"creator"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
}
