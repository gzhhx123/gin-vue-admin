import service from '@/utils/request'
// @Tags Company
// @Summary 创建公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "创建公司信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /company/createCompany [post]
export const createCompany = (data) => {
  return service({
    url: '/company/createCompany',
    method: 'post',
    data
  })
}

// @Tags Company
// @Summary 删除公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "删除公司信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompany = (params) => {
  return service({
    url: '/company/deleteCompany',
    method: 'delete',
    params
  })
}

// @Tags Company
// @Summary 批量删除公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公司信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /company/deleteCompany [delete]
export const deleteCompanyByIds = (params) => {
  return service({
    url: '/company/deleteCompanyByIds',
    method: 'delete',
    params
  })
}

// @Tags Company
// @Summary 更新公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "更新公司信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /company/updateCompany [put]
export const updateCompany = (data) => {
  return service({
    url: '/company/updateCompany',
    method: 'put',
    data
  })
}

// @Tags Company
// @Summary 用id查询公司信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Company true "用id查询公司信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /company/findCompany [get]
export const findCompany = (params) => {
  return service({
    url: '/company/findCompany',
    method: 'get',
    params
  })
}

// @Tags Company
// @Summary 分页获取公司信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公司信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /company/getCompanyList [get]
export const getCompanyList = (params) => {
  return service({
    url: '/company/getCompanyList',
    method: 'get',
    params
  })
}

// @Tags Company
// @Summary 不需要鉴权的公司信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.CompanySearch true "分页获取公司信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /company/getCompanyPublic [get]
export const getCompanyPublic = () => {
  return service({
    url: '/company/getCompanyPublic',
    method: 'get',
  })
}
// RestoreCompany 恢复单条company的数据，也就是将deleted_at设置为null
// @Tags Company
// @Summary 恢复单条company的数据，也就是将deleted_at设置为null
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /company/restoreCompany [PUT]
export const restoreCompany = (params) => {
  return service({
    url: '/company/restoreCompany',
    method: 'PUT',
    params
  })
}
