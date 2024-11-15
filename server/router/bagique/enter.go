package bagique

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BrandRouter
	ProductRouter
	CompanyRouter
	SellerRouter
	EvaluateRouter
	EvaluatePriceRouter
	CommonRouter
}

var (
	brandApi         = api.ApiGroupApp.BagiqueApiGroup.BrandApi
	productApi       = api.ApiGroupApp.BagiqueApiGroup.ProductApi
	companyApi       = api.ApiGroupApp.BagiqueApiGroup.CompanyApi
	sellerApi        = api.ApiGroupApp.BagiqueApiGroup.SellerApi
	evaluateApi      = api.ApiGroupApp.BagiqueApiGroup.EvaluateApi
	evaluatePriceApi = api.ApiGroupApp.BagiqueApiGroup.EvaluatePriceApi
	commonApi        = api.ApiGroupApp.BagiqueApiGroup.CommonApi
)
