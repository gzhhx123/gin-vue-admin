package bagique

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct {
	BrandRouter
	ProductRouter
	CompanyRouter
}

var (
	brandApi   = api.ApiGroupApp.BagiqueApiGroup.BrandApi
	productApi = api.ApiGroupApp.BagiqueApiGroup.ProductApi
	companyApi = api.ApiGroupApp.BagiqueApiGroup.CompanyApi
)
