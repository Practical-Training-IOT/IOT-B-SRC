package devicePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DevicesRouter struct{}

// InitDevicesRouter 初始化 devices表 路由信息
func (s *DevicesRouter) InitDevicesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	devicesRouter := Router.Group("devices").Use(middleware.OperationRecord())
	devicesRouterWithoutRecord := Router.Group("devices")
	devicesRouterWithoutAuth := PublicRouter.Group("devices")
	{
		devicesRouter.POST("createDevices", devicesApi.CreateDevices)             // 新建devices表
		devicesRouter.DELETE("deleteDevices", devicesApi.DeleteDevices)           // 删除devices表
		devicesRouter.DELETE("deleteDevicesByIds", devicesApi.DeleteDevicesByIds) // 批量删除devices表
		devicesRouter.PUT("updateDevices", devicesApi.UpdateDevices)              // 更新devices表
	}
	{
		devicesRouterWithoutRecord.GET("findDevices", devicesApi.FindDevices)       // 根据ID获取devices表
		devicesRouterWithoutRecord.GET("getDevicesList", devicesApi.GetDevicesList) // 获取devices表列表
	}
	{
		devicesRouterWithoutAuth.GET("getDevicesPublic", devicesApi.GetDevicesPublic)       // devices表开放接口
		devicesRouterWithoutAuth.GET("getProductGroups", devicesApi.GetProductGroups)       // devices表开放接口
		devicesRouterWithoutAuth.GET("getProductTwoGroups", devicesApi.GetProductTwoGroups) // devices表开放接口
		devicesRouterWithoutAuth.GET("getDriverGroups", devicesApi.GetDriverGroups)         // devices表开放接口
	}
}
