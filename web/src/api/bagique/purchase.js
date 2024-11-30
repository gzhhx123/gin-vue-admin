import service from '@/utils/request'
// @Tags Purchase
// @Summary 创建采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Purchase true "创建采购信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /purchase/createPurchase [post]
export const createPurchase = (data) => {
  return service({
    url: '/purchase/createPurchase',
    method: 'post',
    data
  })
}

// @Tags Purchase
// @Summary 删除采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Purchase true "删除采购信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchase/deletePurchase [delete]
export const deletePurchase = (params) => {
  return service({
    url: '/purchase/deletePurchase',
    method: 'delete',
    params
  })
}

// @Tags Purchase
// @Summary 批量删除采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除采购信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /purchase/deletePurchase [delete]
export const deletePurchaseByIds = (params) => {
  return service({
    url: '/purchase/deletePurchaseByIds',
    method: 'delete',
    params
  })
}

// @Tags Purchase
// @Summary 更新采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Purchase true "更新采购信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /purchase/updatePurchase [put]
export const updatePurchase = (data) => {
  return service({
    url: '/purchase/updatePurchase',
    method: 'put',
    data
  })
}

// @Tags Purchase
// @Summary 用id查询采购信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Purchase true "用id查询采购信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /purchase/findPurchase [get]
export const findPurchase = (params) => {
  return service({
    url: '/purchase/findPurchase',
    method: 'get',
    params
  })
}

// @Tags Purchase
// @Summary 分页获取采购信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取采购信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /purchase/getPurchaseList [get]
export const getPurchaseList = (params) => {
  return service({
    url: '/purchase/getPurchaseList',
    method: 'get',
    params
  })
}

// @Tags Purchase
// @Summary 不需要鉴权的采购信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.PurchaseSearch true "分页获取采购信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /purchase/getPurchasePublic [get]
export const getPurchasePublic = () => {
  return service({
    url: '/purchase/getPurchasePublic',
    method: 'get',
  })
}
// RefuseEvaluate 根据ID驳回估价
// @Tags Purchase
// @Summary 根据ID驳回估价
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /purchase/refuseEvaluate [PUT]
export const refuseEvaluate = (params) => {
  return service({
    url: '/purchase/refuseEvaluate',
    method: 'PUT',
    params
  })
}
