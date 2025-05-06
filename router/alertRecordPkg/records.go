package alertRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type RecordsRouter struct {}

// InitRecordsRouter 初始化 告警记录 路由信息
func (s *RecordsRouter) InitRecordsRouter(Router *gin.RouterGroup,PublicRouter *gin.RouterGroup) {
	recordsRouter := Router.Group("records").Use(middleware.OperationRecord())
	recordsRouterWithoutRecord := Router.Group("records")
	recordsRouterWithoutAuth := PublicRouter.Group("records")
	{
		recordsRouter.POST("createRecords", recordsApi.CreateRecords)   // 新建告警记录
		recordsRouter.DELETE("deleteRecords", recordsApi.DeleteRecords) // 删除告警记录
		recordsRouter.DELETE("deleteRecordsByIds", recordsApi.DeleteRecordsByIds) // 批量删除告警记录
		recordsRouter.PUT("updateRecords", recordsApi.UpdateRecords)    // 更新告警记录
	}
	{
		recordsRouterWithoutRecord.GET("findRecords", recordsApi.FindRecords)        // 根据ID获取告警记录
		recordsRouterWithoutRecord.GET("getRecordsList", recordsApi.GetRecordsList)  // 获取告警记录列表
	}
	{
	    recordsRouterWithoutAuth.GET("getRecordsPublic", recordsApi.GetRecordsPublic)  // 告警记录开放接口
	}
}
