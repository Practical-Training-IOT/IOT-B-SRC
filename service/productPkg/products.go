package productPkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	productPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/productPkg/request"
)

type ProductsService struct{}

// CreateProducts 创建products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) CreateProducts(ctx context.Context, products *productPkg.Products) (err error) {
	newPlatform := "cloud"
	products.Platform = &newPlatform
	err = global.GVA_DB.Create(products).Error
	return err
}

// DeleteProducts 删除products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) DeleteProducts(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&productPkg.Products{}, "id = ?", ID).Error
	return err
}

// DeleteProductsByIds 批量删除products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) DeleteProductsByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]productPkg.Products{}, "id in ?", IDs).Error
	return err
}

// UpdateProducts 更新products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) UpdateProducts(ctx context.Context, products productPkg.Products) (err error) {
	err = global.GVA_DB.Model(&productPkg.Products{}).Where("id = ?", products.ID).Updates(&products).Error
	return err
}

// GetProducts 根据ID获取products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) GetProducts(ctx context.Context, ID string) (products productPkg.Products, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&products).Error
	return
}

// GetProductsInfoList 分页获取products表记录
// Author [yourname](https://github.com/yourname)
func (productsService *ProductsService) GetProductsInfoList(ctx context.Context, info productPkgReq.ProductsSearch) (list []productPkg.Products, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&productPkg.Products{})
	var productss []productPkg.Products
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Platform != nil && *info.Platform != "" {
		db = db.Where("platform = ?", *info.Platform)
	}
	if info.ProductName != nil && *info.ProductName != "" {
		db = db.Where("product_name = ?", *info.ProductName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&productss).Error
	return productss, total, err
}
func (productsService *ProductsService) GetProductsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
