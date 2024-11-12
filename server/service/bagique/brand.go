package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type BrandService struct{}

// CreateBrand 创建品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) CreateBrand(brand *bagique.Brand) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(brand).Error
	return err
}

// DeleteBrand 删除品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) DeleteBrand(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Brand{}, "id = ?", ID).Error
	return err
}

// DeleteBrandByIds 批量删除品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) DeleteBrandByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Brand{}, "id in ?", IDs).Error
	return err
}

// UpdateBrand 更新品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) UpdateBrand(brand bagique.Brand) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Brand{}).Where("id = ?", brand.ID).Updates(&brand).Error
	return err
}

// GetBrand 根据ID获取品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) GetBrand(ID string) (brand bagique.Brand, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&brand).Error
	return
}

// GetBrandInfoList 分页获取品牌信息记录
// Author [yourname](https://github.com/yourname)
func (brandService *BrandService) GetBrandInfoList(info bagiqueReq.BrandSearch) (list []bagique.Brand, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Brand{})
	var brands []bagique.Brand
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BrandName != nil && *info.BrandName != "" {
		db = db.Where("brand_name LIKE ?", "%"+*info.BrandName+"%")
	}
	if info.BrandShortName != nil && *info.BrandShortName != "" {
		db = db.Where("brand_short_name LIKE ?", "%"+*info.BrandShortName+"%")
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
	orderMap["brand_name"] = true
	orderMap["brand_short_name"] = true
	orderMap["brand_logo"] = true
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

	err = db.Find(&brands).Error
	return brands, total, err
}
func (brandService *BrandService) GetBrandPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
