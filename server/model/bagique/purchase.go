// 自动生成模板Purchase
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 采购信息 结构体  Purchase
type Purchase struct {
	global.GVA_MODEL
	EvaluateId      *int           `json:"evaluateId" form:"evaluateId" gorm:"column:evaluate_id;comment:估价公司id;" binding:"required"`                   //估价公司
	EvaluatePriceId *int           `json:"evaluatePriceId" form:"evaluatePriceId" gorm:"column:evaluate_price_id;comment:估价公司估价id;" binding:"required"` //估价公司估价
	Status          *string        `json:"status" form:"status" gorm:"column:status;comment:采购状态;"`                                                     //采购状态
	Remark          *string        `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                                       //备注
	Evaluate        *Evaluate      `json:"evaluate" gorm:"foreignKey:EvaluateId;references:ID;comment:估价公司"`                                            //估价公司
	EvaluatePrice   *EvaluatePrice `json:"evaluatePrice" gorm:"foreignKey:EvaluatePriceId;references:ID;comment:估价公司估价"`                                //估价公司估价
}

// TableName 采购信息 Purchase自定义表名 bagique_purchases
func (Purchase) TableName() string {
	return "bagique_purchases"
}
