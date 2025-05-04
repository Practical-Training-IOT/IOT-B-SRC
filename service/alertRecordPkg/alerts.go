
package alertRecordPkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
    alertRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg/request"
    "gorm.io/gorm"
)

type AlertsService struct {}
// CreateAlerts 创建alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService) CreateAlerts(ctx context.Context, alerts *alertRecordPkg.Alerts) (err error) {
	err = global.GVA_DB.Create(alerts).Error
	return err
}

// DeleteAlerts 删除alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService)DeleteAlerts(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRecordPkg.Alerts{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&alertRecordPkg.Alerts{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAlertsByIds 批量删除alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService)DeleteAlertsByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRecordPkg.Alerts{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&alertRecordPkg.Alerts{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAlerts 更新alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService)UpdateAlerts(ctx context.Context, alerts alertRecordPkg.Alerts) (err error) {
	err = global.GVA_DB.Model(&alertRecordPkg.Alerts{}).Where("id = ?",alerts.ID).Updates(&alerts).Error
	return err
}

// GetAlerts 根据ID获取alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService)GetAlerts(ctx context.Context, ID string) (alerts alertRecordPkg.Alerts, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alerts).Error
	return
}
// GetAlertsInfoList 分页获取alerts表记录
// Author [yourname](https://github.com/yourname)
func (alertsService *AlertsService)GetAlertsInfoList(ctx context.Context, info alertRecordPkgReq.AlertsSearch) (list []alertRecordPkg.Alerts, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&alertRecordPkg.Alerts{})
    var alertss []alertRecordPkg.Alerts
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

	err = db.Find(&alertss).Error
	return  alertss, total, err
}
func (alertsService *AlertsService)GetAlertsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
