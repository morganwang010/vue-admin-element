package api

import(
	"vue-admin-element/service"
	"github.com/gin-gonic/gin"
	"vue-admin-element/response"

)

type DictApi struct {
	dictService *service.DictService

}

func NewDictApi() *DictApi {
	dictApi := DictApi{
		dictService: service.NewDictService(),
	}
	return &dictApi
}


func (u *DictApi) GetList(context *gin.Context) {
	 
    response.Result(response.ErrCodeSuccess, nil, context)
	}
