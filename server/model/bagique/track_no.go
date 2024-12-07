// 自动生成模板TrackNo
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 物流单号 结构体  TrackNo
type TrackNo struct {
	global.GVA_MODEL
	OrderSn        *string `json:"orderSn" form:"orderSn" gorm:"column:order_sn;comment:物流单号;" binding:"required"`      //物流单号
	TrackCompanyId *int    `json:"trackCompanyId" form:"trackCompanyId" gorm:"column:track_company_id;comment:转运公司id;"` //转运公司id
	PurchaseId     *int    `json:"purchaseId" form:"purchaseId" gorm:"column:purchase_id;comment:采购id;"`                //采购id
	Type           *string `json:"type" form:"type" gorm:"column:type;comment:物流类型;"`                                   //物流类型
	Status         *string `json:"status" form:"status" gorm:"column:status;comment:物流状态;"`                             //物流状态
	Sort           *int    `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`                                     //排序
	Remark         *string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                               //备注
}

// TableName 物流单号 TrackNo自定义表名 bagique_track_nos
func (TrackNo) TableName() string {
	return "bagique_track_nos"
}
