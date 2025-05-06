package scene

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	rulesReq "github.com/flipped-aurora/gin-vue-admin/server/model/rules/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/scene"
	sceneReq "github.com/flipped-aurora/gin-vue-admin/server/model/scene/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ScenesApi struct{}

// CreateScenes 创建scenes表
// @Tags Scenes
// @Summary 创建scenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scene.Scenes true "创建scenes表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /scenes/createScenes [post]
func (scenesApi *ScenesApi) CreateScenes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var scenes scene.Scenes
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scenesService.CreateScenes(ctx, &scenes)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteScenes 删除scenes表
// @Tags Scenes
// @Summary 删除scenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scene.Scenes true "删除scenes表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /scenes/deleteScenes [delete]
func (scenesApi *ScenesApi) DeleteScenes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	err := scenesService.DeleteScenes(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteScenesByIds 批量删除scenes表
// @Tags Scenes
// @Summary 批量删除scenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /scenes/deleteScenesByIds [delete]
func (scenesApi *ScenesApi) DeleteScenesByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := scenesService.DeleteScenesByIds(ctx, IDs)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateScenes 更新scenes表
// @Tags Scenes
// @Summary 更新scenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body scene.Scenes true "更新scenes表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /scenes/updateScenes [put]
func (scenesApi *ScenesApi) UpdateScenes(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var scenes scene.Scenes
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scenesService.UpdateScenes(ctx, scenes)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindScenes 用id查询scenes表
// @Tags Scenes
// @Summary 用id查询scenes表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询scenes表"
// @Success 200 {object} response.Response{data=scene.Scenes,msg=string} "查询成功"
// @Router /scenes/findScenes [get]
func (scenesApi *ScenesApi) FindScenes(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	rescenes, err := scenesService.GetScenes(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(rescenes, c)
}

// GetScenesList 分页获取scenes表列表
// @Tags Scenes
// @Summary 分页获取scenes表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query sceneReq.ScenesSearch true "分页获取scenes表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /scenes/getScenesList [get]
func (scenesApi *ScenesApi) GetScenesList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo sceneReq.ScenesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := scenesService.GetScenesInfoList(ctx, pageInfo)
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

// GetScenesPublic 不需要鉴权的scenes表接口
// @Tags Scenes
// @Summary 不需要鉴权的scenes表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /scenes/getScenesPublic [get]
func (scenesApi *ScenesApi) GetScenesPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	scenesService.GetScenesPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的scenes表接口信息",
	}, "获取成功", c)
}

func (scenesApi *ScenesApi) ScenesSwitchChange(c *gin.Context) {
	ctx := c.Request.Context()

	var pageInfo rulesReq.HandSearch
	err := c.ShouldBindJSON(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = scenesService.HandleSwitchChange(ctx, pageInfo)
	if err != nil {
		global.GVA_LOG.Error("修改失败!", zap.Error(err))
		response.FailWithMessage("修改失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(nil, "修改成功", c)
}

func (scenesApi *ScenesApi) GetScenDevicesList(c *gin.Context) {
	ctx := c.Request.Context()
	value := c.Query("id")
	va, err := scenesService.GetScenDevicesList(ctx, value)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(va, "修改成功", c)
}

func (scenesApi *ScenesApi) GetScenFuncList(c *gin.Context) {
	ctx := c.Request.Context()
	value := c.Query("id")
	va, err := scenesService.GetScenFuncList(ctx, value)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithDetailed(va, "修改成功", c)
}
