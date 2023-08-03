package service

import (
	// "vue-admin-element/common"
	"vue-admin-element/dao"
	"vue-admin-element/models"
	"vue-admin-element/response"
	// "strconv"
)

type UrlsService struct {
	urlsDao *dao.UrlsDao
}

func NewUrlsService() *UrlsService {
	return &UrlsService{
		urlsDao: dao.NewUrlsDao(),
	}
}

// // 创建产品
// func (p *ProductService) Create(param *models.ProductCreateParam) int {
// 	if p.productDao.IsExists(param.Name, param.Creator) {
// 		return response.ErrCodeProductHasExist
// 	}
// 	if err := p.productDao.Create(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// // 更新产品
// func (p *ProductService) Update(param *models.ProductUpdateParam) int {
// 	if err := p.productDao.Update(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// // 删除产品
// func (p *ProductService) Delete(param *models.ProductDeleteParam) int {
// 	if err := p.productDao.Delete(param); err != nil {
// 		return response.ErrCodeFailed
// 	}
// 	return response.ErrCodeSuccess
// }

// 查询产品列表
func (p *UrlsService) GetList(param *models.UrlsQueryParam) ([]*models.UrlsList, int64, int) {
	urlsList, rows, err := p.urlsDao.GetList(param)
	if err != nil {
		return nil, NumberNull, response.ErrCodeFailed
	}
	return urlsList, rows, response.ErrCodeSuccess
}

// // 查询产品信息
// func (p *ProductService) GetInfo(param *models.ProductQueryParam) (*models.ProductInfo, int) {
// 	productInfo, err := p.productDao.GetInfo(param)
// 	if err != nil {
// 		return nil, response.ErrCodeFailed
// 	}
// 	return productInfo, response.ErrCodeSuccess
// }

// // 导出Excel文件
// func (p *ProductService) Export(uid int64) (string, int) {
// 	products, err := p.productDao.GetListByUid(uid)
// 	if err != nil {
// 		return StringNull, response.ErrCodeFileExportFailed
// 	}
// 	excelRows := make([]models.ProductExcelRow, 0)
// 	var row models.ProductExcelRow
// 	for _, p := range products {
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
// 	fileName := "product_" + strconv.FormatInt(uid, 10)
// 	file, err := common.GenExcelFile(sheet, columns, excelRows, fileName)
// 	if err != nil {
// 		return StringNull, response.ErrCodeFileExportFailed
// 	}
// 	return file, response.ErrCodeSuccess
// }
