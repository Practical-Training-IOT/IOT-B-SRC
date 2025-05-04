package alertRulePkg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
    alertRulePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type AlertRulesApi struct {}



// CreateAlertRules 创建alertRules表
// @Tags AlertRules
// @Summary 创建alertRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.AlertRules true "创建alertRules表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /alertRules/createAlertRules [post]
func (alertRulesApi *AlertRulesApi) CreateAlertRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var alertRules alertRulePkg.AlertRules
	err := c.ShouldBindJSON(&alertRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alertRules.CreatedBy = utils.GetUserID(c)
	err = alertRulesService.CreateAlertRules(ctx,&alertRules)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteAlertRules 删除alertRules表
// @Tags AlertRules
// @Summary 删除alertRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.AlertRules true "删除alertRules表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /alertRules/deleteAlertRules [delete]
func (alertRulesApi *AlertRulesApi) DeleteAlertRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
    userID := utils.GetUserID(c)
	err := alertRulesService.DeleteAlertRules(ctx,ID,userID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteAlertRulesByIds 批量删除alertRules表
// @Tags AlertRules
// @Summary 批量删除alertRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /alertRules/deleteAlertRulesByIds [delete]
func (alertRulesApi *AlertRulesApi) DeleteAlertRulesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
    userID := utils.GetUserID(c)
	err := alertRulesService.DeleteAlertRulesByIds(ctx,IDs,userID)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateAlertRules 更新alertRules表
// @Tags AlertRules
// @Summary 更新alertRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body alertRulePkg.AlertRules true "更新alertRules表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /alertRules/updateAlertRules [put]
func (alertRulesApi *AlertRulesApi) UpdateAlertRules(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var alertRules alertRulePkg.AlertRules
	err := c.ShouldBindJSON(&alertRules)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
    alertRules.UpdatedBy = utils.GetUserID(c)
	err = alertRulesService.UpdateAlertRules(ctx,alertRules)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindAlertRules 用id查询alertRules表
// @Tags AlertRules
// @Summary 用id查询alertRules表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询alertRules表"
// @Success 200 {object} response.Response{data=alertRulePkg.AlertRules,msg=string} "查询成功"
// @Router /alertRules/findAlertRules [get]
func (alertRulesApi *AlertRulesApi) FindAlertRules(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	realertRules, err := alertRulesService.GetAlertRules(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(realertRules, c)
}
// GetAlertRulesList 分页获取alertRules表列表
// @Tags AlertRules
// @Summary 分页获取alertRules表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query alertRulePkgReq.AlertRulesSearch true "分页获取alertRules表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /alertRules/getAlertRulesList [get]
func (alertRulesApi *AlertRulesApi) GetAlertRulesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo alertRulePkgReq.AlertRulesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := alertRulesService.GetAlertRulesInfoList(ctx,pageInfo)
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

// GetAlertRulesPublic 不需要鉴权的alertRules表接口
// @Tags AlertRules
// @Summary 不需要鉴权的alertRules表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /alertRules/getAlertRulesPublic [get]
func (alertRulesApi *AlertRulesApi) GetAlertRulesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    alertRulesService.GetAlertRulesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的alertRules表接口信息",
    }, "获取成功", c)
}
