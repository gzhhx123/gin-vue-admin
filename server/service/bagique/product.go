package bagique

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/bagique"
	bagiqueReq "github.com/flipped-aurora/gin-vue-admin/server/model/bagique/request"
)

type ProductService struct{}

// CreateProduct 创建产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) CreateProduct(product *bagique.Product) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Create(product).Error
	return err
}

// DeleteProduct 删除产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) DeleteProduct(ID string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&bagique.Product{}, "id = ?", ID).Error
	return err
}

// DeleteProductByIds 批量删除产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) DeleteProductByIds(IDs []string) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Delete(&[]bagique.Product{}, "id in ?", IDs).Error
	return err
}

// UpdateProduct 更新产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) UpdateProduct(product bagique.Product) (err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Product{}).Where("id = ?", product.ID).Updates(&product).Error
	return err
}

// GetProduct 根据ID获取产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) GetProduct(ID string) (product bagique.Product, err error) {
	err = global.MustGetGlobalDBByDBName("bagique").Where("id = ?", ID).First(&product).Error
	return
}

// GetProductInfoList 分页获取产品信息记录
// Author [yourname](https://github.com/yourname)
func (productService *ProductService) GetProductInfoList(info bagiqueReq.ProductSearch) (list []bagique.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.MustGetGlobalDBByDBName("bagique").Model(&bagique.Product{})
	var products []bagique.Product
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.BrandId != nil {
		db = db.Where("brand_id = ?", *info.BrandId)
	}
	if info.ProductName != nil && *info.ProductName != "" {
		db = db.Where("product_name LIKE ?", "%"+*info.ProductName+"%")
	}
	if info.ProductSku != nil && *info.ProductSku != "" {
		db = db.Where("product_sku LIKE ?", "%"+*info.ProductSku+"%")
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
	orderMap["brand_id"] = true
	orderMap["product_name"] = true
	orderMap["product_sku"] = true
	orderMap["refer_price_min"] = true
	orderMap["refer_price_max"] = true
	orderMap["refer_pics"] = true
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

	err = db.Find(&products).Error
	return products, total, err
}
func (productService *ProductService) GetProductDataSource() (res map[string][]map[string]any, err error) {
	res = make(map[string][]map[string]any)

	brandId := make([]map[string]any, 0)

	global.MustGetGlobalDBByDBName("bagique").Table("bagique_brands").Where("deleted_at IS NULL").Select("brand_name as label,id as value").Scan(&brandId)
	res["brandId"] = brandId
	return
}
func (productService *ProductService) GetProductPublic() {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
