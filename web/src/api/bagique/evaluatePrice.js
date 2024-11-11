import service from '@/utils/request'
// @Tags EvaluatePrice
// @Summary 创建公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluatePrice true "创建公司估价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /evaluatePrice/createEvaluatePrice [post]
export const createEvaluatePrice = (data) => {
  return service({
    url: '/evaluatePrice/createEvaluatePrice',
    method: 'post',
    data
  })
}

// @Tags EvaluatePrice
// @Summary 删除公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluatePrice true "删除公司估价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /evaluatePrice/deleteEvaluatePrice [delete]
export const deleteEvaluatePrice = (params) => {
  return service({
    url: '/evaluatePrice/deleteEvaluatePrice',
    method: 'delete',
    params
  })
}

// @Tags EvaluatePrice
// @Summary 批量删除公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除公司估价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /evaluatePrice/deleteEvaluatePrice [delete]
export const deleteEvaluatePriceByIds = (params) => {
  return service({
    url: '/evaluatePrice/deleteEvaluatePriceByIds',
    method: 'delete',
    params
  })
}

// @Tags EvaluatePrice
// @Summary 更新公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.EvaluatePrice true "更新公司估价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /evaluatePrice/updateEvaluatePrice [put]
export const updateEvaluatePrice = (data) => {
  return service({
    url: '/evaluatePrice/updateEvaluatePrice',
    method: 'put',
    data
  })
}

// @Tags EvaluatePrice
// @Summary 用id查询公司估价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.EvaluatePrice true "用id查询公司估价"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /evaluatePrice/findEvaluatePrice [get]
export const findEvaluatePrice = (params) => {
  return service({
    url: '/evaluatePrice/findEvaluatePrice',
    method: 'get',
    params
  })
}

// @Tags EvaluatePrice
// @Summary 分页获取公司估价列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取公司估价列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /evaluatePrice/getEvaluatePriceList [get]
export const getEvaluatePriceList = (params) => {
  return service({
    url: '/evaluatePrice/getEvaluatePriceList',
    method: 'get',
    params
  })
}
// @Tags EvaluatePrice
// @Summary 获取数据源
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /evaluatePrice/findEvaluatePriceDataSource [get]
export const getEvaluatePriceDataSource = () => {
  return service({
    url: '/evaluatePrice/getEvaluatePriceDataSource',
    method: 'get',
  })
}

// @Tags EvaluatePrice
// @Summary 不需要鉴权的公司估价接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.EvaluatePriceSearch true "分页获取公司估价列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /evaluatePrice/getEvaluatePricePublic [get]
export const getEvaluatePricePublic = () => {
  return service({
    url: '/evaluatePrice/getEvaluatePricePublic',
    method: 'get',
  })
}
