package alertRulePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlarmsRouter struct{}

// InitAlarmsRouter 初始化 告警规则 路由信息
func (s *AlarmsRouter) InitAlarmsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	alarmsRouter := Router.Group("alarms").Use(middleware.OperationRecord())
	alarmsRouterWithoutRecord := Router.Group("alarms")
	alarmsRouterWithoutAuth := PublicRouter.Group("alarms")
	{
		alarmsRouter.POST("createAlarms", alarmsApi.CreateAlarms)             // 新建告警规则
		alarmsRouter.DELETE("deleteAlarms", alarmsApi.DeleteAlarms)           // 删除告警规则
		alarmsRouter.DELETE("deleteAlarmsByIds", alarmsApi.DeleteAlarmsByIds) // 批量删除告警规则
		alarmsRouter.PUT("updateAlarms", alarmsApi.UpdateAlarms)              // 更新告警规则
	}
	{
		alarmsRouterWithoutRecord.GET("findAlarms", alarmsApi.FindAlarms)       // 根据ID获取告警规则
		alarmsRouterWithoutRecord.GET("getAlarmsList", alarmsApi.GetAlarmsList) // 获取告警规则列表
	}
	{
		alarmsRouterWithoutAuth.GET("getAlarmsPublic", alarmsApi.GetAlarmsPublic)     // 告警规则开放接口
		alarmsRouterWithoutAuth.GET("getAllProductList", alarmsApi.GetAllProductList) // 告警规则开放接口
		alarmsRouterWithoutAuth.GET("getEquipmentList", alarmsApi.GetEquipmentList)   // 告警规则开放接口
	}
}
