package devicePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg"
	devicePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type DevicesApi struct{}

// CreateDevices 创建devices表
// @Tags Devices
// @Summary 创建devices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body devicePkg.Devices true "创建devices表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /devices/createDevices [post]
func (devicesApi *DevicesApi) CreateDevices(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var devices devicePkg.Devices
	err := c.ShouldBindJSON(&devices)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	devices.CreatedBy = utils.GetUserID(c)
	err = devicesService.CreateDevices(ctx, &devices)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteDevices 删除devices表
// @Tags Devices
// @Summary 删除devices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body devicePkg.Devices true "删除devices表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /devices/deleteDevices [delete]
func (devicesApi *DevicesApi) DeleteDevices(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	userID := utils.GetUserID(c)
	err := devicesService.DeleteDevices(ctx, ID, userID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteDevicesByIds 批量删除devices表
// @Tags Devices
// @Summary 批量删除devices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /devices/deleteDevicesByIds [delete]
func (devicesApi *DevicesApi) DeleteDevicesByIds(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	userID := utils.GetUserID(c)
	err := devicesService.DeleteDevicesByIds(ctx, IDs, userID)
	if err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateDevices 更新devices表
// @Tags Devices
// @Summary 更新devices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body devicePkg.Devices true "更新devices表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /devices/updateDevices [put]
func (devicesApi *DevicesApi) UpdateDevices(c *gin.Context) {
	// 从ctx获取标准context进行业务行为
	ctx := c.Request.Context()

	var devices devicePkg.Devices
	err := c.ShouldBindJSON(&devices)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	devices.UpdatedBy = utils.GetUserID(c)
	err = devicesService.UpdateDevices(ctx, devices)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindDevices 用id查询devices表
// @Tags Devices
// @Summary 用id查询devices表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询devices表"
// @Success 200 {object} response.Response{data=devicePkg.Devices,msg=string} "查询成功"
// @Router /devices/findDevices [get]
func (devicesApi *DevicesApi) FindDevices(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	ID := c.Query("ID")
	redevices, err := devicesService.GetDevices(ctx, ID)
	if err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:"+err.Error(), c)
		return
	}
	response.OkWithData(redevices, c)
}

// GetDevicesList 分页获取devices表列表
// @Tags Devices
// @Summary 分页获取devices表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query devicePkgReq.DevicesSearch true "分页获取devices表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /devices/getDevicesList [get]
func (devicesApi *DevicesApi) GetDevicesList(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	var pageInfo devicePkgReq.DevicesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := devicesService.GetDevicesInfoList(ctx, pageInfo)
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

// GetDevicesPublic 不需要鉴权的devices表接口
// @Tags Devices
// @Summary 不需要鉴权的devices表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /devices/getDevicesPublic [get]
func (devicesApi *DevicesApi) GetDevicesPublic(c *gin.Context) {
	// 创建业务用Context
	ctx := c.Request.Context()

	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	devicesService.GetDevicesPublic(ctx)
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的devices表接口信息",
	}, "获取成功", c)
}

func (devicesApi *DevicesApi) GetProductGroups(c *gin.Context) {
	var getProductNameData []productPkg.Products
	global.GVA_DB.Model(&productPkg.Products{}).Find(&getProductNameData)
	response.OkWithDetailed(getProductNameData, "获取成功", c)
}

func (devicesApi *DevicesApi) GetProductTwoGroups(c *gin.Context) {
	var getProductNameData []productPkg.Products
	global.GVA_DB.Model(&productPkg.Products{}).Find(&getProductNameData)
	type names struct {
		Name *string `json:"name"`
	}
	var newNames []*names
	for _, v := range getProductNameData {
		name := names{Name: v.ProductName}
		newNames = append(newNames, &name)
	}
	response.OkWithDetailed(newNames, "获取成功", c)
}

func (devicesApi *DevicesApi) GetDriverGroups(c *gin.Context) {
	var getDriverNameData []driverPkg.Drivers
	global.GVA_DB.Model(&driverPkg.Drivers{}).Find(&getDriverNameData)
	response.OkWithDetailed(getDriverNameData, "获取成功", c)
}
