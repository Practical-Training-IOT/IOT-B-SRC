package alertRecordPkg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
    alertRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AlertsApi struct {}



// CreateAlerts 创建alerts表
// @Tags Alerts
// @Summary 创建alerts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRecordPkg.Alerts true "创建alerts表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alerts/createAlerts [post]
func (alertsApi *AlertsApi) CreateAlerts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var alerts alertRecordPkg.Alerts
	err := c.ShouldBindJSON(&alerts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alerts.CreatedBy = utils.GetUserID(c)
	err = alertsService.CreateAlerts(ctx,&alerts)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAlerts 删除alerts表
// @Tags Alerts
// @Summary 删除alerts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRecordPkg.Alerts true "删除alerts表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alerts/deleteAlerts [delete]
func (alertsApi *AlertsApi) DeleteAlerts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := alertsService.DeleteAlerts(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlertsByIds 批量删除alerts表
// @Tags Alerts
// @Summary 批量删除alerts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alerts/deleteAlertsByIds [delete]
func (alertsApi *AlertsApi) DeleteAlertsByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := alertsService.DeleteAlertsByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlerts 更新alerts表
// @Tags Alerts
// @Summary 更新alerts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRecordPkg.Alerts true "更新alerts表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alerts/updateAlerts [put]
func (alertsApi *AlertsApi) UpdateAlerts(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var alerts alertRecordPkg.Alerts
	err := c.ShouldBindJSON(&alerts)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alerts.UpdatedBy = utils.GetUserID(c)
	err = alertsService.UpdateAlerts(ctx,alerts)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlerts 用id查询alerts表
// @Tags Alerts
// @Summary 用id查询alerts表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询alerts表"
// @Success 200 {object} response.Response{data=alertRecordPkg.Alerts,msg=string} "查询成功"
// @Router /alerts/findAlerts [get]
func (alertsApi *AlertsApi) FindAlerts(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	realerts, err := alertsService.GetAlerts(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(realerts, c)
}
// GetAlertsList 分页获取alerts表列表
// @Tags Alerts
// @Summary 分页获取alerts表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query alertRecordPkgReq.AlertsSearch true "分页获取alerts表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alerts/getAlertsList [get]
func (alertsApi *AlertsApi) GetAlertsList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo alertRecordPkgReq.AlertsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alertsService.GetAlertsInfoList(ctx,pageInfo)
	if err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败:" + err.Error(), c)
        return
    }
    response.OkWithDetailed(response.PageResult{
        List:     list,
        Total:    total,
        Page:     pageInfo.Page,
        PageSize: pageInfo.PageSize,
    }, "获取成功", c)
}

// GetAlertsPublic 不需要鉴权的alerts表接口
// @Tags Alerts
// @Summary 不需要鉴权的alerts表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alerts/getAlertsPublic [get]
func (alertsApi *AlertsApi) GetAlertsPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    alertsService.GetAlertsPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的alerts表接口信息",
    }, "获取成功", c)
}
