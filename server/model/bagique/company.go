// 自动生成模板Company
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 公司信息 结构体  Company
type Company struct {
	global.GVA_MODEL
	CompanyName      *string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;" binding:"required"` //公司名称
	CompanyShortName *string `json:"companyShortName" form:"companyShortName" gorm:"column:company_short_name;comment:公司简称;"`    //公司简称
	CompanyLogo      string  `json:"companyLogo" form:"companyLogo" gorm:"column:company_logo;comment:公司logo;"`                  //公司logo
	Remark           *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                      //备注
}

// TableName 公司信息 Company自定义表名 bagique_companies
func (Company) TableName() string {
	return "bagique_companies"
}
