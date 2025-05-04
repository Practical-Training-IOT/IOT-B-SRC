package myDrivePkg

import (
	
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg"
    myDrivePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg/request"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type MyDriversApi struct {}



// CreateMyDrivers 创建myDrivers表
// @Tags MyDrivers
// @Summary 创建myDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body myDrivePkg.MyDrivers true "创建myDrivers表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /myDrivers/createMyDrivers [post]
func (myDriversApi *MyDriversApi) CreateMyDrivers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var myDrivers myDrivePkg.MyDrivers
	err := c.ShouldBindJSON(&myDrivers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = myDriversService.CreateMyDrivers(ctx,&myDrivers)
	if err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败:" + err.Error(), c)
		return
	}
    response.OkWithMessage("创建成功", c)
}

// DeleteMyDrivers 删除myDrivers表
// @Tags MyDrivers
// @Summary 删除myDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body myDrivePkg.MyDrivers true "删除myDrivers表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /myDrivers/deleteMyDrivers [delete]
func (myDriversApi *MyDriversApi) DeleteMyDrivers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	err := myDriversService.DeleteMyDrivers(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// DeleteMyDriversByIds 批量删除myDrivers表
// @Tags MyDrivers
// @Summary 批量删除myDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /myDrivers/deleteMyDriversByIds [delete]
func (myDriversApi *MyDriversApi) DeleteMyDriversByIds(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	IDs := c.QueryArray("IDs[]")
	err := myDriversService.DeleteMyDriversByIds(ctx,IDs)
	if err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("批量删除成功", c)
}

// UpdateMyDrivers 更新myDrivers表
// @Tags MyDrivers
// @Summary 更新myDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data body myDrivePkg.MyDrivers true "更新myDrivers表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /myDrivers/updateMyDrivers [put]
func (myDriversApi *MyDriversApi) UpdateMyDrivers(c *gin.Context) {
    // 从ctx获取标准context进行业务行为
    ctx := c.Request.Context()

	var myDrivers myDrivePkg.MyDrivers
	err := c.ShouldBindJSON(&myDrivers)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = myDriversService.UpdateMyDrivers(ctx,myDrivers)
	if err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:" + err.Error(), c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// FindMyDrivers 用id查询myDrivers表
// @Tags MyDrivers
// @Summary 用id查询myDrivers表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param ID query uint true "用id查询myDrivers表"
// @Success 200 {object} response.Response{data=myDrivePkg.MyDrivers,msg=string} "查询成功"
// @Router /myDrivers/findMyDrivers [get]
func (myDriversApi *MyDriversApi) FindMyDrivers(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	ID := c.Query("ID")
	remyDrivers, err := myDriversService.GetMyDrivers(ctx,ID)
	if err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败:" + err.Error(), c)
		return
	}
	response.OkWithData(remyDrivers, c)
}
// GetMyDriversList 分页获取myDrivers表列表
// @Tags MyDrivers
// @Summary 分页获取myDrivers表列表
// @Security ApiKeyAuth
// @Accept application/json
// @Produce application/json
// @Param data query myDrivePkgReq.MyDriversSearch true "分页获取myDrivers表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /myDrivers/getMyDriversList [get]
func (myDriversApi *MyDriversApi) GetMyDriversList(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

	var pageInfo myDrivePkgReq.MyDriversSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := myDriversService.GetMyDriversInfoList(ctx,pageInfo)
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

// GetMyDriversPublic 不需要鉴权的myDrivers表接口
// @Tags MyDrivers
// @Summary 不需要鉴权的myDrivers表接口
// @Accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /myDrivers/getMyDriversPublic [get]
func (myDriversApi *MyDriversApi) GetMyDriversPublic(c *gin.Context) {
    // 创建业务用Context
    ctx := c.Request.Context()

    // 此接口不需要鉴权
    // 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
    myDriversService.GetMyDriversPublic(ctx)
    response.OkWithDetailed(gin.H{
       "info": "不需要鉴权的myDrivers表接口信息",
    }, "获取成功", c)
}
