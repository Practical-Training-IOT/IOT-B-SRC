package rules

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RuleInfoRouter struct{}

// InitRuleInfoRouter 初始化 ruleInfo表 路由信息
func (s *RuleInfoRouter) InitRuleInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	ruleInfoRouter := Router.Group("ruleInfo").Use(middleware.OperationRecord())
	ruleInfoRouterWithoutRecord := Router.Group("ruleInfo")
	ruleInfoRouterWithoutAuth := PublicRouter.Group("ruleInfo")
	{
		ruleInfoRouter.POST("createRuleInfo", ruleInfoApi.CreateRuleInfo)             // 新建ruleInfo表
		ruleInfoRouter.DELETE("deleteRuleInfo", ruleInfoApi.DeleteRuleInfo)           // 删除ruleInfo表
		ruleInfoRouter.DELETE("deleteRuleInfoByIds", ruleInfoApi.DeleteRuleInfoByIds) // 批量删除ruleInfo表
		ruleInfoRouter.PUT("updateRuleInfo", ruleInfoApi.UpdateRuleInfo)              // 更新ruleInfo表
	}
	{
		ruleInfoRouterWithoutRecord.GET("findRuleInfo", ruleInfoApi.FindRuleInfo)       // 根据ID获取ruleInfo表
		ruleInfoRouterWithoutRecord.GET("getRuleInfoList", ruleInfoApi.GetRuleInfoList) // 获取ruleInfo表列表
	}
	{
		ruleInfoRouterWithoutAuth.GET("getRuleInfoPublic", ruleInfoApi.GetRuleInfoPublic)    // ruleInfo表开放接口
		ruleInfoRouterWithoutAuth.POST("handleSwitchChange", ruleInfoApi.HandleSwitchChange) // ruleInfo表开放接口
	}
}
