package resources_iot

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ResourcesRouter struct {}

// InitResourcesRouter 初始化 resources表 路由信息
func (s *ResourcesRouter) InitResourcesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	resourcesRouter := Router.Group("resources").Use(middleware.OperationRecord())
	resourcesRouterWithoutRecord := Router.Group("resources")
	resourcesRouterWithoutAuth := PublicRouter.Group("resources")
	{
		resourcesRouter.POST("createResources", resourcesApi.CreateResources)   // 新建resources表
		resourcesRouter.DELETE("deleteResources", resourcesApi.DeleteResources) // 删除resources表
		resourcesRouter.DELETE("deleteResourcesByIds", resourcesApi.DeleteResourcesByIds) // 批量删除resources表
		resourcesRouter.PUT("updateResources", resourcesApi.UpdateResources)    // 更新resources表
	}
	{
		resourcesRouterWithoutRecord.GET("findResources", resourcesApi.FindResources)        // 根据ID获取resources表
		resourcesRouterWithoutRecord.GET("getResourcesList", resourcesApi.GetResourcesList)  // 获取resources表列表
	}
	{
	    resourcesRouterWithoutAuth.GET("getResourcesPublic", resourcesApi.GetResourcesPublic)  // resources表开放接口
	}
}
