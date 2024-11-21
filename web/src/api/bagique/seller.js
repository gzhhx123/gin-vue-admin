import service from '@/utils/request'
// @Tags Seller
// @Summary 创建卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seller true "创建卖家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /seller/createSeller [post]
export const createSeller = (data) => {
  return service({
    url: '/seller/createSeller',
    method: 'post',
    data
  })
}

// @Tags Seller
// @Summary 删除卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seller true "删除卖家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /seller/deleteSeller [delete]
export const deleteSeller = (params) => {
  return service({
    url: '/seller/deleteSeller',
    method: 'delete',
    params
  })
}

// @Tags Seller
// @Summary 批量删除卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除卖家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /seller/deleteSeller [delete]
export const deleteSellerByIds = (params) => {
  return service({
    url: '/seller/deleteSellerByIds',
    method: 'delete',
    params
  })
}

// @Tags Seller
// @Summary 更新卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Seller true "更新卖家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /seller/updateSeller [put]
export const updateSeller = (data) => {
  return service({
    url: '/seller/updateSeller',
    method: 'put',
    data
  })
}

// @Tags Seller
// @Summary 用id查询卖家信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Seller true "用id查询卖家信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /seller/findSeller [get]
export const findSeller = (params) => {
  return service({
    url: '/seller/findSeller',
    method: 'get',
    params
  })
}

// @Tags Seller
// @Summary 分页获取卖家信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取卖家信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /seller/getSellerList [get]
export const getSellerList = (params) => {
  return service({
    url: '/seller/getSellerList',
    method: 'get',
    params
  })
}

// @Tags Seller
// @Summary 不需要鉴权的卖家信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.SellerSearch true "分页获取卖家信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /seller/getSellerPublic [get]
export const getSellerPublic = () => {
  return service({
    url: '/seller/getSellerPublic',
    method: 'get',
  })
}
// RestoreSeller 恢复单条seller的数据，也就是将deleted_at设置为null
// @Tags Seller
// @Summary 恢复单条seller的数据，也就是将deleted_at设置为null
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "成功"
// @Router /seller/restoreSeller [PUT]
export const restoreSeller = (params) => {
  return service({
    url: '/seller/restoreSeller',
    method: 'PUT',
    params
  })
}
