// 自动生成模板Evaluate
package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/datatypes"
)

// 估价信息 结构体  Evaluate
type Evaluate struct {
	global.GVA_MODEL
	ProductId      *int            `json:"productId" form:"productId" gorm:"column:product_id;comment:产品id;" binding:"required"`                                    //产品
	SellerId       *int            `json:"sellerId" form:"sellerId" gorm:"column:seller_id;comment:卖家id;" binding:"required"`                                       //卖家
	EvaluatePics   datatypes.JSON  `json:"evaluatePics" form:"evaluatePics" gorm:"column:evaluate_pics;comment:细节图;" binding:"required" swaggertype:"array,object"` //细节图
	Status         *string         `json:"status" form:"status" gorm:"column:status;comment:估价状态;" binding:"required"`                                              //估价状态
	Remark         *string         `json:"remark" form:"remark" gorm:"column:remark;comment:备注;"`                                                                   //备注
	EvaluatePrices []EvaluatePrice `json:"evaluatePrices" gorm:"foreignKey:EvaluateId;references:ID;"`                                                              //公司估价
}

// TableName 估价信息 Evaluate自定义表名 bagique_evaluates
func (Evaluate) TableName() string {
	return "bagique_evaluates"
}
