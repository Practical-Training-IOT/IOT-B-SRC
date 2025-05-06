package scene

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ScenesRouter struct{}

// InitScenesRouter 初始化 scenes表 路由信息
func (s *ScenesRouter) InitScenesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	scenesRouter := Router.Group("scenes").Use(middleware.OperationRecord())
	scenesRouterWithoutRecord := Router.Group("scenes")
	scenesRouterWithoutAuth := PublicRouter.Group("scenes")
	{
		scenesRouter.POST("createScenes", scenesApi.CreateScenes)             // 新建scenes表
		scenesRouter.DELETE("deleteScenes", scenesApi.DeleteScenes)           // 删除scenes表
		scenesRouter.DELETE("deleteScenesByIds", scenesApi.DeleteScenesByIds) // 批量删除scenes表
		scenesRouter.PUT("updateScenes", scenesApi.UpdateScenes)              // 更新scenes表
	}
	{
		scenesRouterWithoutRecord.GET("findScenes", scenesApi.FindScenes)       // 根据ID获取scenes表
		scenesRouterWithoutRecord.GET("getScenesList", scenesApi.GetScenesList) // 获取scenes表列表
	}
	{
		scenesRouterWithoutAuth.GET("getScenesPublic", scenesApi.GetScenesPublic)        // scenes表开放接口
		scenesRouterWithoutAuth.POST("scenesSwitchChange", scenesApi.ScenesSwitchChange) // scenes表开放接口
		scenesRouterWithoutAuth.GET("getScenDevicesList", scenesApi.GetScenDevicesList)  // scenes表开放接口
		scenesRouterWithoutAuth.GET("getScenFuncList", scenesApi.GetScenFuncList)        // scenes表开放接口
	}
}
