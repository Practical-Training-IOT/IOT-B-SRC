package devicePkg

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg"
	devicePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg/request"
	"gorm.io/gorm"
)

type DevicesService struct{}

// CreateDevices 创建devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) CreateDevices(ctx context.Context, devices *devicePkg.Devices) (err error) {
	newStatus := "2"
	newPlatform := "cloud"
	devices.Status = &newStatus
	devices.Platform = &newPlatform
	err = global.GVA_DB.Create(devices).Error
	return err
}

// DeleteDevices 删除devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) DeleteDevices(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&devicePkg.Devices{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&devicePkg.Devices{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteDevicesByIds 批量删除devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) DeleteDevicesByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&devicePkg.Devices{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&devicePkg.Devices{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateDevices 更新devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) UpdateDevices(ctx context.Context, devices devicePkg.Devices) (err error) {
	err = global.GVA_DB.Model(&devicePkg.Devices{}).Where("id = ?", devices.ID).Updates(&devices).Error
	return err
}

// GetDevices 根据ID获取devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) GetDevices(ctx context.Context, ID string) (devices devicePkg.Devices, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&devices).Error
	return
}

// GetDevicesInfoList 分页获取devices表记录
// Author [yourname](https://github.com/yourname)
func (devicesService *DevicesService) GetDevicesInfoList(ctx context.Context, info devicePkgReq.DevicesSearch) (list []devicePkg.Devices, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&devicePkg.Devices{})
	var devicess []devicePkg.Devices
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	fmt.Println(info.ProductName)
	if info.ProductName != nil && *info.ProductName != "" {
		db = db.Where("belonging_products = ?", *info.ProductName)
	}

	if info.ProductRange != nil && *info.ProductRange != "" {
		db = db.Where("product_range = ?", *info.ProductRange)
	}
	if info.Platform != nil && *info.Platform != "" {
		db = db.Where("platform = ?", *info.Platform)
	}
	if info.Drive != nil && *info.Drive != "" {
		db = db.Where("drive = ?", *info.Drive)
	}
	if info.Status != nil && *info.Status != "" {
		db = db.Where("status = ?", *info.Status)
	}
	if info.DriveName != nil && *info.DriveName != "" {
		db = db.Where("drive_name = ?", *info.DriveName)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&devicess).Error
	return devicess, total, err
}
func (devicesService *DevicesService) GetDevicesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}
