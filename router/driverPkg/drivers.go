package driverPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type DriversRouter struct {}

// InitDriversRouter 初始化 drivers表 路由信息
func (s *DriversRouter) InitDriversRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	driversRouter := Router.Group("drivers").Use(middleware.OperationRecord())
	driversRouterWithoutRecord := Router.Group("drivers")
	driversRouterWithoutAuth := PublicRouter.Group("drivers")
	{
		driversRouter.POST("createDrivers", driversApi.CreateDrivers)   // 新建drivers表
		driversRouter.DELETE("deleteDrivers", driversApi.DeleteDrivers) // 删除drivers表
		driversRouter.DELETE("deleteDriversByIds", driversApi.DeleteDriversByIds) // 批量删除drivers表
		driversRouter.PUT("updateDrivers", driversApi.UpdateDrivers)    // 更新drivers表
	}
	{
		driversRouterWithoutRecord.GET("findDrivers", driversApi.FindDrivers)        // 根据ID获取drivers表
		driversRouterWithoutRecord.GET("getDriversList", driversApi.GetDriversList)  // 获取drivers表列表
	}
	{
	    driversRouterWithoutAuth.GET("getDriversPublic", driversApi.GetDriversPublic)  // drivers表开放接口
	}
}
