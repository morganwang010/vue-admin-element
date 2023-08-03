package models
import (
	"time"
)
type Urls struct {
	Id          int64   `gorm:"primaryKey"`
	Url        string  `gorm:"name"`
	Description string  `gorm:"description"`
	Keywords        string     `gorm:"keywords"`
	Status      int     `gorm:"status"`
	Views      int     `gorm:"views"`
	Like      int     `gorm:"like"`
	Created     time.Time   `gorm:"created"`
	Updated     time.Time   `gorm:"updated"`
	remarks      string  `gorm:"remarks"` 

}

// type ProductCreateParam struct {
// 	Name        string  `json:"name" binding:"required"`
// 	Type        int     `json:"type" binding:"required,len=1"`
// 	Unit        string  `json:"unit" binding:"-"`
// 	Code        string  `json:"code" binding:"-"`
// 	Price       float64 `json:"price" binding:"required,gt=0"`
// 	Description string  `json:"description" binding:"-"`
// 	Status      int     `json:"status" binding:"required,oneof=1 2"`
// 	Creator     int64   `json:"creator,omitempty" binding:"-"`
// }

// type ProductUpdateParam struct {
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

// type ProductDeleteParam struct {
// 	Ids []int64 `json:"ids" binding:"required"`
// }

type UrlsQueryParam struct {
	Keywords    string  `form:"keywords" binding:"-"`
	Status  int     `form:"status" binding:"omitempty"`
	Page    Page
}

type UrlsList struct {
	Id          int64   `json:"id"`
	Url        string  `json:"url"`
	Description string  `json:"description"`
	Keywords        string     `json:"keywords"`
	Status      int     `json:"status"`
	Views      int     `json:"views"`
	Like      int     `json:"like"`
	Created     time.Time   `json:"created"`
	Updated     time.Time   `json:"updated"`
	remarks       string  `json:"remarks"` 

}

// type ProductInfo struct {
// 	Id          int64   `json:"id"`
// 	Name        string  `json:"name"`
// 	Type        int     `json:"type"`
// 	Unit        string  `json:"unit"`
// 	Code        string  `json:"code"`
// 	Price       float64 `json:"price"`
// 	Description string  `json:"description"`
// 	Status      int     `json:"status"`
// }

// type ProductExcelRow struct {
// 	Name        string  `json:"name"`
// 	Type        string  `json:"type"`
// 	Unit        string  `json:"unit"`
// 	Code        string  `json:"code"`
// 	Price       float64 `json:"price"`
// 	Description string  `json:"description"`
// 	Status      string  `json:"status"`
// 	Created     string  `json:"created"`
// 	Updated     string  `json:"updated"`
// }
