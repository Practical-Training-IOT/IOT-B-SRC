package productPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	productPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/productPkg/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductsApi struct{}

// CreateProducts 创建products表
// @Tags Products
// @Summary 创建products表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body productPkg.Products true "创建products表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /products/createProducts [post]
func (productsApi *ProductsApi) CreateProducts(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var products productPkg.Products
	err := c.ShouldBindJSON(&products)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productsService.CreateProducts(ctx, &products)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteProducts 删除products表
// @Tags Products
// @Summary 删除products表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body productPkg.Products true "删除products表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /products/deleteProducts [delete]
func (productsApi *ProductsApi) DeleteProducts(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := productsService.DeleteProducts(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteProductsByIds 批量删除products表
// @Tags Products
// @Summary 批量删除products表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /products/deleteProductsByIds [delete]
func (productsApi *ProductsApi) DeleteProductsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := productsService.DeleteProductsByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateProducts 更新products表
// @Tags Products
// @Summary 更新products表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body productPkg.Products true "更新products表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /products/updateProducts [put]
func (productsApi *ProductsApi) UpdateProducts(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var products productPkg.Products
	err := c.ShouldBindJSON(&products)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = productsService.UpdateProducts(ctx, products)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindProducts 用id查询products表
// @Tags Products
// @Summary 用id查询products表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询products表"
// @Success 200 {object} response.Response{data=productPkg.Products,msg=string} "查询成功"
// @Router /products/findProducts [get]
func (productsApi *ProductsApi) FindProducts(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	reproducts, err := productsService.GetProducts(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(reproducts, c)
}

// GetProductsList 分页获取products表列表
// @Tags Products
// @Summary 分页获取products表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query productPkgReq.ProductsSearch true "分页获取products表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /products/getProductsList [get]
func (productsApi *ProductsApi) GetProductsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo productPkgReq.ProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := productsService.GetProductsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetProductsPublic 不需要鉴权的products表接口
// @Tags Products
// @Summary 不需要鉴权的products表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /products/getProductsPublic [get]
func (productsApi *ProductsApi) GetProductsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	productsService.GetProductsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的products表接口信息",
	}, "获取成功", c)
}

func (productsApi *ProductsApi) GetStandardCategoryList(c *gin.Context) {

	var pageInfo productPkgReq.ProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var cate []productPkg.Category
	var total int64
	global.GVA_DB.Model(&productPkg.Category{}).Count(&total)
	global.GVA_DB.Model(&productPkg.Category{}).Scopes(Paginate(pageInfo.Page, pageInfo.PageSize)).Find(&cate)
	response.OkWithDetailed(response.PageResult{
		List:     cate,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}

func (productsApi *ProductsApi) GetPropertyList(c *gin.Context) {
	categoryName := c.Query("CategoryName")
	var pageInfo productPkgReq.ProductsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var category productPkg.Category
	err = global.GVA_DB.Model(&productPkg.Category{}).Where("category_name = ?", categoryName).First(&category).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	//获得key
	key := category.CategoryKey
	var Property []productPkg.Property
	var total int64
	global.GVA_DB.Model(&productPkg.Property{}).Where("category_key = ?", key).Count(&total)
	err = global.GVA_DB.Model(&productPkg.Property{}).Scopes(Paginate(pageInfo.Page, pageInfo.PageSize)).Where("category_key = ?", key).Find(&Property).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     Property,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)

}

func Paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
