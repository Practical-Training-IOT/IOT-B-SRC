package ai

import (
	"github.com/gin-gonic/gin"
)

type AiRouter struct{}

// InitScenesRouter 初始化 scenes表 路由信息
func (s *AiRouter) InitScenesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	scenesRouterWithoutAuth := PublicRouter.Group("ai")
	{
		scenesRouterWithoutAuth.POST("chat", aiApi.Chats)    // scenes表开放接口
		scenesRouterWithoutAuth.POST("change", aiApi.Change) // scenes表开放接口
		scenesRouterWithoutAuth.GET("history", aiApi.History)
		scenesRouterWithoutAuth.POST("oneHistory", aiApi.OneHistory)
	}
}
