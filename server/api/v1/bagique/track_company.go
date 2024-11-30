package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type TrackCompanyApi struct{}

// CreateTrackCompany 创建物流公司
// @Tags TrackCompany
// @Summary 创建物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.TrackCompany true "创建物流公司"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /trackCompany/createTrackCompany [post]
func (trackCompanyApi *TrackCompanyApi) CreateTrackCompany(c *gin.Context) {
	var trackCompany bagique.TrackCompany
	err := c.ShouldBindJSON(&trackCompany)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = trackCompanyService.CreateTrackCompany(&trackCompany)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteTrackCompany 删除物流公司
// @Tags TrackCompany
// @Summary 删除物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.TrackCompany true "删除物流公司"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /trackCompany/deleteTrackCompany [delete]
func (trackCompanyApi *TrackCompanyApi) DeleteTrackCompany(c *gin.Context) {
	ID := c.Query("ID")
	err := trackCompanyService.DeleteTrackCompany(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteTrackCompanyByIds 批量删除物流公司
// @Tags TrackCompany
// @Summary 批量删除物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /trackCompany/deleteTrackCompanyByIds [delete]
func (trackCompanyApi *TrackCompanyApi) DeleteTrackCompanyByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := trackCompanyService.DeleteTrackCompanyByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateTrackCompany 更新物流公司
// @Tags TrackCompany
// @Summary 更新物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.TrackCompany true "更新物流公司"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /trackCompany/updateTrackCompany [put]
func (trackCompanyApi *TrackCompanyApi) UpdateTrackCompany(c *gin.Context) {
	var trackCompany bagique.TrackCompany
	err := c.ShouldBindJSON(&trackCompany)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = trackCompanyService.UpdateTrackCompany(trackCompany)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindTrackCompany 用id查询物流公司
// @Tags TrackCompany
// @Summary 用id查询物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.TrackCompany true "用id查询物流公司"
// @Success 200 {object} response.Response{data=bagique.TrackCompany,msg=string} "查询成功"
// @Router /trackCompany/findTrackCompany [get]
func (trackCompanyApi *TrackCompanyApi) FindTrackCompany(c *gin.Context) {
	ID := c.Query("ID")
	retrackCompany, err := trackCompanyService.GetTrackCompany(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(retrackCompany, c)
}

// GetTrackCompanyList 分页获取物流公司列表
// @Tags TrackCompany
// @Summary 分页获取物流公司列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.TrackCompanySearch true "分页获取物流公司列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /trackCompany/getTrackCompanyList [get]
func (trackCompanyApi *TrackCompanyApi) GetTrackCompanyList(c *gin.Context) {
	var pageInfo bagiqueReq.TrackCompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := trackCompanyService.GetTrackCompanyInfoList(pageInfo)
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

// GetTrackCompanyPublic 不需要鉴权的物流公司接口
// @Tags TrackCompany
// @Summary 不需要鉴权的物流公司接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.TrackCompanySearch true "分页获取物流公司列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /trackCompany/getTrackCompanyPublic [get]
func (trackCompanyApi *TrackCompanyApi) GetTrackCompanyPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	trackCompanyService.GetTrackCompanyPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的物流公司接口信息",
	}, "获取成功", c)
}
