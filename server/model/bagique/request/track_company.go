package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type TrackCompanySearch struct {
	StartCreatedAt   *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt     *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	CompanyName      *string    `json:"companyName" form:"companyName" `
	CompanyShortName *string    `json:"companyShortName" form:"companyShortName" `
	CompanyUrl       *string    `json:"companyUrl" form:"companyUrl" `
	Remark           *string    `json:"remark" form:"remark" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
