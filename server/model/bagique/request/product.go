package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProductSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	BrandId        *int       `json:"brandId" form:"brandId" `
	ProductName    *string    `json:"productName" form:"productName" `
	ProductSku     *string    `json:"productSku" form:"productSku" `
	Remark         *string    `json:"remark" form:"remark" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
