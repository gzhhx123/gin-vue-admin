package bagique

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BrandApi
	ProductApi
}

var (
	brandService   = service.ServiceGroupApp.BagiqueServiceGroup.BrandService
	productService = service.ServiceGroupApp.BagiqueServiceGroup.ProductService
)
