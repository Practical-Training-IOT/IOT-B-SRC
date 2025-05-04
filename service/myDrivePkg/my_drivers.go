
package myDrivePkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg"
    myDrivePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg/request"
)

type MyDriversService struct {}
// CreateMyDrivers 创建myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService) CreateMyDrivers(ctx context.Context, myDrivers *myDrivePkg.MyDrivers) (err error) {
	err = global.GVA_DB.Create(myDrivers).Error
	return err
}

// DeleteMyDrivers 删除myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService)DeleteMyDrivers(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&myDrivePkg.MyDrivers{},"id = ?",ID).Error
	return err
}

// DeleteMyDriversByIds 批量删除myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService)DeleteMyDriversByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]myDrivePkg.MyDrivers{},"id in ?",IDs).Error
	return err
}

// UpdateMyDrivers 更新myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService)UpdateMyDrivers(ctx context.Context, myDrivers myDrivePkg.MyDrivers) (err error) {
	err = global.GVA_DB.Model(&myDrivePkg.MyDrivers{}).Where("id = ?",myDrivers.ID).Updates(&myDrivers).Error
	return err
}

// GetMyDrivers 根据ID获取myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService)GetMyDrivers(ctx context.Context, ID string) (myDrivers myDrivePkg.MyDrivers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&myDrivers).Error
	return
}
// GetMyDriversInfoList 分页获取myDrivers表记录
// Author [yourname](https://github.com/yourname)
func (myDriversService *MyDriversService)GetMyDriversInfoList(ctx context.Context, info myDrivePkgReq.MyDriversSearch) (list []myDrivePkg.MyDrivers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&myDrivePkg.MyDrivers{})
    var myDriverss []myDrivePkg.MyDrivers
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
    if info.DriverName != nil && *info.DriverName != "" {
        db = db.Where("driver_name = ?", *info.DriverName)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&myDriverss).Error
	return  myDriverss, total, err
}
func (myDriversService *MyDriversService)GetMyDriversPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
