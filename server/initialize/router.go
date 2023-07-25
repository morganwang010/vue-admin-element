package initialize

import (
	"crm/api"
	"crm/global"
	"crm/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Router() {
	engine := gin.Default()

	// 开启跨域
	engine.Use(middleware.Cors())

	route := engine.Group("/api/v1")

	{
		// 用户模块，订阅模块
		route.GET("/user/verifycode", api.NewUserApi().GetVerifyCode)
		route.GET("/user/info", api.NewUserApi().GetInfo)
		route.POST("/user/login", api.NewUserApi().Login)
		route.POST("/user/register", api.NewUserApi().Register)
		route.POST("/user/pass", api.NewUserApi().ForgotPass)
		route.DELETE("/user/delete", api.NewUserApi().Delete)
		route.POST("/subscribe/payback", api.NewSubscribeApi().PayBack)

		// 通用接口
		route.POST("/common/database/init", api.NewCommonApi().InitDatabase)
		route.POST("/common/file/upload", api.NewCommonApi().FileUpload)
		route.DELETE("/common/file/remove", api.NewCommonApi().FileRemove)

		// Jwt中间件
		route.Use(middleware.JwtAuth())

		// 客户模块
		route.GET("/customer/list", api.NewCustomerApi().GetList)
		route.GET("/customer/info", api.NewCustomerApi().GetInfo)
		route.GET("/customer/option", api.NewCustomerApi().GetOption)
		route.GET("/customer/export", api.NewCustomerApi().Export)
		route.POST("/customer/create", api.NewCustomerApi().Create)
		route.POST("/customer/send", api.NewCustomerApi().SendMail)
		route.PUT("/customer/update", api.NewCustomerApi().Update)
		route.DELETE("/customer/delete", api.NewCustomerApi().Delete)

		// 合同模块
		route.GET("/contract/list", api.NewContractApi().GetList)
		route.GET("/contract/info", api.NewContractApi().GetInfo)
		route.GET("/contract/export", api.NewContractApi().Export)
		route.POST("/contract/plist", api.NewContractApi().GetProductList)
		route.PUT("/contract/update", api.NewContractApi().Update)
		route.POST("/contract/create", api.NewContractApi().Create)
		route.DELETE("/contract/delete", api.NewContractApi().Delete)

		// 产品模块
		route.GET("/product/list", api.NewProductApi().GetList)
		route.GET("/product/info", api.NewProductApi().GetInfo)
		route.GET("/product/export", api.NewProductApi().Export)
		route.POST("/product/create", api.NewProductApi().Create)
		route.PUT("/product/update", api.NewProductApi().Update)
		route.DELETE("/product/delete", api.NewProductApi().Delete)

		// 仪表盘模块
		route.GET("/dashboard/sum", api.NewDashboardApi().Summary)

		// 配置模块
		route.GET("/config/info", api.NewMailConfigApi().GetInfo)
		route.GET("/config/check", api.NewMailConfigApi().Check)
		route.PUT("/config/save", api.NewMailConfigApi().Save)
		route.PUT("/config/status", api.NewMailConfigApi().UpdateStatus)
		route.DELETE("/config/delete", api.NewMailConfigApi().Delete)

		// 订阅模块
		route.GET("/subscribe/info", api.NewSubscribeApi().GetInfo)
		route.POST("/subscribe/pay", api.NewSubscribeApi().Pay)

		// 通知模块
		route.GET("/notice/list", api.NewNoticeApi().GetList)
		route.GET("/notice/count", api.NewNoticeApi().GetUnReadCount)
		route.PUT("/notice/update", api.NewNoticeApi().Update)
		route.DELETE("/notice/delete", api.NewNoticeApi().Delete)
	}

	// 启动、监听端口
	_ = engine.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}
