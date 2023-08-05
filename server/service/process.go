package service

import (
	// "vue-admin-element/common"
	"vue-admin-element/dao"
	"vue-admin-element/models"
	"vue-admin-element/response"
	// "strconv"
)

type ProcessService struct {
	processDao *dao.ProcessDao
}

func NewProcessService() *ProcessService {
	return &ProcessService{
		processDao: dao.NewProcessDao(),
	}
}

// // 创建产品
// func (p *ProcessService) Create(param *models.ProcessCreateParam) int {
// 	if p.processDao.IsExists(param.Name, param.Creator) {
// 		return response.ErrCodeProcessHasExist
// 	}
// 	if err := p.processDao.Create(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// // 更新产品
// func (p *ProcessService) Update(param *models.ProcessUpdateParam) int {
// 	if err := p.processDao.Update(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// // 删除产品
// func (p *ProcessService) Delete(param *models.ProcessDeleteParam) int {
// 	if err := p.processDao.Delete(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// 查询产品列表
func (p *ProcessService) GetList(param *models.ProcessQueryParam) ([]*models.ProcessList, int64, int) {
	processList, rows, err := p.processDao.GetList(param)
	if err != nil {
		return nil, NumberNull, response.ErrCodeFailed
	}
	return processList, rows, response.ErrCodeSuccess
}

// // 查询产品信息
// func (p *ProcessService) GetInfo(param *models.ProcessQueryParam) (*models.ProcessInfo, int) {
// 	processInfo, err := p.processDao.GetInfo(param)
// 	if err != nil {
// 		return nil, response.ErrCodeFailed
// 	}
// 	return processInfo, response.ErrCodeSuccess
// }

// // 导出Excel文件
// func (p *ProcessService) Export(uid int64) (string, int) {
// 	processs, err := p.processDao.GetListByUid(uid)
// 	if err != nil {
// 		return StringNull, response.ErrCodeFileExportFailed
// 	}
// 	excelRows := make([]models.ProcessExcelRow, 0)
// 	var row models.ProcessExcelRow
// 	for _, p := range processs {
// 		row.Name = p.Name
// 		if p.Status == 1 {
// 			row.Status = "已上架"
// 		}
// 		if p.Status == 2 {
// 			row.Status = "已下架"
// 		}
// 		if p.Type == 1 {
// 			row.Type = "默认"
// 		}
// 		row.Unit = p.Unit
// 		row.Code = p.Code
// 		row.Price = p.Price
// 		row.Description = p.Description
// 		row.Created = p.Created.Format("2006-01-02")
// 		if !p.Updated.IsZero(){
// 			row.Updated = p.Updated.Format("2006-01-02")
// 		}
// 		excelRows = append(excelRows, row)
// 	}
// 	sheet := "产品信息"
// 	columns := []string{"名称", "是否上下架", "类型", "单位", "编码", "价格", "描述", "创建时间", "更新时间"}
// 	fileName := "process_" + strconv.FormatInt(uid, 10)
// 	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
// 	if err != nil {
// 		return StringNull, response.ErrCodeFileExportFailed
// 	}
// 	return file, response.ErrCodeSuccess
// }
