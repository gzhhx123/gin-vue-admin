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
func (evaluateService *EvaluateService) DeleteEvaluate(ID string, TYPE string) (err error) {
	if TYPE == "HARD" {
		err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Delete(&bagique.Evaluate{}, "id = ?", ID).Error
	} else {
		err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Evaluate{}, "id = ?", ID).Error
	}
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
	tx := global.MustGetGlobalDBByDBName("bagique").Begin()
	//更新估价主记录
	err = tx.Model(&bagique.Evaluate{}).Where("id = ?", evaluate.ID).Updates(&evaluate).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//更新公司估价记录
	//1.获取数据库已存在数据
	var evaluatePrices []bagique.EvaluatePrice
	err = tx.Where("evaluate_id = ?", evaluate.ID).Find(&evaluatePrices).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//2.构建ID映射方便比对
	existingMap := make(map[uint]*bagique.EvaluatePrice)
	for _, v := range evaluatePrices {
		existingMap[v.ID] = &v
	}
	// 3. 遍历前端传递的数据 进行相应的操作
	for _, v := range evaluate.EvaluatePrices {
		if existingPrice, exists := existingMap[v.ID]; exists {
			if existingPrice.CompanyId != v.CompanyId || existingPrice.Price != v.Price || existingPrice.Fee != v.Fee || existingPrice.Remark != v.Remark {
				err = tx.Model(&bagique.EvaluatePrice{}).Where("id = ?", v.ID).Updates(&v).Error
				if err != nil {
					tx.Rollback()
					return err
				}
			}
			delete(existingMap, v.ID)
		} else {
			//新增
			v.EvaluateId = new(int)
			*v.EvaluateId = int(evaluate.ID)
			err = tx.Create(&v).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//4.删除数据库中未前端未提交的数据
	for k := range existingMap {
		err = tx.Delete(&bagique.EvaluatePrice{}, "id = ?", k).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return nil
}

// GetEvaluate 根据ID获取估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) GetEvaluate(ID string) (evaluate bagique.Evaluate, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Preload("EvaluatePrices").Where("id = ?", ID).First(&evaluate).Error
	return
}

// GetEvaluateInfoList 分页获取估价信息记录
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) GetEvaluateInfoList(info bagiqueReq.EvaluateSearch) (list []bagique.Evaluate, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Preload("EvaluatePrices").Model(&bagique.Evaluate{})
	var evaluates []bagique.Evaluate
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.IsRemove != nil && *info.IsRemove == true {
		db = db.Unscoped()
	}
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
	orderMap["created_at"] = true
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
	companyId := make([]map[string]any, 0)

	global.MustGetGlobalDBByDBName("bagique").Table("bagique_companies").Where("deleted_at IS NULL").Select("company_name as label,id as value").Scan(&companyId)
	res["companyId"] = companyId
	return
}
func (evaluateService *EvaluateService) GetEvaluatePublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

// RestoreEvaluate 恢复单条evaluate的数据，也就是将deleted_at设置为null
// Author [yourname](https://github.com/yourname)
func (evaluateService *EvaluateService) RestoreEvaluate(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Unscoped().Model(&bagique.Evaluate{}).Where("id = ?", ID).Update("deleted_at", nil).Error
	return err
}
