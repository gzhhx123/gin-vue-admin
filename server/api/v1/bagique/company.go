package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyApi struct{}

// CreateCompany 创建估价公司信息
// @Tags Company
// @Summary 创建估价公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Company true "创建估价公司信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /company/createCompany [post]
func (companyApi *CompanyApi) CreateCompany(c *gin.Context) {
	var company bagique.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = companyService.CreateCompany(&company)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteCompany 删除估价公司信息
// @Tags Company
// @Summary 删除估价公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Company true "删除估价公司信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /company/deleteCompany [delete]
func (companyApi *CompanyApi) DeleteCompany(c *gin.Context) {
	ID := c.Query("ID")
	err := companyService.DeleteCompany(ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteCompanyByIds 批量删除估价公司信息
// @Tags Company
// @Summary 批量删除估价公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /company/deleteCompanyByIds [delete]
func (companyApi *CompanyApi) DeleteCompanyByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	err := companyService.DeleteCompanyByIds(IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateCompany 更新估价公司信息
// @Tags Company
// @Summary 更新估价公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bagique.Company true "更新估价公司信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /company/updateCompany [put]
func (companyApi *CompanyApi) UpdateCompany(c *gin.Context) {
	var company bagique.Company
	err := c.ShouldBindJSON(&company)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = companyService.UpdateCompany(company)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindCompany 用id查询估价公司信息
// @Tags Company
// @Summary 用id查询估价公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagique.Company true "用id查询估价公司信息"
// @Success 200 {object} response.Response{data=bagique.Company,msg=string} "查询成功"
// @Router /company/findCompany [get]
func (companyApi *CompanyApi) FindCompany(c *gin.Context) {
	ID := c.Query("ID")
	recompany, err := companyService.GetCompany(ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(recompany, c)
}

// GetCompanyList 分页获取估价公司信息列表
// @Tags Company
// @Summary 分页获取估价公司信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.CompanySearch true "分页获取估价公司信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /company/getCompanyList [get]
func (companyApi *CompanyApi) GetCompanyList(c *gin.Context) {
	var pageInfo bagiqueReq.CompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := companyService.GetCompanyInfoList(pageInfo)
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

// GetCompanyPublic 不需要鉴权的估价公司信息接口
// @Tags Company
// @Summary 不需要鉴权的估价公司信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.CompanySearch true "分页获取估价公司信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /company/getCompanyPublic [get]
func (companyApi *CompanyApi) GetCompanyPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	companyService.GetCompanyPublic()
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的估价公司信息接口信息",
	}, "获取成功", c)
}
