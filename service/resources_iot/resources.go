
package resources_iot

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
    resources_iotReq "github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/request"
)

type ResourcesService struct {}
// CreateResources 创建resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) CreateResources(ctx context.Context, resources *resources_iot.Resources) (err error) {
	err = global.GVA_DB.Create(resources).Error
	return err
}

// DeleteResources 删除resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService)DeleteResources(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&resources_iot.Resources{},"id = ?",ID).Error
	return err
}

// DeleteResourcesByIds 批量删除resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService)DeleteResourcesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]resources_iot.Resources{},"id in ?",IDs).Error
	return err
}

// UpdateResources 更新resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService)UpdateResources(ctx context.Context, resources resources_iot.Resources) (err error) {
	err = global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?",resources.ID).Updates(&resources).Error
	return err
}

// GetResources 根据ID获取resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService)GetResources(ctx context.Context, ID string) (resources resources_iot.Resources, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&resources).Error
	return
}
// GetResourcesInfoList 分页获取resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService)GetResourcesInfoList(ctx context.Context, info resources_iotReq.ResourcesSearch) (list []resources_iot.Resources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&resources_iot.Resources{})
    var resourcess []resources_iot.Resources
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&resourcess).Error
	return  resourcess, total, err
}
func (resourcesService *ResourcesService)GetResourcesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
