package bagique

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BrandApi
	ProductApi
	CompanyApi
	SellerApi
	EvaluateApi
	CommonApi
}

var (
	brandService    = service.ServiceGroupApp.BagiqueServiceGroup.BrandService
	productService  = service.ServiceGroupApp.BagiqueServiceGroup.ProductService
	companyService  = service.ServiceGroupApp.BagiqueServiceGroup.CompanyService
	sellerService   = service.ServiceGroupApp.BagiqueServiceGroup.SellerService
	evaluateService = service.ServiceGroupApp.BagiqueServiceGroup.EvaluateService
	commonService   = service.ServiceGroupApp.BagiqueServiceGroup.CommonService
)
