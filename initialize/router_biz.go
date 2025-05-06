package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router"
	"github.com/gin-gonic/gin"
)

func holder(routers ...*gin.RouterGroup) {
	_ = routers
	_ = router.RouterGroupApp
}
func initBizRouter(routers ...*gin.RouterGroup) {
	privateGroup := routers[0]
	publicGroup := routers[1]
	holder(publicGroup, privateGroup)
	holder(publicGroup, privateGroup)
	{
		alertRulePkgRouter := router.RouterGroupApp.AlertRulePkg
		alertRulePkgRouter.InitAlertRulesRouter(privateGroup, publicGroup)
	}
	{
		alertRecordPkgRouter := router.RouterGroupApp.AlertRecordPkg
		alertRecordPkgRouter.InitAlertsRouter(privateGroup, publicGroup)
	}
	{
		productPkgRouter := router.RouterGroupApp.ProductPkg
		productPkgRouter.InitProductsRouter(privateGroup, publicGroup)
	}
	{
		devicePkgRouter := router.RouterGroupApp.DevicePkg
		devicePkgRouter.InitDevicesRouter(privateGroup, publicGroup)
	}
	{
		driverPkgRouter := router.RouterGroupApp.DriverPkg
		driverPkgRouter.InitDriversRouter(privateGroup, publicGroup)
	}
	{
		myDrivePkgRouter := router.RouterGroupApp.MyDrivePkg
		myDrivePkgRouter.InitMyDriversRouter(privateGroup, publicGroup)
	}
	{
		resources_iotRouter := router.RouterGroupApp.Resources_iot
		resources_iotRouter.InitResourcesRouter(privateGroup, publicGroup)
	}
	{
		rulesRouter := router.RouterGroupApp.Rules
		rulesRouter.InitRuleInfoRouter(privateGroup, publicGroup)
	} // 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	// 占位方法，保证文件可以正确加载，避免go空变量检测报错，请勿删除。
	{
		sceneRouter := router.RouterGroupApp.Scene
		sceneRouter.InitScenesRouter(privateGroup, publicGroup)
	}
	{
		aiRouter := router.RouterGroupApp.AI
		aiRouter.InitScenesRouter(privateGroup, publicGroup)
	}
}
