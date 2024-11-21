package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	productRouterWithoutAuth := PublicRouter.Group("product")
	{
		productRouter.POST("createProduct", productApi.CreateProduct)
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct)
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds)
		productRouter.PUT("updateProduct", productApi.UpdateProduct)
		productRouter.PUT("restoreProduct", productApi.RestoreProduct)
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList)
	}
	{
		productRouterWithoutAuth.GET("getProductDataSource", productApi.GetProductDataSource)
		productRouterWithoutAuth.GET("getProductPublic", productApi.GetProductPublic)
	}
}
