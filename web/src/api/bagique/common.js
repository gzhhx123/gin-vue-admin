import service from '@/utils/request'

// @Tags Common
// @Summary 获取美元汇率
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query {object} true "获取美元汇率"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /common/getRate [get]
export const getRate = (params) => {
  return service({
    url: '/common/getRate',
    method: 'get',
    params
  })
}