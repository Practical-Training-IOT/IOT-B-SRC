package alertRulePkg

import (
	"context"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
	alertRulePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	"gorm.io/gorm"
)

type AlarmsService struct{}

// CreateAlarms 创建告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) CreateAlarms(ctx context.Context, alarms *alertRulePkg.Alarms) (err error) {
	err = global.GVA_DB.Create(alarms).Error
	return err
}

// DeleteAlarms 删除告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) DeleteAlarms(ctx context.Context, ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alertRulePkg.Alarms{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&alertRulePkg.Alarms{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteAlarmsByIds 批量删除告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) DeleteAlarmsByIds(ctx context.Context, IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&alertRulePkg.Alarms{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&alertRulePkg.Alarms{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateAlarms 更新告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) UpdateAlarms(ctx context.Context, alarms alertRulePkg.Alarms) (err error) {
	err = global.GVA_DB.Model(&alertRulePkg.Alarms{}).Where("id = ?", alarms.ID).Updates(&alarms).Error
	return err
}

// GetAlarms 根据ID获取告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) GetAlarms(ctx context.Context, ID string) (alarms alertRulePkg.Alarms, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alarms).Error
	return
}

// GetAlarmsInfoList 分页获取告警规则记录
// Author [yourname](https://github.com/yourname)
func (alarmsService *AlarmsService) GetAlarmsInfoList(ctx context.Context, info alertRulePkgReq.AlarmsSearch) (list []alertRulePkg.Alarms, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&alertRulePkg.Alarms{})
	var alarmss []alertRulePkg.Alarms
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&alarmss).Error
	return alarmss, total, err
}
func (alarmsService *AlarmsService) GetAlarmsPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (alarmsService *AlarmsService) GetProductList(ctx context.Context) (list []productPkg.Products, err error) {
	db := global.GVA_DB.Model(&productPkg.Products{})
	var product []productPkg.Products
	err = db.Find(&product).Error
	fmt.Println(product)
	return product, err

}

func (alarmsService *AlarmsService) GetEquipment(ctx context.Context, ID string) (alarms alertRulePkg.Alarms, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alarms).Error
	return
}
