package dao

import (
	"vue-admin-element/global"
	"vue-admin-element/models"
	// "time"
	log "github.com/sirupsen/logrus"
)

type ProcessDao struct {
}

func NewProcessDao() *ProcessDao {
	return &ProcessDao{}
}

// func (p *ProcessDao) Create(param *models.ProcessCreateParam) error {
// 	process := models.Process{
// 		Name:        param.Name,
// 		Type:        param.Type,
// 		Unit:        param.Unit,
// 		Code:        param.Code,
// 		Price:       param.Price,
// 		Description: param.Description,
// 		Status:      param.Status,
// 		Creator:     param.Creator,
// 		Created:     time.Now(),
// 	}
// 	return global.Db.Create(&process).Error
// }

// func (p *ProcessDao) Update(param *models.ProcessUpdateParam) error {
// 	process := models.Process{
// 		Id:          param.Id,
// 		Name:        param.Name,
// 		Type:        param.Type,
// 		Unit:        param.Unit,
// 		Code:        param.Code,
// 		Price:       param.Price,
// 		Description: param.Description,
// 		Status:      param.Status,
// 		Updated:     time.Now(),
// 	}
// 	db := global.Db.Model(&process).Select("*").Omit("id", "creator", "created")
// 	return db.Updates(&process).Error
// }

// func (p *ProcessDao) Delete(param *models.ProcessDeleteParam) error {
// 	return global.Db.Delete(&models.Process{}, param.Ids).Error
// }

func (p *ProcessDao) IsExists(name string, uid int64) bool {
	var process models.Process
	db := global.Db.Table(PROCESS).Where("name = ? and creator = ?", name, uid).First(&process)
	return db.RowsAffected != NumberNull
}

func (p *ProcessDao) GetList(param *models.ProcessQueryParam) ([]*models.ProcessList, int64, error) {
	process := models.Process{
		Id:    param.Id,
		Contractid:  param.Contractid,
		Creatorid: param.Creatorid,
	}
	processList := make([]*models.ProcessList, 0)
	rows, err := restPage(param.Page, PROCESS, process, &processList, &[]*models.ProcessList{})
	log.Println(processList)
	if err != nil {
		return nil, 0, err
	}
	return processList, rows, err
}

// func (p *ProcessDao) GetListByIds(ids []int64) ([]*models.Processs, error) {
// 	processs := make([]*models.Processs, 0)
// 	if err := global.Db.Table(PROCESS).Find(&processs, ids).Error; err != nil {
// 		return nil, err
// 	}
// 	return processs, nil
// }

// func (p *ProcessDao) GetInfo(param *models.ProcessQueryParam) (*models.ProcessInfo, error) {
// 	process := models.Process{
// 		Id: param.Id,
// 	}
// 	processInfo := models.ProcessInfo{}
// 	err := global.Db.Table(PROCESS).Where(&process).First(&processInfo).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &processInfo, err
// }

// func (p *ProcessDao) GetListByUid(uid int64) ([]*models.Process, error) {
// 	processs := make([]*models.Process, 0)
// 	err := global.Db.Where("creator = ?", uid).Find(&processs).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return processs, nil
// }
