package alertRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlertsRouter struct {}

// InitAlertsRouter 初始化 alerts表 路由信息
func (s *AlertsRouter) InitAlertsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	alertsRouter := Router.Group("alerts").Use(middleware.OperationRecord())
	alertsRouterWithoutRecord := Router.Group("alerts")
	alertsRouterWithoutAuth := PublicRouter.Group("alerts")
	{
		alertsRouter.POST("createAlerts", alertsApi.CreateAlerts)   // 新建alerts表
		alertsRouter.DELETE("deleteAlerts", alertsApi.DeleteAlerts) // 删除alerts表
		alertsRouter.DELETE("deleteAlertsByIds", alertsApi.DeleteAlertsByIds) // 批量删除alerts表
		alertsRouter.PUT("updateAlerts", alertsApi.UpdateAlerts)    // 更新alerts表
	}
	{
		alertsRouterWithoutRecord.GET("findAlerts", alertsApi.FindAlerts)        // 根据ID获取alerts表
		alertsRouterWithoutRecord.GET("getAlertsList", alertsApi.GetAlertsList)  // 获取alerts表列表
	}
	{
	    alertsRouterWithoutAuth.GET("getAlertsPublic", alertsApi.GetAlertsPublic)  // alerts表开放接口
	}
}
