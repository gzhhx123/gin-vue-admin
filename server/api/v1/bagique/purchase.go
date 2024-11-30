package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PurchaseApi struct{}

// CreatePurchase 创建采购信息
// @Tags Purchase
// @Summary 创建采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Purchase true "创建采购信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /purchase/createPurchase [post]
func (purchaseApi *PurchaseApi) CreatePurchase(c *gin.Context) {
	var purchase bagique.Purchase
	err := c.ShouldBindJSON(&purchase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = purchaseService.CreatePurchase(&purchase)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeletePurchase 删除采购信息
// @Tags Purchase
// @Summary 删除采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Purchase true "删除采购信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /purchase/deletePurchase [delete]
func (purchaseApi *PurchaseApi) DeletePurchase(c *gin.Context) {
	ID := c.Query("ID")
	err := purchaseService.DeletePurchase(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeletePurchaseByIds 批量删除采购信息
// @Tags Purchase
// @Summary 批量删除采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /purchase/deletePurchaseByIds [delete]
func (purchaseApi *PurchaseApi) DeletePurchaseByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := purchaseService.DeletePurchaseByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdatePurchase 更新采购信息
// @Tags Purchase
// @Summary 更新采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Purchase true "更新采购信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /purchase/updatePurchase [put]
func (purchaseApi *PurchaseApi) UpdatePurchase(c *gin.Context) {
	var purchase bagique.Purchase
	err := c.ShouldBindJSON(&purchase)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = purchaseService.UpdatePurchase(purchase)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindPurchase 用id查询采购信息
// @Tags Purchase
// @Summary 用id查询采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Purchase true "用id查询采购信息"
// @Success 200 {object} response.Response{data=bagique.Purchase,msg=string} "查询成功"
// @Router /purchase/findPurchase [get]
func (purchaseApi *PurchaseApi) FindPurchase(c *gin.Context) {
	ID := c.Query("ID")
	repurchase, err := purchaseService.GetPurchase(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(repurchase, c)
}

// GetPurchaseList 分页获取采购信息列表
// @Tags Purchase
// @Summary 分页获取采购信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.PurchaseSearch true "分页获取采购信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /purchase/getPurchaseList [get]
func (purchaseApi *PurchaseApi) GetPurchaseList(c *gin.Context) {
	var pageInfo bagiqueReq.PurchaseSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := purchaseService.GetPurchaseInfoList(pageInfo)
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

// GetPurchasePublic 不需要鉴权的采购信息接口
// @Tags Purchase
// @Summary 不需要鉴权的采购信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.PurchaseSearch true "分页获取采购信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /purchase/getPurchasePublic [get]
func (purchaseApi *PurchaseApi) GetPurchasePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	purchaseService.GetPurchasePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的采购信息接口信息",
	}, "获取成功", c)
}

// RefuseEvaluate 根据ID驳回估价
// @Tags Purchase
// @Summary 根据ID驳回估价
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.PurchaseSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /purchase/refuseEvaluate [PUT]
func (purchaseApi *PurchaseApi) RefuseEvaluate(c *gin.Context) {
	ID := c.Query("ID")
	err := purchaseService.RefuseEvaluate(ID)
	if err != nil {
		global.GVA_LOG.Error("驳回失败!", zap.Error(err))
		response.FailWithMessage("驳回失败", c)
		return
	}
	response.OkWithData("驳回成功", c)
}
