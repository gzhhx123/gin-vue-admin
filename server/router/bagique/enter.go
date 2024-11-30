package bagique

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BrandRouter
	ProductRouter
	CompanyRouter
	SellerRouter
	EvaluateRouter
	CommonRouter
	PurchaseRouter
	TrackCompanyRouter
}

var (
	brandApi        = api.ApiGroupApp.BagiqueApiGroup.BrandApi
	productApi      = api.ApiGroupApp.BagiqueApiGroup.ProductApi
	companyApi      = api.ApiGroupApp.BagiqueApiGroup.CompanyApi
	sellerApi       = api.ApiGroupApp.BagiqueApiGroup.SellerApi
	evaluateApi     = api.ApiGroupApp.BagiqueApiGroup.EvaluateApi
	commonApi       = api.ApiGroupApp.BagiqueApiGroup.CommonApi
	purchaseApi     = api.ApiGroupApp.BagiqueApiGroup.PurchaseApi
	trackCompanyApi = api.ApiGroupApp.BagiqueApiGroup.TrackCompanyApi
)
