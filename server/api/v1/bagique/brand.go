package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BrandApi struct{}

// CreateBrand 创建品牌信息
// @Tags Brand
// @Summary 创建品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Brand true "创建品牌信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /brand/createBrand [post]
func (brandApi *BrandApi) CreateBrand(c *gin.Context) {
	var brand bagique.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = brandService.CreateBrand(&brand)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteBrand 删除品牌信息
// @Tags Brand
// @Summary 删除品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Brand true "删除品牌信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /brand/deleteBrand [delete]
func (brandApi *BrandApi) DeleteBrand(c *gin.Context) {
	ID := c.Query("ID")
	err := brandService.DeleteBrand(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteBrandByIds 批量删除品牌信息
// @Tags Brand
// @Summary 批量删除品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /brand/deleteBrandByIds [delete]
func (brandApi *BrandApi) DeleteBrandByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := brandService.DeleteBrandByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateBrand 更新品牌信息
// @Tags Brand
// @Summary 更新品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Brand true "更新品牌信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /brand/updateBrand [put]
func (brandApi *BrandApi) UpdateBrand(c *gin.Context) {
	var brand bagique.Brand
	err := c.ShouldBindJSON(&brand)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = brandService.UpdateBrand(brand)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindBrand 用id查询品牌信息
// @Tags Brand
// @Summary 用id查询品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Brand true "用id查询品牌信息"
// @Success 200 {object} response.Response{data=bagique.Brand,msg=string} "查询成功"
// @Router /brand/findBrand [get]
func (brandApi *BrandApi) FindBrand(c *gin.Context) {
	ID := c.Query("ID")
	rebrand, err := brandService.GetBrand(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rebrand, c)
}

// GetBrandList 分页获取品牌信息列表
// @Tags Brand
// @Summary 分页获取品牌信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.BrandSearch true "分页获取品牌信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /brand/getBrandList [get]
func (brandApi *BrandApi) GetBrandList(c *gin.Context) {
	var pageInfo bagiqueReq.BrandSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := brandService.GetBrandInfoList(pageInfo)
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

// GetBrandPublic 不需要鉴权的品牌信息接口
// @Tags Brand
// @Summary 不需要鉴权的品牌信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.BrandSearch true "分页获取品牌信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /brand/getBrandPublic [get]
func (brandApi *BrandApi) GetBrandPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	brandService.GetBrandPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的品牌信息接口信息",
	}, "获取成功", c)
}
