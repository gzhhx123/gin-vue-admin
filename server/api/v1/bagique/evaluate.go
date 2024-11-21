package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type EvaluateApi struct{}

// CreateEvaluate 创建估价信息
// @Tags Evaluate
// @Summary 创建估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Evaluate true "创建估价信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /evaluate/createEvaluate [post]
func (evaluateApi *EvaluateApi) CreateEvaluate(c *gin.Context) {
	var evaluate bagique.Evaluate
	err := c.ShouldBindJSON(&evaluate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = evaluateService.CreateEvaluate(&evaluate)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteEvaluate 删除估价信息
// @Tags Evaluate
// @Summary 删除估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Evaluate true "删除估价信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /evaluate/deleteEvaluate [delete]
func (evaluateApi *EvaluateApi) DeleteEvaluate(c *gin.Context) {
	ID := c.Query("ID")
	TYPE := c.Query("TYPE")
	err := evaluateService.DeleteEvaluate(ID, TYPE)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteEvaluateByIds 批量删除估价信息
// @Tags Evaluate
// @Summary 批量删除估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /evaluate/deleteEvaluateByIds [delete]
func (evaluateApi *EvaluateApi) DeleteEvaluateByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := evaluateService.DeleteEvaluateByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateEvaluate 更新估价信息
// @Tags Evaluate
// @Summary 更新估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Evaluate true "更新估价信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /evaluate/updateEvaluate [put]
func (evaluateApi *EvaluateApi) UpdateEvaluate(c *gin.Context) {
	var evaluate bagique.Evaluate
	err := c.ShouldBindJSON(&evaluate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = evaluateService.UpdateEvaluate(evaluate)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindEvaluate 用id查询估价信息
// @Tags Evaluate
// @Summary 用id查询估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Evaluate true "用id查询估价信息"
// @Success 200 {object} response.Response{data=bagique.Evaluate,msg=string} "查询成功"
// @Router /evaluate/findEvaluate [get]
func (evaluateApi *EvaluateApi) FindEvaluate(c *gin.Context) {
	ID := c.Query("ID")
	reevaluate, err := evaluateService.GetEvaluate(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reevaluate, c)
}

// GetEvaluateList 分页获取估价信息列表
// @Tags Evaluate
// @Summary 分页获取估价信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluateSearch true "分页获取估价信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /evaluate/getEvaluateList [get]
func (evaluateApi *EvaluateApi) GetEvaluateList(c *gin.Context) {
	var pageInfo bagiqueReq.EvaluateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := evaluateService.GetEvaluateInfoList(pageInfo)
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

// GetEvaluateDataSource 获取Evaluate的数据源
// @Tags Evaluate
// @Summary 获取Evaluate的数据源
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "查询成功"
// @Router /evaluate/getEvaluateDataSource [get]
func (evaluateApi *EvaluateApi) GetEvaluateDataSource(c *gin.Context) {
	// 此接口为获取数据源定义的数据
	dataSource, err := evaluateService.GetEvaluateDataSource()
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(dataSource, c)
}

// GetEvaluatePublic 不需要鉴权的估价信息接口
// @Tags Evaluate
// @Summary 不需要鉴权的估价信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluateSearch true "分页获取估价信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /evaluate/getEvaluatePublic [get]
func (evaluateApi *EvaluateApi) GetEvaluatePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	evaluateService.GetEvaluatePublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的估价信息接口信息",
	}, "获取成功", c)
}

// RestoreEvaluate 恢复单条evaluate的数据，也就是将deleted_at设置为null
// @Tags Evaluate
// @Summary 恢复单条evaluate的数据，也就是将deleted_at设置为null
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluateSearch true "成功"
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /evaluate/restoreEvaluate [PUT]
func (evaluateApi *EvaluateApi) RestoreEvaluate(c *gin.Context) {
	ID := c.Query("ID")
	err := evaluateService.RestoreEvaluate(ID)
	if err != nil {
		global.GVA_LOG.Error("恢复失败!", zap.Error(err))
		response.FailWithMessage("恢复失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("恢复成功", c)
}
