package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouter struct{}

func (s *CompanyRouter) InitCompanyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	companyRouter := Router.Group("company").Use(middleware.OperationRecord())
	companyRouterWithoutRecord := Router.Group("company")
	companyRouterWithoutAuth := PublicRouter.Group("company")
	{
		companyRouter.POST("createCompany", companyApi.CreateCompany)
		companyRouter.DELETE("deleteCompany", companyApi.DeleteCompany)
		companyRouter.DELETE("deleteCompanyByIds", companyApi.DeleteCompanyByIds)
		companyRouter.PUT("updateCompany", companyApi.UpdateCompany)
		companyRouter.PUT("restoreCompany", companyApi.RestoreCompany)
	}
	{
		companyRouterWithoutRecord.GET("findCompany", companyApi.FindCompany)
		companyRouterWithoutRecord.GET("getCompanyList", companyApi.GetCompanyList)
	}
	{
		companyRouterWithoutAuth.GET("getCompanyPublic", companyApi.GetCompanyPublic)
	}
}
