
package alertRecordPkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
    alertRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg/request"
    "gorm.io/gorm"
)

type RecordsService struct {}
// CreateRecords 创建告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService) CreateRecords(ctx context.Context, records *alertRecordPkg.Records) (err error) {
	err = global.GVA_DB.Create(records).Error
	return err
}

// DeleteRecords 删除告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService)DeleteRecords(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRecordPkg.Records{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&alertRecordPkg.Records{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteRecordsByIds 批量删除告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService)DeleteRecordsByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRecordPkg.Records{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&alertRecordPkg.Records{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateRecords 更新告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService)UpdateRecords(ctx context.Context, records alertRecordPkg.Records) (err error) {
	err = global.GVA_DB.Model(&alertRecordPkg.Records{}).Where("id = ?",records.ID).Updates(&records).Error
	return err
}

// GetRecords 根据ID获取告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService)GetRecords(ctx context.Context, ID string) (records alertRecordPkg.Records, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&records).Error
	return
}
// GetRecordsInfoList 分页获取告警记录记录
// Author [yourname](https://github.com/yourname)
func (recordsService *RecordsService)GetRecordsInfoList(ctx context.Context, info alertRecordPkgReq.RecordsSearch) (list []alertRecordPkg.Records, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&alertRecordPkg.Records{})
    var recordss []alertRecordPkg.Records
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

	err = db.Find(&recordss).Error
	return  recordss, total, err
}
func (recordsService *RecordsService)GetRecordsPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
