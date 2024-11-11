package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type EvaluatePriceService struct{}

// CreateEvaluatePrice 创建公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) CreateEvaluatePrice(evaluatePrice *bagique.EvaluatePrice) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(evaluatePrice).Error
	return err
}

// DeleteEvaluatePrice 删除公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) DeleteEvaluatePrice(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.EvaluatePrice{}, "id = ?", ID).Error
	return err
}

// DeleteEvaluatePriceByIds 批量删除公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) DeleteEvaluatePriceByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.EvaluatePrice{}, "id in ?", IDs).Error
	return err
}

// UpdateEvaluatePrice 更新公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) UpdateEvaluatePrice(evaluatePrice bagique.EvaluatePrice) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.EvaluatePrice{}).Where("id = ?", evaluatePrice.ID).Updates(&evaluatePrice).Error
	return err
}

// GetEvaluatePrice 根据ID获取公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) GetEvaluatePrice(ID string) (evaluatePrice bagique.EvaluatePrice, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&evaluatePrice).Error
	return
}

// GetEvaluatePriceInfoList 分页获取公司估价记录
// Author [yourname](https://github.com/yourname)
func (evaluatePriceService *EvaluatePriceService) GetEvaluatePriceInfoList(info bagiqueReq.EvaluatePriceSearch) (list []bagique.EvaluatePrice, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.EvaluatePrice{})
	var evaluatePrices []bagique.EvaluatePrice
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyId != nil {
		db = db.Where("company_id = ?", *info.CompanyId)
	}
	if info.StartEvaluateId != nil && info.EndEvaluateId != nil {
		db = db.Where("evaluate_id BETWEEN ? AND ? ", info.StartEvaluateId, info.EndEvaluateId)
	}
	if info.StartPrice != nil && info.EndPrice != nil {
		db = db.Where("price BETWEEN ? AND ? ", info.StartPrice, info.EndPrice)
	}
	if info.StartFee != nil && info.EndFee != nil {
		db = db.Where("fee BETWEEN ? AND ? ", info.StartFee, info.EndFee)
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
	orderMap["company_id"] = true
	orderMap["evaluate_id"] = true
	orderMap["price"] = true
	orderMap["fee"] = true
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

	err = db.Find(&evaluatePrices).Error
	return evaluatePrices, total, err
}
func (evaluatePriceService *EvaluatePriceService) GetEvaluatePriceDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	companyId := make([]map[string]any, 0)

	global.MustGetGlobalDBByDBName("bagique").Table("bagique_companies").Where("deleted_at IS NULL").Select("company_name as label,id as value").Scan(&companyId)
	res["companyId"] = companyId
	return
}
func (evaluatePriceService *EvaluatePriceService) GetEvaluatePricePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
