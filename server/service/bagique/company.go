package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type CompanyService struct{}

// CreateCompany 创建公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) CreateCompany(company *bagique.Company) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(company).Error
	return err
}

// DeleteCompany 删除公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) DeleteCompany(ID string, TYPE string) (err error) {
	if TYPE == "HARD" {
		err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Delete(&bagique.Company{}, "id = ?", ID).Error
	} else {
		err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Company{}, "id = ?", ID).Error
	}
	return err
}

// DeleteCompanyByIds 批量删除公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) DeleteCompanyByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Company{}, "id in ?", IDs).Error
	return err
}

// UpdateCompany 更新公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) UpdateCompany(company bagique.Company) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Company{}).Where("id = ?", company.ID).Updates(&company).Error
	return err
}

// GetCompany 根据ID获取公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) GetCompany(ID string) (company bagique.Company, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&company).Error
	return
}

// GetCompanyInfoList 分页获取公司信息记录
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) GetCompanyInfoList(info bagiqueReq.CompanySearch) (list []bagique.Company, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Company{})
	var companys []bagique.Company
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
	if info.Remark != nil && *info.Remark != "" {
		db = db.Where("remark LIKE ?", "%"+*info.Remark+"%")
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
	orderMap["remark"] = true
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

	err = db.Find(&companys).Error
	return companys, total, err
}
func (companyService *CompanyService) GetCompanyPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// RestoreCompany 恢复单条company的数据，也就是将deleted_at设置为null
// Author [yourname](https://github.com/yourname)
func (companyService *CompanyService) RestoreCompany(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Model(&bagique.Company{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	return err
}
