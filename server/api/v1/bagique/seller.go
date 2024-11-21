package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SellerApi struct{}

// CreateSeller 创建卖家信息
// @Tags Seller
// @Summary 创建卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Seller true "创建卖家信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /seller/createSeller [post]
func (sellerApi *SellerApi) CreateSeller(c *gin.Context) {
	var seller bagique.Seller
	err := c.ShouldBindJSON(&seller)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sellerService.CreateSeller(&seller)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteSeller 删除卖家信息
// @Tags Seller
// @Summary 删除卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Seller true "删除卖家信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /seller/deleteSeller [delete]
func (sellerApi *SellerApi) DeleteSeller(c *gin.Context) {
	ID := c.Query("ID")
	TYPE := c.Query("TYPE")
	err := sellerService.DeleteSeller(ID, TYPE)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteSellerByIds 批量删除卖家信息
// @Tags Seller
// @Summary 批量删除卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /seller/deleteSellerByIds [delete]
func (sellerApi *SellerApi) DeleteSellerByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := sellerService.DeleteSellerByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateSeller 更新卖家信息
// @Tags Seller
// @Summary 更新卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Seller true "更新卖家信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /seller/updateSeller [put]
func (sellerApi *SellerApi) UpdateSeller(c *gin.Context) {
	var seller bagique.Seller
	err := c.ShouldBindJSON(&seller)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = sellerService.UpdateSeller(seller)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindSeller 用id查询卖家信息
// @Tags Seller
// @Summary 用id查询卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Seller true "用id查询卖家信息"
// @Success 200 {object} response.Response{data=bagique.Seller,msg=string} "查询成功"
// @Router /seller/findSeller [get]
func (sellerApi *SellerApi) FindSeller(c *gin.Context) {
	ID := c.Query("ID")
	reseller, err := sellerService.GetSeller(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reseller, c)
}

// GetSellerList 分页获取卖家信息列表
// @Tags Seller
// @Summary 分页获取卖家信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.SellerSearch true "分页获取卖家信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /seller/getSellerList [get]
func (sellerApi *SellerApi) GetSellerList(c *gin.Context) {
	var pageInfo bagiqueReq.SellerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := sellerService.GetSellerInfoList(pageInfo)
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

// GetSellerPublic 不需要鉴权的卖家信息接口
// @Tags Seller
// @Summary 不需要鉴权的卖家信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.SellerSearch true "分页获取卖家信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /seller/getSellerPublic [get]
func (sellerApi *SellerApi) GetSellerPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	sellerService.GetSellerPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的卖家信息接口信息",
	}, "获取成功", c)
}

// RestoreSeller 恢复单条seller的数据，也就是将deleted_at设置为null
// @Tags Seller
// @Summary 恢复单条seller的数据，也就是将deleted_at设置为null
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.SellerSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /seller/restoreSeller [PUT]
func (sellerApi *SellerApi) RestoreSeller(c *gin.Context) {
	ID := c.Query("ID")
	err := sellerService.RestoreSeller(ID)
	if err != nil {
		global.GVA_LOG.Error("恢复失败!", zap.Error(err))
		response.FailWithMessage("恢复失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("恢复成功", c)
}
