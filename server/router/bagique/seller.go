package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SellerRouter struct{}

// InitSellerRouter 初始化 卖家信息 路由信息
func (s *SellerRouter) InitSellerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	sellerRouter := Router.Group("seller").Use(middleware.OperationRecord())
	sellerRouterWithoutRecord := Router.Group("seller")
	sellerRouterWithoutAuth := PublicRouter.Group("seller")
	{
		sellerRouter.POST("createSeller", sellerApi.CreateSeller)             // 新建卖家信息
		sellerRouter.DELETE("deleteSeller", sellerApi.DeleteSeller)           // 删除卖家信息
		sellerRouter.DELETE("deleteSellerByIds", sellerApi.DeleteSellerByIds) // 批量删除卖家信息
		sellerRouter.PUT("updateSeller", sellerApi.UpdateSeller)              // 更新卖家信息
	}
	{
		sellerRouterWithoutRecord.GET("findSeller", sellerApi.FindSeller)       // 根据ID获取卖家信息
		sellerRouterWithoutRecord.GET("getSellerList", sellerApi.GetSellerList) // 获取卖家信息列表
	}
	{
		sellerRouterWithoutAuth.GET("getSellerPublic", sellerApi.GetSellerPublic) // 卖家信息开放接口
	}
}
