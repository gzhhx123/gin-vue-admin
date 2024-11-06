// 自动生成模板Seller
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 卖家信息 结构体  Seller
type Seller struct {
	global.GVA_MODEL
	SellerName         *string `json:"sellerName" form:"sellerName" gorm:"column:seller_name;comment:卖家名称;" binding:"required"`         //卖家名称
	SellerPlatformCode *string `json:"sellerPlatformCode" form:"sellerPlatformCode" gorm:"column:seller_platform_code;comment:卖家平台id;"` //卖家平台id
	Remark             *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                           //备注
}

// TableName 卖家信息 Seller自定义表名 bagique_sellers
func (Seller) TableName() string {
	return "bagique_sellers"
}
