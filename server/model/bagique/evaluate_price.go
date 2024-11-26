package bagique

import "github.com/flipped-aurora/gin-vue-admin/server/global"

// 公司估价 结构体  EvaluatePrice
type EvaluatePrice struct {
	global.GVA_MODEL
	CompanyId  *int     `json:"companyId" form:"companyId" gorm:"column:company_id;comment:公司id;" binding:"required"`      //估价公司
	EvaluateId *int     `json:"evaluateId" form:"evaluateId" gorm:"column:evaluate_id;comment:估价记录id;" binding:"required"` //估价记录id
	Price      *float64 `json:"price" form:"price" gorm:"column:price;comment:估价;default:0;"`                              //估价
	Fee        *float64 `json:"fee" form:"fee" gorm:"column:fee;comment:估价成本;default:0;"`                                  //估价成本
	Remark     *string  `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                     //备注
}

// TableName 公司估价 EvaluatePrice自定义表名 bagique_evaluate_prices
func (EvaluatePrice) TableName() string {
	return "bagique_evaluate_prices"
}
