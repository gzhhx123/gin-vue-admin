package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SellerSearch struct {
	StartCreatedAt     *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt       *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	SellerName         *string    `json:"sellerName" form:"sellerName" `
	SellerPlatformCode *string    `json:"sellerPlatformCode" form:"sellerPlatformCode" `
	IsRemove           *bool      `json:"isRemove" form:"isRemove" `
	Remark             *string    `json:"remark" form:"remark" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
