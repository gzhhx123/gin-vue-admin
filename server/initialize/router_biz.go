package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	{
		bagiqueRouter := router.RouterGroupApp.Bagique
		bagiqueRouter.InitBrandRouter(privateGroup, publicGroup)
		bagiqueRouter.InitProductRouter(privateGroup, publicGroup)
		bagiqueRouter.InitCompanyRouter(privateGroup, publicGroup)
		bagiqueRouter.InitSellerRouter(privateGroup, publicGroup)
		bagiqueRouter.InitEvaluateRouter(privateGroup, publicGroup) // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
		bagiqueRouter.InitEvaluatePriceRouter(privateGroup, publicGroup)
	}
}
