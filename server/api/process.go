package api

import (
	"vue-admin-element/models"
	"vue-admin-element/response"
	"vue-admin-element/service"
	"strconv"
	log "github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
)

type ProcessApi struct {
	processService *service.ProcessService
}

func NewProcessApi() *ProcessApi {
	processApi := ProcessApi{
		processService: &service.ProcessService{},
	}
	return &processApi
}

// // 创建产品
// func (p *ProcessApi) Create(context *gin.Context) {
// 	var param models.ProcessCreateParam
// 	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
// 	err := context.ShouldBind(&param)
// 	if uid <= 0 || err != nil {
// 		response.Result(response.ErrCodeParamInvalid, nil, context)
// 		return
// 	}
// 	param.Creator = int64(uid)
// 	errCode := p.processService.Create(&param)
// 	response.Result(errCode, nil, context)
// }

// // 更新产品
// func (p *ProcessApi) Update(context *gin.Context) {
// 	var param models.ProcessUpdateParam
// 	if err := context.ShouldBind(&param); err != nil {
// 		response.Result(response.ErrCodeParamInvalid, nil, context)
// 		return
// 	}
// 	errCode := p.processService.Update(&param)
// 	response.Result(errCode, nil, context)
// }

// // 删除产品
// func (p *ProcessApi) Delete(context *gin.Context) {
// 	var param models.ProcessDeleteParam
// 	if err := context.ShouldBind(&param); err != nil {
// 		response.Result(response.ErrCodeParamInvalid, nil, context)
// 		return
// 	}
// 	errCode := p.processService.Delete(&param)
// 	response.Result(errCode, nil, context)
// }

// 查询产品列表
func (p *ProcessApi) GetList(context *gin.Context) {
	var param models.ProcessQueryParam
	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
	err := context.ShouldBind(&param)
	if uid <= 0 || err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	processList, rows, errCode := p.processService.GetList(&param)
	response.PageResult(errCode, processList, rows, context)
}

// // 查询产品信息
// func (p *ProcessApi) GetInfo(context *gin.Context) {
// 	var param models.ProcessQueryParam
// 	if err := context.ShouldBind(&param); err != nil {
// 		response.Result(response.ErrCodeParamInvalid, nil, context)
// 		return
// 	}
// 	processInfo, errCode := p.processService.GetInfo(&param)
// 	response.Result(errCode, processInfo, context)
// }

// // 导出Excel文件
// func (p *ProcessApi) Export(context *gin.Context) {
// 	uid, _ := strconv.Atoi(context.Request.Header.Get("uid"))
// 	if uid <= 0 {
// 		response.Result(response.ErrCodeParamInvalid, nil, context)
// 		return
// 	}
// 	file, errCode := p.processService.Export(int64(uid))
// 	if len(file) >= 0 && errCode != 0 {
// 		response.Result(errCode, nil, context)
// 		return
// 	}
// 	context.File(file)
// }
