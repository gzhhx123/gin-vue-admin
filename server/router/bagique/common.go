package bagique

import (
	"github.com/gin-gonic/gin"
)

type CommonRouter struct{}

// InitCommonRouter 初始化 通用 路由信息
func (s *CommonRouter) InitCommonRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	commonRouterWithoutRecord := Router.Group("common")
	{
		commonRouterWithoutRecord.GET("getRate", commonApi.GetRate) // 获取美元汇率
	}
}
