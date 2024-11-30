package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PurchaseRouter struct{}

// InitPurchaseRouter 初始化 采购信息 路由信息
func (s *PurchaseRouter) InitPurchaseRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	purchaseRouter := Router.Group("purchase").Use(middleware.OperationRecord())
	purchaseRouterWithoutRecord := Router.Group("purchase")
	purchaseRouterWithoutAuth := PublicRouter.Group("purchase")
	{
		purchaseRouter.POST("createPurchase", purchaseApi.CreatePurchase)             // 新建采购信息
		purchaseRouter.DELETE("deletePurchase", purchaseApi.DeletePurchase)           // 删除采购信息
		purchaseRouter.DELETE("deletePurchaseByIds", purchaseApi.DeletePurchaseByIds) // 批量删除采购信息
		purchaseRouter.PUT("updatePurchase", purchaseApi.UpdatePurchase)              // 更新采购信息
	}
	{
		purchaseRouterWithoutRecord.GET("findPurchase", purchaseApi.FindPurchase)       // 根据ID获取采购信息
		purchaseRouterWithoutRecord.GET("getPurchaseList", purchaseApi.GetPurchaseList) // 获取采购信息列表
	}
	{
		purchaseRouterWithoutAuth.GET("getPurchasePublic", purchaseApi.GetPurchasePublic) // 采购信息开放接口
	}
}
