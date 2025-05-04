package myDrivePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MyDriversRouter struct {}

// InitMyDriversRouter 初始化 myDrivers表 路由信息
func (s *MyDriversRouter) InitMyDriversRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	myDriversRouter := Router.Group("myDrivers").Use(middleware.OperationRecord())
	myDriversRouterWithoutRecord := Router.Group("myDrivers")
	myDriversRouterWithoutAuth := PublicRouter.Group("myDrivers")
	{
		myDriversRouter.POST("createMyDrivers", myDriversApi.CreateMyDrivers)   // 新建myDrivers表
		myDriversRouter.DELETE("deleteMyDrivers", myDriversApi.DeleteMyDrivers) // 删除myDrivers表
		myDriversRouter.DELETE("deleteMyDriversByIds", myDriversApi.DeleteMyDriversByIds) // 批量删除myDrivers表
		myDriversRouter.PUT("updateMyDrivers", myDriversApi.UpdateMyDrivers)    // 更新myDrivers表
	}
	{
		myDriversRouterWithoutRecord.GET("findMyDrivers", myDriversApi.FindMyDrivers)        // 根据ID获取myDrivers表
		myDriversRouterWithoutRecord.GET("getMyDriversList", myDriversApi.GetMyDriversList)  // 获取myDrivers表列表
	}
	{
	    myDriversRouterWithoutAuth.GET("getMyDriversPublic", myDriversApi.GetMyDriversPublic)  // myDrivers表开放接口
	}
}
