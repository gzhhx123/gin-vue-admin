package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BrandRouter struct{}

// InitBrandRouter 初始化 品牌信息 路由信息
func (s *BrandRouter) InitBrandRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	brandRouter := Router.Group("brand").Use(middleware.OperationRecord())
	brandRouterWithoutRecord := Router.Group("brand")
	brandRouterWithoutAuth := PublicRouter.Group("brand")
	{
		brandRouter.POST("createBrand", brandApi.CreateBrand)             // 新建品牌信息
		brandRouter.DELETE("deleteBrand", brandApi.DeleteBrand)           // 删除品牌信息
		brandRouter.DELETE("deleteBrandByIds", brandApi.DeleteBrandByIds) // 批量删除品牌信息
		brandRouter.PUT("updateBrand", brandApi.UpdateBrand)              // 更新品牌信息
	}
	{
		brandRouterWithoutRecord.GET("findBrand", brandApi.FindBrand)       // 根据ID获取品牌信息
		brandRouterWithoutRecord.GET("getBrandList", brandApi.GetBrandList) // 获取品牌信息列表
	}
	{
		brandRouterWithoutAuth.GET("getBrandPublic", brandApi.GetBrandPublic) // 品牌信息开放接口
	}
}
