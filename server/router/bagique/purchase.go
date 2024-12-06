package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PurchaseRouter struct{}

func (s *PurchaseRouter) InitPurchaseRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	purchaseRouter := Router.Group("purchase").Use(middleware.OperationRecord())
	purchaseRouterWithoutRecord := Router.Group("purchase")
	purchaseRouterWithoutAuth := PublicRouter.Group("purchase")
	{
		purchaseRouter.POST("createPurchase", purchaseApi.CreatePurchase)
		purchaseRouter.DELETE("deletePurchase", purchaseApi.DeletePurchase)
		purchaseRouter.DELETE("deletePurchaseByIds", purchaseApi.DeletePurchaseByIds)
		purchaseRouter.PUT("updatePurchase", purchaseApi.UpdatePurchase)
		purchaseRouter.PUT("refuseEvaluate", purchaseApi.RefuseEvaluate)
	}
	{
		purchaseRouterWithoutRecord.GET("findPurchase", purchaseApi.FindPurchase)
		purchaseRouterWithoutRecord.GET("getPurchaseList", purchaseApi.GetPurchaseList)
	}
	{
		purchaseRouterWithoutAuth.GET("getPurchaseDataSource", purchaseApi.GetPurchaseDataSource)
		purchaseRouterWithoutAuth.GET("getPurchasePublic", purchaseApi.GetPurchasePublic)
	}
}
