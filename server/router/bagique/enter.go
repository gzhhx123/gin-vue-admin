package bagique

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ BrandRouter }

var brandApi = api.ApiGroupApp.BagiqueApiGroup.BrandApi
