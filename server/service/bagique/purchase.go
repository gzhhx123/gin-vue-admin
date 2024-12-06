package bagique

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type PurchaseService struct{}

// CreatePurchase 创建采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) CreatePurchase(purchase *bagique.Purchase) (err error) {
	//先判断purchase表是否已经有EvaluateId
	var count int64
	global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Purchase{}).Where("evaluate_id = ?", *purchase.EvaluateId).Count(&count)
	if count > 0 {
		err = errors.New("已经存在该估价公司的采购信息")
		return err
	}
	purchase.Status = new(string)
	*purchase.Status = "ADD"
	err = global.MustGetGlobalDBByDBName("bagique").Create(purchase).Error
	return err
}

// DeletePurchase 删除采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) DeletePurchase(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Purchase{}, "id = ?", ID).Error
	return err
}

// DeletePurchaseByIds 批量删除采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) DeletePurchaseByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Purchase{}, "id in ?", IDs).Error
	return err
}

// UpdatePurchase 更新采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) UpdatePurchase(purchase bagique.Purchase) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Purchase{}).Where("id = ?", purchase.ID).Updates(&purchase).Error
	return err
}

// GetPurchase 根据ID获取采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) GetPurchase(ID string) (purchase bagique.Purchase, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Preload("Evaluate").Preload("EvaluatePrice").Where("id = ?", ID).First(&purchase).Error
	return
}

// GetPurchaseInfoList 分页获取采购信息记录
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) GetPurchaseInfoList(info bagiqueReq.PurchaseSearch) (list []bagique.Purchase, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Preload("Evaluate").Preload("EvaluatePrice").Model(&bagique.Purchase{})
	var purchases []bagique.Purchase
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsRemove != nil && *info.IsRemove == true {
		db = db.Unscoped()
	}
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
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
	orderMap["evaluate_price_id"] = true
	orderMap["status"] = true
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

	err = db.Find(&purchases).Error
	return purchases, total, err
}

func (purchaseService *PurchaseService) GetPurchaseDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	companyId := make([]map[string]any, 0)
	productId := make([]map[string]any, 0)
	global.MustGetGlobalDBByDBName("bagique").Table("bagique_companies").Where("deleted_at IS NULL").Select("company_name as label,id as value").Scan(&companyId)
	global.MustGetGlobalDBByDBName("bagique").Table("bagique_products").Where("deleted_at IS NULL").Select("product_name as label,id as value").Scan(&productId)
	res["companyId"] = companyId
	res["productId"] = productId
	return
}

func (purchaseService *PurchaseService) GetPurchasePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// RefuseEvaluate 根据ID驳回估价
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) RefuseEvaluate(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Evaluate{}).Where("id = ?", ID).Update("status", "REFUSE").Error
	return err
}
