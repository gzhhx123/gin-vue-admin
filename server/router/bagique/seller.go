package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SellerRouter struct{}

func (s *SellerRouter) InitSellerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sellerRouter := Router.Group("seller").Use(middleware.OperationRecord())
	sellerRouterWithoutRecord := Router.Group("seller")
	sellerRouterWithoutAuth := PublicRouter.Group("seller")
	{
		sellerRouter.POST("createSeller", sellerApi.CreateSeller)
		sellerRouter.DELETE("deleteSeller", sellerApi.DeleteSeller)
		sellerRouter.DELETE("deleteSellerByIds", sellerApi.DeleteSellerByIds)
		sellerRouter.PUT("updateSeller", sellerApi.UpdateSeller)
		sellerRouter.PUT("restoreSeller", sellerApi.RestoreSeller)
	}
	{
		sellerRouterWithoutRecord.GET("findSeller", sellerApi.FindSeller)
		sellerRouterWithoutRecord.GET("getSellerList", sellerApi.GetSellerList)
	}
	{
		sellerRouterWithoutAuth.GET("getSellerPublic", sellerApi.GetSellerPublic)
	}
}
