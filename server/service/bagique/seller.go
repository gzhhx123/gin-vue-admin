package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type SellerService struct{}

// CreateSeller 创建卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) CreateSeller(seller *bagique.Seller) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(seller).Error
	return err
}

// DeleteSeller 删除卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) DeleteSeller(ID string, TYPE string) (err error) {
	if TYPE == "HARD" {
		err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Delete(&bagique.Seller{}, "id = ?", ID).Error
	} else {
		err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Seller{}, "id = ?", ID).Error
	}
	return err
}

// DeleteSellerByIds 批量删除卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) DeleteSellerByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Seller{}, "id in ?", IDs).Error
	return err
}

// UpdateSeller 更新卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) UpdateSeller(seller bagique.Seller) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Seller{}).Where("id = ?", seller.ID).Updates(&seller).Error
	return err
}

// GetSeller 根据ID获取卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) GetSeller(ID string) (seller bagique.Seller, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&seller).Error
	return
}

// GetSellerInfoList 分页获取卖家信息记录
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) GetSellerInfoList(info bagiqueReq.SellerSearch) (list []bagique.Seller, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Seller{})
	var sellers []bagique.Seller
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsRemove != nil && *info.IsRemove == true {
		db = db.Unscoped()
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.SellerName != nil && *info.SellerName != "" {
		db = db.Where("seller_name LIKE ?", "%"+*info.SellerName+"%")
	}
	if info.SellerPlatformCode != nil && *info.SellerPlatformCode != "" {
		db = db.Where("seller_platform_code LIKE ?", "%"+*info.SellerPlatformCode+"%")
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
	orderMap["seller_name"] = true
	orderMap["seller_platform_code"] = true
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

	err = db.Find(&sellers).Error
	return sellers, total, err
}
func (sellerService *SellerService) GetSellerPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// RestoreSeller 恢复单条seller的数据，也就是将deleted_at设置为null
// Author [yourname](https://github.com/yourname)
func (sellerService *SellerService) RestoreSeller(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Model(&bagique.Seller{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	return err
}
