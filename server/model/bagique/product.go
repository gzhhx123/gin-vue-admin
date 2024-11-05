// 自动生成模板Product
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 产品信息 结构体  Product
type Product struct {
	global.GVA_MODEL
	BrandId       *int           `json:"brandId" form:"brandId" gorm:"column:brand_id;comment:品牌id;"`                                          //品牌
	ProductName   *string        `json:"productName" form:"productName" gorm:"column:product_name;comment:产品名称;" binding:"required"`           //产品名称
	ProductSku    *string        `json:"productSku" form:"productSku" gorm:"uniqueIndex;column:product_sku;comment:产品sku;" binding:"required"` //产品sku
	ReferPriceMin *float64       `json:"referPriceMin" form:"referPriceMin" gorm:"column:refer_price_min;comment:参考价（Min）;"`                   //参考价（Min）
	ReferPriceMax *float64       `json:"referPriceMax" form:"referPriceMax" gorm:"column:refer_price_max;comment:参考价（Max）;"`                   //参考价（Max）
	ReferPics     datatypes.JSON `json:"referPics" form:"referPics" gorm:"column:refer_pics;comment:参考图片;"swaggertype:"array,object"`          //参考图片
	Remark        *string        `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                                //备注
}

// TableName 产品信息 Product自定义表名 bagique_products
func (Product) TableName() string {
	return "bagique_products"
}
