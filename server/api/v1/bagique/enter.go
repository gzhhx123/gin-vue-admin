package bagique

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ BrandApi }

var brandService = service.ServiceGroupApp.BagiqueServiceGroup.BrandService
