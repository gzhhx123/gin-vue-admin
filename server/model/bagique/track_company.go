// 自动生成模板TrackCompany
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 物流公司 结构体  TrackCompany
type TrackCompany struct {
	global.GVA_MODEL
	CompanyName      *string `json:"companyName" form:"companyName" gorm:"column:company_name;comment:公司名称;" binding:"required"` //公司名称
	CompanyShortName *string `json:"companyShortName" form:"companyShortName" gorm:"column:company_short_name;comment:公司简称;"`    //公司简称
	CompanyLogo      string  `json:"companyLogo" form:"companyLogo" gorm:"column:company_logo;comment:公司logo;"`                  //公司logo
	CompanyUrl       *string `json:"companyUrl" form:"companyUrl" gorm:"column:company_url;comment:查询链接;"`                       //查询链接
	Country          *string `json:"country" form:"country" gorm:"column:country;comment:所属国家;" binding:"required"`              //所属国家
	Remark           *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                      //备注
}

// TableName 物流公司 TrackCompany自定义表名 bagique_track_companies
func (TrackCompany) TableName() string {
	return "bagique_track_companies"
}
