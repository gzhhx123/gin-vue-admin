// 自动生成模板Brand
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 品牌信息 结构体  Brand
type Brand struct {
	global.GVA_MODEL
	BrandName      *string `json:"brandName" form:"brandName" gorm:"column:brand_name;comment:品牌名称;" binding:"required"` //品牌名称
	BrandShortName *string `json:"brandShortName" form:"brandShortName" gorm:"column:brand_short_name;comment:品牌简称;"`    //品牌简称
	BrandLogo      string  `json:"brandLogo" form:"brandLogo" gorm:"column:brand_logo;comment:品牌logo;"`                  //品牌logo
	Remark         *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                //备注
}

// TableName 品牌信息 Brand自定义表名 bagique_brands
func (Brand) TableName() string {
	return "bagique_brands"
}
