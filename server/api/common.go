package api

import (
	"vue-admin-element/models"
	"vue-admin-element/response"
	"vue-admin-element/service"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type CommonApi struct {
	commonService *service.CommonService
}

func NewCommonApi() *CommonApi {
	return &CommonApi{
		commonService: service.NewCommonService(),
	}
}

// 初始化数据库
func (c *CommonApi) InitDatabase(context *gin.Context) {
	errCode := c.commonService.InitDatabase()
	response.Result(errCode, nil, context)
}

// 文件上传
func (c *CommonApi) FileUpload(context *gin.Context) {
	file, _ := context.FormFile("file")
	// log.Println(file)
	fileInfo, errCode := c.commonService.FileUpload(file)
	response.Result(errCode, fileInfo, context)
}

// 文件移除
func (c *CommonApi) FileRemove(context *gin.Context) {
	var param models.FileParam
	if err := context.ShouldBind(&param); err != nil {
		response.Result(response.ErrCodeParamInvalid, nil, context)
		log.Println(err)
		return
	}
	errCode := c.commonService.FileRemove(&param)
	response.Result(errCode, nil, context)
}
