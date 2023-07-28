package service
import (
	"vue-admin-element/dao"
)

type DictService struct {
	dictDao      *dao.DictDao
 
}

func NewDictService() *DictService {
	dictService := DictService{
		dictDao:      dao.NewDictDao(),
 
	}
	return &dictService
}

