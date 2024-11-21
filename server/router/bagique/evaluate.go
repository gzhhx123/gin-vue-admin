package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type EvaluateRouter struct{}

func (s *EvaluateRouter) InitEvaluateRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	evaluateRouter := Router.Group("evaluate").Use(middleware.OperationRecord())
	evaluateRouterWithoutRecord := Router.Group("evaluate")
	evaluateRouterWithoutAuth := PublicRouter.Group("evaluate")
	{
		evaluateRouter.POST("createEvaluate", evaluateApi.CreateEvaluate)
		evaluateRouter.DELETE("deleteEvaluate", evaluateApi.DeleteEvaluate)
		evaluateRouter.DELETE("deleteEvaluateByIds", evaluateApi.DeleteEvaluateByIds)
		evaluateRouter.PUT("updateEvaluate", evaluateApi.UpdateEvaluate)
		evaluateRouter.PUT("restoreEvaluate", evaluateApi.RestoreEvaluate)
	}
	{
		evaluateRouterWithoutRecord.GET("findEvaluate", evaluateApi.FindEvaluate)
		evaluateRouterWithoutRecord.GET("getEvaluateList", evaluateApi.GetEvaluateList)
	}
	{
		evaluateRouterWithoutAuth.GET("getEvaluateDataSource", evaluateApi.GetEvaluateDataSource)
		evaluateRouterWithoutAuth.GET("getEvaluatePublic", evaluateApi.GetEvaluatePublic)
	}
}
