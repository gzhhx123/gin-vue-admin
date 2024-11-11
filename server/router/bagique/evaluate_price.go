package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EvaluatePriceRouter struct{}

// InitEvaluatePriceRouter 初始化 公司估价 路由信息
func (s *EvaluatePriceRouter) InitEvaluatePriceRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	evaluatePriceRouter := Router.Group("evaluatePrice").Use(middleware.OperationRecord())
	evaluatePriceRouterWithoutRecord := Router.Group("evaluatePrice")
	evaluatePriceRouterWithoutAuth := PublicRouter.Group("evaluatePrice")
	{
		evaluatePriceRouter.POST("createEvaluatePrice", evaluatePriceApi.CreateEvaluatePrice)             // 新建公司估价
		evaluatePriceRouter.DELETE("deleteEvaluatePrice", evaluatePriceApi.DeleteEvaluatePrice)           // 删除公司估价
		evaluatePriceRouter.DELETE("deleteEvaluatePriceByIds", evaluatePriceApi.DeleteEvaluatePriceByIds) // 批量删除公司估价
		evaluatePriceRouter.PUT("updateEvaluatePrice", evaluatePriceApi.UpdateEvaluatePrice)              // 更新公司估价
	}
	{
		evaluatePriceRouterWithoutRecord.GET("findEvaluatePrice", evaluatePriceApi.FindEvaluatePrice)       // 根据ID获取公司估价
		evaluatePriceRouterWithoutRecord.GET("getEvaluatePriceList", evaluatePriceApi.GetEvaluatePriceList) // 获取公司估价列表
	}
	{
		evaluatePriceRouterWithoutAuth.GET("getEvaluatePriceDataSource", evaluatePriceApi.GetEvaluatePriceDataSource) // 获取公司估价数据源
		evaluatePriceRouterWithoutAuth.GET("getEvaluatePricePublic", evaluatePriceApi.GetEvaluatePricePublic)         // 公司估价开放接口
	}
}
