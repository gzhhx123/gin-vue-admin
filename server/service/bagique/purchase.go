package bagique

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
	"time"
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
	trackCompanyId := make([]map[string]any, 0)
	global.MustGetGlobalDBByDBName("bagique").Table("bagique_companies").Where("deleted_at IS NULL").Select("company_name as label,id as value").Scan(&companyId)
	global.MustGetGlobalDBByDBName("bagique").Table("bagique_products").Where("deleted_at IS NULL").Select("product_name as label,id as value").Scan(&productId)
	global.MustGetGlobalDBByDBName("bagique").Table("bagique_track_companies").Where("deleted_at IS NULL").Select("company_name as label,id as value").Scan(&trackCompanyId)
	res["companyId"] = companyId
	res["productId"] = productId
	res["trackCompanyId"] = trackCompanyId
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

// FindPurchaseTrackNosById 根据采购ID获取物流单号
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) FindPurchaseTrackNosById(ID string) (list []bagique.TrackNo, err error) {
	var arr []bagique.TrackNo
	arr = make([]bagique.TrackNo, 0)
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.TrackNo{}).Where("purchase_id = ?", ID).Order("sort").Find(&arr).Error
	return arr, err
}

// UpdateTrackNo 根据采购ID更新物流单号
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) UpdateTrackNo(ID uint, trackNos []bagique.TrackNo) (err error) {
	tx := global.MustGetGlobalDBByDBName("bagique").Begin()
	//1.获取数据库已存在的物流单号
	var oldTrackNos []bagique.TrackNo
	err = tx.Where("purchase_id = ?", ID).Find(&oldTrackNos).Error
	if err != nil {
		tx.Rollback()
		return err
	}
	//2.构建ID映射方便比对
	existingMap := make(map[uint]*bagique.TrackNo)
	for _, v := range oldTrackNos {
		existingMap[v.ID] = &v
	}
	// 3. 遍历前端传递的数据 进行相应的操作
	for _, v := range trackNos {
		if existingTrackNo, exists := existingMap[v.ID]; exists {
			if *existingTrackNo.OrderSn != *v.OrderSn || *existingTrackNo.TrackCompanyId != *v.TrackCompanyId || *existingTrackNo.Remark != *v.Remark || *existingTrackNo.Sort != *v.Sort || *existingTrackNo.Status != *v.Status || *existingTrackNo.Type != *v.Type {
				err = tx.Model(&bagique.TrackNo{}).Where("id = ?", v.ID).Updates(&v).Error
				if err != nil {
					tx.Rollback()
					return err
				}
			}
			delete(existingMap, v.ID)
		} else {
			//新增
			v.PurchaseId = new(int)
			*v.PurchaseId = int(ID)
			err = tx.Create(&v).Error
			if err != nil {
				tx.Rollback()
				return
			}
		}
	}
	//4.删除数据库中未前端未提交的数据
	for k := range existingMap {
		err = tx.Delete(&bagique.TrackNo{}, "id = ?", k).Error
		if err != nil {
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	return nil
}

type Item struct {
	Content string    `json:"content"`
	Time    time.Time `json:"time"`
}
type Axis struct {
	List []Item `json:"list"`
}

// PurchaseTimeAxis 查看采购时间轴
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) PurchaseTimeAxis(ID string) (axis_ Axis, err error) {
	var axis Axis
	axis.List = make([]Item, 0)
	//先查询出purchase
	var purchase bagique.Purchase
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&purchase).Error
	//再根据purchase的evaluate_price_id去查询evaluate_price
	var evaluatePrice bagique.EvaluatePrice
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", *purchase.EvaluatePriceId).First(&evaluatePrice).Error
	//再根据evaluate_price的evaluate_id去查询evaluate
	var evaluate bagique.Evaluate
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", evaluatePrice.EvaluateId).First(&evaluate).Error
	if evaluate.ID != 0 {
		axis.List = append(axis.List, Item{Content: "添加估价", Time: evaluate.CreatedAt})
		axis.List = append(axis.List, Item{Content: "完成估价", Time: evaluate.UpdatedAt})
	}
	if purchase.ID != 0 {
		axis.List = append(axis.List, Item{Content: "添加采购", Time: purchase.CreatedAt})
		if *purchase.Status == "FINISH" {
			axis.List = append(axis.List, Item{Content: "完成采购", Time: purchase.UpdatedAt})
		}
		if *purchase.Status == "CANCEL" {
			axis.List = append(axis.List, Item{Content: "取消采购", Time: purchase.UpdatedAt})
		}
	}
	axis_ = axis
	return axis_, err
}

// FinishPurchase 根据ID完成采购
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) FinishPurchase(ID string) (err error) {
	// 请在这里实现自己的业务逻辑
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Purchase{}).Where("id = ?", ID).Update("status", "FINISH").Error
	return err
}

// CancelPurchase 根据ID取消采购
// Author [yourname](https://github.com/yourname)
func (purchaseService *PurchaseService) CancelPurchase(ID string) (err error) {
	// 请在这里实现自己的业务逻辑
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Purchase{}).Where("id = ?", ID).Update("status", "CANCEL").Error
	return err
}
