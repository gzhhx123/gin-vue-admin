package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type BrandRouter struct{}

func (s *BrandRouter) InitBrandRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	brandRouter := Router.Group("brand").Use(middleware.OperationRecord())
	brandRouterWithoutRecord := Router.Group("brand")
	brandRouterWithoutAuth := PublicRouter.Group("brand")
	{
		brandRouter.POST("createBrand", brandApi.CreateBrand)
		brandRouter.DELETE("deleteBrand", brandApi.DeleteBrand)
		brandRouter.DELETE("deleteBrandByIds", brandApi.DeleteBrandByIds)
		brandRouter.PUT("updateBrand", brandApi.UpdateBrand)
		brandRouter.PUT("restoreBrand", brandApi.RestoreBrand)
	}
	{
		brandRouterWithoutRecord.GET("findBrand", brandApi.FindBrand)
		brandRouterWithoutRecord.GET("getBrandList", brandApi.GetBrandList)
	}
	{
		brandRouterWithoutAuth.GET("getBrandPublic", brandApi.GetBrandPublic)
	}
}
