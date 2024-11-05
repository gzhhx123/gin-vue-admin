package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct{}

// InitProductRouter 初始化 产品信息 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	productRouter := Router.Group("product").Use(middleware.OperationRecord())
	productRouterWithoutRecord := Router.Group("product")
	productRouterWithoutAuth := PublicRouter.Group("product")
	{
		productRouter.POST("createProduct", productApi.CreateProduct)             // 新建产品信息
		productRouter.DELETE("deleteProduct", productApi.DeleteProduct)           // 删除产品信息
		productRouter.DELETE("deleteProductByIds", productApi.DeleteProductByIds) // 批量删除产品信息
		productRouter.PUT("updateProduct", productApi.UpdateProduct)              // 更新产品信息
	}
	{
		productRouterWithoutRecord.GET("findProduct", productApi.FindProduct)       // 根据ID获取产品信息
		productRouterWithoutRecord.GET("getProductList", productApi.GetProductList) // 获取产品信息列表
	}
	{
		productRouterWithoutAuth.GET("getProductDataSource", productApi.GetProductDataSource) // 获取产品信息数据源
		productRouterWithoutAuth.GET("getProductPublic", productApi.GetProductPublic)         // 产品信息开放接口
	}
}
