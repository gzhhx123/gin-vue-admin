package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type EvaluateSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ProductId      *int       `json:"productId" form:"productId" `
	SellerId       *int       `json:"sellerId" form:"sellerId" `
	Status         *string    `json:"status" form:"status" `
	IsRemove       *bool      `json:"isRemove" form:"isRemove" `
	Remark         *string    `json:"remark" form:"remark" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
