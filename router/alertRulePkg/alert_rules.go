package alertRulePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type AlertRulesRouter struct {}

// InitAlertRulesRouter 初始化 alertRules表 路由信息
func (s *AlertRulesRouter) InitAlertRulesRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	alertRulesRouter := Router.Group("alertRules").Use(middleware.OperationRecord())
	alertRulesRouterWithoutRecord := Router.Group("alertRules")
	alertRulesRouterWithoutAuth := PublicRouter.Group("alertRules")
	{
		alertRulesRouter.POST("createAlertRules", alertRulesApi.CreateAlertRules)   // 新建alertRules表
		alertRulesRouter.DELETE("deleteAlertRules", alertRulesApi.DeleteAlertRules) // 删除alertRules表
		alertRulesRouter.DELETE("deleteAlertRulesByIds", alertRulesApi.DeleteAlertRulesByIds) // 批量删除alertRules表
		alertRulesRouter.PUT("updateAlertRules", alertRulesApi.UpdateAlertRules)    // 更新alertRules表
	}
	{
		alertRulesRouterWithoutRecord.GET("findAlertRules", alertRulesApi.FindAlertRules)        // 根据ID获取alertRules表
		alertRulesRouterWithoutRecord.GET("getAlertRulesList", alertRulesApi.GetAlertRulesList)  // 获取alertRules表列表
	}
	{
	    alertRulesRouterWithoutAuth.GET("getAlertRulesPublic", alertRulesApi.GetAlertRulesPublic)  // alertRules表开放接口
	}
}
