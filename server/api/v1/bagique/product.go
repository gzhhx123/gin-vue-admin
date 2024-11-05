package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductApi struct{}

// CreateProduct 创建产品信息
// @Tags Product
// @Summary 创建产品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Product true "创建产品信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /product/createProduct [post]
func (productApi *ProductApi) CreateProduct(c *gin.Context) {
	var product bagique.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.CreateProduct(&product)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProduct 删除产品信息
// @Tags Product
// @Summary 删除产品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Product true "删除产品信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /product/deleteProduct [delete]
func (productApi *ProductApi) DeleteProduct(c *gin.Context) {
	ID := c.Query("ID")
	err := productService.DeleteProduct(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductByIds 批量删除产品信息
// @Tags Product
// @Summary 批量删除产品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /product/deleteProductByIds [delete]
func (productApi *ProductApi) DeleteProductByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := productService.DeleteProductByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProduct 更新产品信息
// @Tags Product
// @Summary 更新产品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Product true "更新产品信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /product/updateProduct [put]
func (productApi *ProductApi) UpdateProduct(c *gin.Context) {
	var product bagique.Product
	err := c.ShouldBindJSON(&product)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productService.UpdateProduct(product)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProduct 用id查询产品信息
// @Tags Product
// @Summary 用id查询产品信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Product true "用id查询产品信息"
// @Success 200 {object} response.Response{data=bagique.Product,msg=string} "查询成功"
// @Router /product/findProduct [get]
func (productApi *ProductApi) FindProduct(c *gin.Context) {
	ID := c.Query("ID")
	reproduct, err := productService.GetProduct(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reproduct, c)
}

// GetProductList 分页获取产品信息列表
// @Tags Product
// @Summary 分页获取产品信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.ProductSearch true "分页获取产品信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /product/getProductList [get]
func (productApi *ProductApi) GetProductList(c *gin.Context) {
	var pageInfo bagiqueReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := productService.GetProductInfoList(pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetProductDataSource 获取Product的数据源
// @Tags Product
// @Summary 获取Product的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /product/getProductDataSource [get]
func (productApi *ProductApi) GetProductDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := productService.GetProductDataSource()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetProductPublic 不需要鉴权的产品信息接口
// @Tags Product
// @Summary 不需要鉴权的产品信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.ProductSearch true "分页获取产品信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /product/getProductPublic [get]
func (productApi *ProductApi) GetProductPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	productService.GetProductPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的产品信息接口信息",
	}, "获取成功", c)
}
