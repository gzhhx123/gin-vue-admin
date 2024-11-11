package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EvaluatePriceSearch struct {
	StartCreatedAt  *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt    *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CompanyId       *int       `json:"companyId" form:"companyId" `
	StartEvaluateId *int       `json:"startEvaluateId" form:"startEvaluateId"`
	EndEvaluateId   *int       `json:"endEvaluateId" form:"endEvaluateId"`
	StartPrice      *float64   `json:"startPrice" form:"startPrice"`
	EndPrice        *float64   `json:"endPrice" form:"endPrice"`
	StartFee        *float64   `json:"startFee" form:"startFee"`
	EndFee          *float64   `json:"endFee" form:"endFee"`
	Remark          *string    `json:"remark" form:"remark" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
