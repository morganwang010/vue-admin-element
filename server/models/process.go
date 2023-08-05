package models
import (
	"time"
)
type Process struct {
	Id          int64   `gorm:"primaryKey"`
	Creatorid        int  `gorm:"creatorid"`
	Contractid        int  `gorm:"coontractid"`       
	Amount       float64 `gorm:"amount"`
	Created     time.Time   `gorm:"created"` 
	Remarks       string  `gorm:"remarks"`
}

// type ProcessCreateParam struct {
// 	Name        string  `json:"name" binding:"required"`
// 	Type        int     `json:"type" binding:"required,len=1"`
// 	Unit        string  `json:"unit" binding:"-"`
// 	Code        string  `json:"code" binding:"-"`
// 	Price       float64 `json:"price" binding:"required,gt=0"`
// 	Description string  `json:"description" binding:"-"`
// 	Status      int     `json:"status" binding:"required,oneof=1 2"`
// 	Creator     int64   `json:"creator,omitempty" binding:"-"`
// }

// type ProcessUpdateParam struct {
// 	Id          int64   `json:"id" binding:"required,gt=0"`
// 	Name        string  `json:"name" binding:"required"`
// 	Type        int     `json:"type" binding:"required,len=1"`
// 	Unit        string  `json:"unit" binding:"-"`
// 	Code        string  `json:"code" binding:"-"`
// 	Price       float64 `json:"price" binding:"required,gt=0"`
// 	Description string  `json:"description" binding:"-"`
// 	Status      int     `json:"status" binding:"required,oneof=1 2"`
// 	Creator     int64   `json:"creator,omitempty" binding:"-"`
// }

// type ProcessDeleteParam struct {
// 	Ids []int64 `json:"ids" binding:"required"`
// }

type ProcessQueryParam struct {
	Id      int64   `form:"id" binding:"omitempty,gt=0"`
	Creatorid    int  `form:"creatorid" binding:"-"`
	Contractid int  `form:"contractid" binding:"-"`
	Page    Page
}

type ProcessList struct {
	Id          int64   `json:"id"`
	Creatorid        int  `json:"creatorid"`
	Contractid        int  `json:"coontractid"`       
	Amount       float64 `json:"amount"`
	Created     time.Time   `json:"created"` 
	Remarks       string  `json:"remarks"`
}

type ProcessInfo struct {
	Id          int64   `json:"id"`
	Name        string  `json:"name"`
	Type        int     `json:"type"`
	Unit        string  `json:"unit"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      int     `json:"status"`
}

type ProcessExcelRow struct {
	Name        string  `json:"name"`
	Type        string  `json:"type"`
	Unit        string  `json:"unit"`
	Code        string  `json:"code"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Status      string  `json:"status"`
	Created     string  `json:"created"`
	Updated     string  `json:"updated"`
}