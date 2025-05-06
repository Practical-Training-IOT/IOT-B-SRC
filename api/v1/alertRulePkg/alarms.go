package alertRulePkg

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
	alertRulePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AlarmsApi struct{}

// CreateAlarms 创建告警规则
// @Tags Alarms
// @Summary 创建告警规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.Alarms true "创建告警规则"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alarms/createAlarms [post]
func (alarmsApi *AlarmsApi) CreateAlarms(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var alarms alertRulePkg.Alarms
	err := c.ShouldBindJSON(&alarms)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	alarms.CreatedBy = utils.GetUserID(c)
	err = alarmsService.CreateAlarms(ctx, &alarms)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteAlarms 删除告警规则
// @Tags Alarms
// @Summary 删除告警规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.Alarms true "删除告警规则"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alarms/deleteAlarms [delete]
func (alarmsApi *AlarmsApi) DeleteAlarms(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := alarmsService.DeleteAlarms(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlarmsByIds 批量删除告警规则
// @Tags Alarms
// @Summary 批量删除告警规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alarms/deleteAlarmsByIds [delete]
func (alarmsApi *AlarmsApi) DeleteAlarmsByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := alarmsService.DeleteAlarmsByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlarms 更新告警规则
// @Tags Alarms
// @Summary 更新告警规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.Alarms true "更新告警规则"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alarms/updateAlarms [put]
func (alarmsApi *AlarmsApi) UpdateAlarms(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var alarms alertRulePkg.Alarms
	err := c.ShouldBindJSON(&alarms)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	alarms.UpdatedBy = utils.GetUserID(c)
	err = alarmsService.UpdateAlarms(ctx, alarms)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlarms 用id查询告警规则
// @Tags Alarms
// @Summary 用id查询告警规则
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询告警规则"
// @Success 200 {object} response.Response{data=alertRulePkg.Alarms,msg=string} "查询成功"
// @Router /alarms/findAlarms [get]
func (alarmsApi *AlarmsApi) FindAlarms(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	realarms, err := alarmsService.GetAlarms(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(realarms, c)
}

// GetAlarmsList 分页获取告警规则列表
// @Tags Alarms
// @Summary 分页获取告警规则列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query alertRulePkgReq.AlarmsSearch true "分页获取告警规则列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alarms/getAlarmsList [get]
func (alarmsApi *AlarmsApi) GetAlarmsList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo alertRulePkgReq.AlarmsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alarmsService.GetAlarmsInfoList(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetAlarmsPublic 不需要鉴权的告警规则接口
// @Tags Alarms
// @Summary 不需要鉴权的告警规则接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alarms/getAlarmsPublic [get]
func (alarmsApi *AlarmsApi) GetAlarmsPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	alarmsService.GetAlarmsPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的告警规则接口信息",
	}, "获取成功", c)
}

func (alarmsApi *AlarmsApi) GetAllProductList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	list, err := alarmsService.GetProductList(ctx)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: list,
	}, "获取成功", c)
}

func (alarmsApi *AlarmsApi) GetEquipmentList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()
	ID := c.Query("ID")
	fmt.Println(ID)
	realarms, err := alarmsService.GetEquipment(ctx, "16")
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(realarms, c)
}
