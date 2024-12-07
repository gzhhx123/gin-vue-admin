package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type TrackCompanyService struct{}

// CreateTrackCompany 创建物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) CreateTrackCompany(trackCompany *bagique.TrackCompany) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(trackCompany).Error
	return err
}

// DeleteTrackCompany 删除物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) DeleteTrackCompany(ID string, TYPE string) (err error) {
	if TYPE == "HARD" {
		err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Delete(&bagique.TrackCompany{}, "id = ?", ID).Error
	} else {
		err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.TrackCompany{}, "id = ?", ID).Error
	}
	return err
}

// DeleteTrackCompanyByIds 批量删除物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) DeleteTrackCompanyByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.TrackCompany{}, "id in ?", IDs).Error
	return err
}

// UpdateTrackCompany 更新物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) UpdateTrackCompany(trackCompany bagique.TrackCompany) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.TrackCompany{}).Where("id = ?", trackCompany.ID).Updates(&trackCompany).Error
	return err
}

// GetTrackCompany 根据ID获取物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) GetTrackCompany(ID string) (trackCompany bagique.TrackCompany, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&trackCompany).Error
	return
}

// GetTrackCompanyInfoList 分页获取物流公司记录
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) GetTrackCompanyInfoList(info bagiqueReq.TrackCompanySearch) (list []bagique.TrackCompany, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.TrackCompany{})
	var trackCompanys []bagique.TrackCompany
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsRemove != nil && *info.IsRemove == true {
		db = db.Unscoped()
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyName != nil && *info.CompanyName != "" {
		db = db.Where("company_name LIKE ?", "%"+*info.CompanyName+"%")
	}
	if info.CompanyShortName != nil && *info.CompanyShortName != "" {
		db = db.Where("company_short_name LIKE ?", "%"+*info.CompanyShortName+"%")
	}
	if info.CompanyUrl != nil && *info.CompanyUrl != "" {
		db = db.Where("company_url LIKE ?", "%"+*info.CompanyUrl+"%")
	}
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
	}
	if info.Country != nil && *info.Country != "" {
		db = db.Where("country = ?", *info.Country)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["created_at"] = true
	orderMap["company_name"] = true
	orderMap["company_short_name"] = true
	orderMap["company_logo"] = true
	orderMap["company_url"] = true
	orderMap["remark"] = true
	orderMap["country"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&trackCompanys).Error
	return trackCompanys, total, err
}
func (trackCompanyService *TrackCompanyService) GetTrackCompanyPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// RestoreTrackCompany 根据ID恢复物流公司
// Author [yourname](https://github.com/yourname)
func (trackCompanyService *TrackCompanyService) RestoreTrackCompany(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Model(&bagique.TrackCompany{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	return err
}
