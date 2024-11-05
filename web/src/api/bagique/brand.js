import service from '@/utils/request'
// @Tags Brand
// @Summary 创建品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "创建品牌信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /brand/createBrand [post]
export const createBrand = (data) => {
  return service({
    url: '/brand/createBrand',
    method: 'post',
    data
  })
}

// @Tags Brand
// @Summary 删除品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "删除品牌信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brand/deleteBrand [delete]
export const deleteBrand = (params) => {
  return service({
    url: '/brand/deleteBrand',
    method: 'delete',
    params
  })
}

// @Tags Brand
// @Summary 批量删除品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除品牌信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /brand/deleteBrand [delete]
export const deleteBrandByIds = (params) => {
  return service({
    url: '/brand/deleteBrandByIds',
    method: 'delete',
    params
  })
}

// @Tags Brand
// @Summary 更新品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Brand true "更新品牌信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /brand/updateBrand [put]
export const updateBrand = (data) => {
  return service({
    url: '/brand/updateBrand',
    method: 'put',
    data
  })
}

// @Tags Brand
// @Summary 用id查询品牌信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Brand true "用id查询品牌信息"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /brand/findBrand [get]
export const findBrand = (params) => {
  return service({
    url: '/brand/findBrand',
    method: 'get',
    params
  })
}

// @Tags Brand
// @Summary 分页获取品牌信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取品牌信息列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /brand/getBrandList [get]
export const getBrandList = (params) => {
  return service({
    url: '/brand/getBrandList',
    method: 'get',
    params
  })
}

// @Tags Brand
// @Summary 不需要鉴权的品牌信息接口
// @accept application/json
// @Produce application/json
// @Param data query bagiqueReq.BrandSearch true "分页获取品牌信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /brand/getBrandPublic [get]
export const getBrandPublic = () => {
  return service({
    url: '/brand/getBrandPublic',
    method: 'get',
  })
}
