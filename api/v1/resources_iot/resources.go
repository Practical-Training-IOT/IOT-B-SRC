package resources_iot

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
    resources_iotReq "github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ResourcesApi struct {}



// CreateResources 创建resources表
// @Tags Resources
// @Summary 创建resources表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body resources_iot.Resources true "创建resources表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /resources/createResources [post]
func (resourcesApi *ResourcesApi) CreateResources(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var resources resources_iot.Resources
	err := c.ShouldBindJSON(&resources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = resourcesService.CreateResources(ctx,&resources)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteResources 删除resources表
// @Tags Resources
// @Summary 删除resources表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body resources_iot.Resources true "删除resources表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /resources/deleteResources [delete]
func (resourcesApi *ResourcesApi) DeleteResources(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := resourcesService.DeleteResources(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteResourcesByIds 批量删除resources表
// @Tags Resources
// @Summary 批量删除resources表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /resources/deleteResourcesByIds [delete]
func (resourcesApi *ResourcesApi) DeleteResourcesByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := resourcesService.DeleteResourcesByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateResources 更新resources表
// @Tags Resources
// @Summary 更新resources表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body resources_iot.Resources true "更新resources表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /resources/updateResources [put]
func (resourcesApi *ResourcesApi) UpdateResources(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var resources resources_iot.Resources
	err := c.ShouldBindJSON(&resources)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = resourcesService.UpdateResources(ctx,resources)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindResources 用id查询resources表
// @Tags Resources
// @Summary 用id查询resources表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询resources表"
// @Success 200 {object} response.Response{data=resources_iot.Resources,msg=string} "查询成功"
// @Router /resources/findResources [get]
func (resourcesApi *ResourcesApi) FindResources(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	reresources, err := resourcesService.GetResources(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(reresources, c)
}
// GetResourcesList 分页获取resources表列表
// @Tags Resources
// @Summary 分页获取resources表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query resources_iotReq.ResourcesSearch true "分页获取resources表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /resources/getResourcesList [get]
func (resourcesApi *ResourcesApi) GetResourcesList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo resources_iotReq.ResourcesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := resourcesService.GetResourcesInfoList(ctx,pageInfo)
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

// GetResourcesPublic 不需要鉴权的resources表接口
// @Tags Resources
// @Summary 不需要鉴权的resources表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /resources/getResourcesPublic [get]
func (resourcesApi *ResourcesApi) GetResourcesPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    resourcesService.GetResourcesPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的resources表接口信息",
    }, "获取成功", c)
}
