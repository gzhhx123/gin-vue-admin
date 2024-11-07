package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type EvaluateService struct{}

// CreateEvaluate 创建估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) CreateEvaluate(evaluate *bagique.Evaluate) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(evaluate).Error
	return err
}

// DeleteEvaluate 删除估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) DeleteEvaluate(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Evaluate{}, "id = ?", ID).Error
	return err
}

// DeleteEvaluateByIds 批量删除估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) DeleteEvaluateByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Evaluate{}, "id in ?", IDs).Error
	return err
}

// UpdateEvaluate 更新估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) UpdateEvaluate(evaluate bagique.Evaluate) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Evaluate{}).Where("id = ?", evaluate.ID).Updates(&evaluate).Error
	return err
}

// GetEvaluate 根据ID获取估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) GetEvaluate(ID string) (evaluate bagique.Evaluate, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&evaluate).Error
	return
}

// GetEvaluateInfoList 分页获取估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) GetEvaluateInfoList(info bagiqueReq.EvaluateSearch) (list []bagique.Evaluate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Evaluate{})
	var evaluates []bagique.Evaluate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductId != nil {
		db = db.Where("product_id = ?", *info.ProductId)
	}
	if info.SellerId != nil {
		db = db.Where("seller_id = ?", *info.SellerId)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
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
	orderMap["product_id"] = true
	orderMap["seller_id"] = true
	orderMap["evaluate_pics"] = true
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

	err = db.Find(&evaluates).Error
	return evaluates, total, err
}
func (evaluateService *EvaluateService) GetEvaluateDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	productId := make([]map[string]any, 0)

	global.MustGetGlobalDBByDBName("bagique").Table("bagique_products").Where("deleted_at IS NULL").Select("product_name as label,id as value").Scan(&productId)
	res["productId"] = productId
	sellerId := make([]map[string]any, 0)

	global.MustGetGlobalDBByDBName("bagique").Table("bagique_sellers").Where("deleted_at IS NULL").Select("seller_name as label,id as value").Scan(&sellerId)
	res["sellerId"] = sellerId
	return
}
func (evaluateService *EvaluateService) GetEvaluatePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
