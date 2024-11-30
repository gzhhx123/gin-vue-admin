package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TrackCompanyRouter struct{}

// InitTrackCompanyRouter 初始化 物流公司 路由信息
func (s *TrackCompanyRouter) InitTrackCompanyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	trackCompanyRouter := Router.Group("trackCompany").Use(middleware.OperationRecord())
	trackCompanyRouterWithoutRecord := Router.Group("trackCompany")
	trackCompanyRouterWithoutAuth := PublicRouter.Group("trackCompany")
	{
		trackCompanyRouter.POST("createTrackCompany", trackCompanyApi.CreateTrackCompany)             // 新建物流公司
		trackCompanyRouter.DELETE("deleteTrackCompany", trackCompanyApi.DeleteTrackCompany)           // 删除物流公司
		trackCompanyRouter.DELETE("deleteTrackCompanyByIds", trackCompanyApi.DeleteTrackCompanyByIds) // 批量删除物流公司
		trackCompanyRouter.PUT("updateTrackCompany", trackCompanyApi.UpdateTrackCompany)              // 更新物流公司
	}
	{
		trackCompanyRouterWithoutRecord.GET("findTrackCompany", trackCompanyApi.FindTrackCompany)       // 根据ID获取物流公司
		trackCompanyRouterWithoutRecord.GET("getTrackCompanyList", trackCompanyApi.GetTrackCompanyList) // 获取物流公司列表
	}
	{
		trackCompanyRouterWithoutAuth.GET("getTrackCompanyPublic", trackCompanyApi.GetTrackCompanyPublic) // 物流公司开放接口
	}
}
