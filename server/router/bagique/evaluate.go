package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EvaluateRouter struct{}

// InitEvaluateRouter 初始化 估价信息 路由信息
func (s *EvaluateRouter) InitEvaluateRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	evaluateRouter := Router.Group("evaluate").Use(middleware.OperationRecord())
	evaluateRouterWithoutRecord := Router.Group("evaluate")
	evaluateRouterWithoutAuth := PublicRouter.Group("evaluate")
	{
		evaluateRouter.POST("createEvaluate", evaluateApi.CreateEvaluate)             // 新建估价信息
		evaluateRouter.DELETE("deleteEvaluate", evaluateApi.DeleteEvaluate)           // 删除估价信息
		evaluateRouter.DELETE("deleteEvaluateByIds", evaluateApi.DeleteEvaluateByIds) // 批量删除估价信息
		evaluateRouter.PUT("updateEvaluate", evaluateApi.UpdateEvaluate)              // 更新估价信息
	}
	{
		evaluateRouterWithoutRecord.GET("findEvaluate", evaluateApi.FindEvaluate)       // 根据ID获取估价信息
		evaluateRouterWithoutRecord.GET("getEvaluateList", evaluateApi.GetEvaluateList) // 获取估价信息列表
	}
	{
		evaluateRouterWithoutAuth.GET("getEvaluateDataSource", evaluateApi.GetEvaluateDataSource) // 获取估价信息数据源
		evaluateRouterWithoutAuth.GET("getEvaluatePublic", evaluateApi.GetEvaluatePublic)         // 估价信息开放接口
	}
}
