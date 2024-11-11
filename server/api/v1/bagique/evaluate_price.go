package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EvaluatePriceApi struct{}

// CreateEvaluatePrice 创建公司估价
// @Tags EvaluatePrice
// @Summary 创建公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.EvaluatePrice true "创建公司估价"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /evaluatePrice/createEvaluatePrice [post]
func (evaluatePriceApi *EvaluatePriceApi) CreateEvaluatePrice(c *gin.Context) {
	var evaluatePrice bagique.EvaluatePrice
	err := c.ShouldBindJSON(&evaluatePrice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = evaluatePriceService.CreateEvaluatePrice(&evaluatePrice)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEvaluatePrice 删除公司估价
// @Tags EvaluatePrice
// @Summary 删除公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.EvaluatePrice true "删除公司估价"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /evaluatePrice/deleteEvaluatePrice [delete]
func (evaluatePriceApi *EvaluatePriceApi) DeleteEvaluatePrice(c *gin.Context) {
	ID := c.Query("ID")
	err := evaluatePriceService.DeleteEvaluatePrice(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEvaluatePriceByIds 批量删除公司估价
// @Tags EvaluatePrice
// @Summary 批量删除公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /evaluatePrice/deleteEvaluatePriceByIds [delete]
func (evaluatePriceApi *EvaluatePriceApi) DeleteEvaluatePriceByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := evaluatePriceService.DeleteEvaluatePriceByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEvaluatePrice 更新公司估价
// @Tags EvaluatePrice
// @Summary 更新公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.EvaluatePrice true "更新公司估价"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /evaluatePrice/updateEvaluatePrice [put]
func (evaluatePriceApi *EvaluatePriceApi) UpdateEvaluatePrice(c *gin.Context) {
	var evaluatePrice bagique.EvaluatePrice
	err := c.ShouldBindJSON(&evaluatePrice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = evaluatePriceService.UpdateEvaluatePrice(evaluatePrice)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEvaluatePrice 用id查询公司估价
// @Tags EvaluatePrice
// @Summary 用id查询公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.EvaluatePrice true "用id查询公司估价"
// @Success 200 {object} response.Response{data=bagique.EvaluatePrice,msg=string} "查询成功"
// @Router /evaluatePrice/findEvaluatePrice [get]
func (evaluatePriceApi *EvaluatePriceApi) FindEvaluatePrice(c *gin.Context) {
	ID := c.Query("ID")
	reevaluatePrice, err := evaluatePriceService.GetEvaluatePrice(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reevaluatePrice, c)
}

// GetEvaluatePriceList 分页获取公司估价列表
// @Tags EvaluatePrice
// @Summary 分页获取公司估价列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluatePriceSearch true "分页获取公司估价列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /evaluatePrice/getEvaluatePriceList [get]
func (evaluatePriceApi *EvaluatePriceApi) GetEvaluatePriceList(c *gin.Context) {
	var pageInfo bagiqueReq.EvaluatePriceSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := evaluatePriceService.GetEvaluatePriceInfoList(pageInfo)
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

// GetEvaluatePriceDataSource 获取EvaluatePrice的数据源
// @Tags EvaluatePrice
// @Summary 获取EvaluatePrice的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /evaluatePrice/getEvaluatePriceDataSource [get]
func (evaluatePriceApi *EvaluatePriceApi) GetEvaluatePriceDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := evaluatePriceService.GetEvaluatePriceDataSource()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetEvaluatePricePublic 不需要鉴权的公司估价接口
// @Tags EvaluatePrice
// @Summary 不需要鉴权的公司估价接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluatePriceSearch true "分页获取公司估价列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /evaluatePrice/getEvaluatePricePublic [get]
func (evaluatePriceApi *EvaluatePriceApi) GetEvaluatePricePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	evaluatePriceService.GetEvaluatePricePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的公司估价接口信息",
	}, "获取成功", c)
}
