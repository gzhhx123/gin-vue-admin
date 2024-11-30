import service from '@/utils/request'
// @Tags Evaluate
// @Summary 创建估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluate true "创建估价信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /evaluate/createEvaluate [post]
export const createEvaluate = (data) => {
  return service({
    url: '/evaluate/createEvaluate',
    method: 'post',
    data
  })
}

// @Tags Evaluate
// @Summary 删除估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluate true "删除估价信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /evaluate/deleteEvaluate [delete]
export const deleteEvaluate = (params) => {
  return service({
    url: '/evaluate/deleteEvaluate',
    method: 'delete',
    params
  })
}

// @Tags Evaluate
// @Summary 批量删除估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除估价信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /evaluate/deleteEvaluate [delete]
export const deleteEvaluateByIds = (params) => {
  return service({
    url: '/evaluate/deleteEvaluateByIds',
    method: 'delete',
    params
  })
}

// @Tags Evaluate
// @Summary 更新估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Evaluate true "更新估价信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /evaluate/updateEvaluate [put]
export const updateEvaluate = (data) => {
  return service({
    url: '/evaluate/updateEvaluate',
    method: 'put',
    data
  })
}

// @Tags Evaluate
// @Summary 用id查询估价信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Evaluate true "用id查询估价信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /evaluate/findEvaluate [get]
export const findEvaluate = (params) => {
  return service({
    url: '/evaluate/findEvaluate',
    method: 'get',
    params
  })
}

// @Tags Evaluate
// @Summary 分页获取估价信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取估价信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /evaluate/getEvaluateList [get]
export const getEvaluateList = (params) => {
  return service({
    url: '/evaluate/getEvaluateList',
    method: 'get',
    params
  })
}
// @Tags Evaluate
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /evaluate/findEvaluateDataSource [get]
export const getEvaluateDataSource = () => {
  return service({
    url: '/evaluate/getEvaluateDataSource',
    method: 'get',
  })
}

// @Tags Evaluate
// @Summary 不需要鉴权的估价信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluateSearch true "分页获取估价信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /evaluate/getEvaluatePublic [get]
export const getEvaluatePublic = () => {
  return service({
    url: '/evaluate/getEvaluatePublic',
    method: 'get',
  })
}
// RestoreEvaluate 恢复单条evaluate的数据，也就是将deleted_at设置为null
// @Tags Evaluate
// @Summary 恢复单条evaluate的数据，也就是将deleted_at设置为null
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /evaluate/restoreEvaluate [PUT]
export const restoreEvaluate = (params) => {
  return service({
    url: '/evaluate/restoreEvaluate',
    method: 'PUT',
    params
  })
}
// FinishEvaluate 根据ID完成估价
// @Tags Evaluate
// @Summary 根据ID完成估价
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /evaluate/finishEvaluate [PUT]
export const finishEvaluate = (params) => {
  return service({
    url: '/evaluate/finishEvaluate',
    method: 'PUT',
    params
  })
}
// CancelEvaluate 根据ID取消估价
// @Tags Evaluate
// @Summary 根据ID取消估价
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /evaluate/cancelEvaluate [PUT]
export const cancelEvaluate = (params) => {
  return service({
    url: '/evaluate/cancelEvaluate',
    method: 'PUT',
    params
  })
}
