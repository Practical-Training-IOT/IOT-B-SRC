package driverPkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/driverPkg"
	driverPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/driverPkg/request"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type DriversService struct{}

// CreateDrivers 创建drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) CreateDrivers(ctx context.Context, drivers *driverPkg.Drivers) (err error) {
	newMirrorNumber := "Mirror" + uuid.NewString()
	drivers.MirrorNumber = &newMirrorNumber
	err = global.GVA_DB.Create(drivers).Error
	return err
}

// DeleteDrivers 删除drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) DeleteDrivers(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&driverPkg.Drivers{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&driverPkg.Drivers{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDriversByIds 批量删除drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) DeleteDriversByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&driverPkg.Drivers{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&driverPkg.Drivers{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateDrivers 更新drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) UpdateDrivers(ctx context.Context, drivers driverPkg.Drivers) (err error) {
	err = global.GVA_DB.Model(&driverPkg.Drivers{}).Where("id = ?", drivers.ID).Updates(&drivers).Error
	return err
}

// GetDrivers 根据ID获取drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) GetDrivers(ctx context.Context, ID string) (drivers driverPkg.Drivers, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&drivers).Error
	return
}

// GetDriversInfoList 分页获取drivers表记录
// Author [yourname](https://github.com/yourname)
func (driversService *DriversService) GetDriversInfoList(ctx context.Context, info driverPkgReq.DriversSearch) (list []driverPkg.Drivers, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&driverPkg.Drivers{})
	var driverss []driverPkg.Drivers
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.DriverName != nil && *info.DriverName != "" {
		db = db.Where("driver_name = ?", *info.DriverName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Where("driver_type = ?", "1").Find(&driverss).Error
	return driverss, total, err
}
func (driversService *DriversService) GetDriversPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
