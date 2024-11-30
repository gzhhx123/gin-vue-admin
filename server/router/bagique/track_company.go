package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type TrackCompanyRouter struct{}

func (s *TrackCompanyRouter) InitTrackCompanyRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	trackCompanyRouter := Router.Group("trackCompany").Use(middleware.OperationRecord())
	trackCompanyRouterWithoutRecord := Router.Group("trackCompany")
	trackCompanyRouterWithoutAuth := PublicRouter.Group("trackCompany")
	{
		trackCompanyRouter.POST("createTrackCompany", trackCompanyApi.CreateTrackCompany)
		trackCompanyRouter.DELETE("deleteTrackCompany", trackCompanyApi.DeleteTrackCompany)
		trackCompanyRouter.DELETE("deleteTrackCompanyByIds", trackCompanyApi.DeleteTrackCompanyByIds)
		trackCompanyRouter.PUT("updateTrackCompany", trackCompanyApi.UpdateTrackCompany)
		trackCompanyRouter.PUT("restoreTrackCompany", trackCompanyApi.RestoreTrackCompany)
	}
	{
		trackCompanyRouterWithoutRecord.GET("findTrackCompany", trackCompanyApi.FindTrackCompany)
		trackCompanyRouterWithoutRecord.GET("getTrackCompanyList", trackCompanyApi.GetTrackCompanyList)
	}
	{
		trackCompanyRouterWithoutAuth.GET("getTrackCompanyPublic", trackCompanyApi.GetTrackCompanyPublic)
	}
}
