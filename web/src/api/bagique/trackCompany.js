import service from '@/utils/request'
// @Tags TrackCompany
// @Summary 创建物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrackCompany true "创建物流公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /trackCompany/createTrackCompany [post]
export const createTrackCompany = (data) => {
  return service({
    url: '/trackCompany/createTrackCompany',
    method: 'post',
    data
  })
}

// @Tags TrackCompany
// @Summary 删除物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrackCompany true "删除物流公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trackCompany/deleteTrackCompany [delete]
export const deleteTrackCompany = (params) => {
  return service({
    url: '/trackCompany/deleteTrackCompany',
    method: 'delete',
    params
  })
}

// @Tags TrackCompany
// @Summary 批量删除物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除物流公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /trackCompany/deleteTrackCompany [delete]
export const deleteTrackCompanyByIds = (params) => {
  return service({
    url: '/trackCompany/deleteTrackCompanyByIds',
    method: 'delete',
    params
  })
}

// @Tags TrackCompany
// @Summary 更新物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.TrackCompany true "更新物流公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /trackCompany/updateTrackCompany [put]
export const updateTrackCompany = (data) => {
  return service({
    url: '/trackCompany/updateTrackCompany',
    method: 'put',
    data
  })
}

// @Tags TrackCompany
// @Summary 用id查询物流公司
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.TrackCompany true "用id查询物流公司"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /trackCompany/findTrackCompany [get]
export const findTrackCompany = (params) => {
  return service({
    url: '/trackCompany/findTrackCompany',
    method: 'get',
    params
  })
}

// @Tags TrackCompany
// @Summary 分页获取物流公司列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取物流公司列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /trackCompany/getTrackCompanyList [get]
export const getTrackCompanyList = (params) => {
  return service({
    url: '/trackCompany/getTrackCompanyList',
    method: 'get',
    params
  })
}

// @Tags TrackCompany
// @Summary 不需要鉴权的物流公司接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.TrackCompanySearch true "分页获取物流公司列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /trackCompany/getTrackCompanyPublic [get]
export const getTrackCompanyPublic = () => {
  return service({
    url: '/trackCompany/getTrackCompanyPublic',
    method: 'get',
  })
}
// RestoreTrackCompany 根据ID恢复物流公司
// @Tags TrackCompany
// @Summary 根据ID恢复物流公司
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /trackCompany/restoreTrackCompany [PUT]
export const restoreTrackCompany = (params) => {
  return service({
    url: '/trackCompany/restoreTrackCompany',
    method: 'PUT',
    params
  })
}
