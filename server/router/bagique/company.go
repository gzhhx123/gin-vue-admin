package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouter struct{}

// InitCompanyRouter 初始化 公司信息 路由信息
func (s *CompanyRouter) InitCompanyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	companyRouter := Router.Group("company").Use(middleware.OperationRecord())
	companyRouterWithoutRecord := Router.Group("company")
	companyRouterWithoutAuth := PublicRouter.Group("company")
	{
		companyRouter.POST("createCompany", companyApi.CreateCompany)             // 新建公司信息
		companyRouter.DELETE("deleteCompany", companyApi.DeleteCompany)           // 删除公司信息
		companyRouter.DELETE("deleteCompanyByIds", companyApi.DeleteCompanyByIds) // 批量删除公司信息
		companyRouter.PUT("updateCompany", companyApi.UpdateCompany)              // 更新公司信息
	}
	{
		companyRouterWithoutRecord.GET("findCompany", companyApi.FindCompany)       // 根据ID获取公司信息
		companyRouterWithoutRecord.GET("getCompanyList", companyApi.GetCompanyList) // 获取公司信息列表
	}
	{
		companyRouterWithoutAuth.GET("getCompanyPublic", companyApi.GetCompanyPublic) // 公司信息开放接口
	}
}
