package productPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductsRouter struct{}

// InitProductsRouter 初始化 products表 路由信息
func (s *ProductsRouter) InitProductsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	productsRouter := Router.Group("products").Use(middleware.OperationRecord())
	productsRouterWithoutRecord := Router.Group("products")
	productsRouterWithoutAuth := PublicRouter.Group("products")
	{
		productsRouter.POST("createProducts", productsApi.CreateProducts)             // 新建products表
		productsRouter.DELETE("deleteProducts", productsApi.DeleteProducts)           // 删除products表
		productsRouter.DELETE("deleteProductsByIds", productsApi.DeleteProductsByIds) // 批量删除products表
		productsRouter.PUT("updateProducts", productsApi.UpdateProducts)              // 更新products表
	}
	{
		productsRouterWithoutRecord.GET("findProducts", productsApi.FindProducts)       // 根据ID获取products表
		productsRouterWithoutRecord.GET("getProductsList", productsApi.GetProductsList) // 获取products表列表
	}
	{
		productsRouterWithoutAuth.GET("getProductsPublic", productsApi.GetProductsPublic)             // products表开放接口
		productsRouterWithoutAuth.GET("getStandardCategoryList", productsApi.GetStandardCategoryList) // products表开放接口
	}
}
