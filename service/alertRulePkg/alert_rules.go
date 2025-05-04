
package alertRulePkg

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
    alertRulePkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg/request"
    "gorm.io/gorm"
)

type AlertRulesService struct {}
// CreateAlertRules 创建alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService) CreateAlertRules(ctx context.Context, alertRules *alertRulePkg.AlertRules) (err error) {
	err = global.GVA_DB.Create(alertRules).Error
	return err
}

// DeleteAlertRules 删除alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService)DeleteAlertRules(ctx context.Context, ID string,userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRulePkg.AlertRules{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
              return err
        }
        if err = tx.Delete(&alertRulePkg.AlertRules{},"id = ?",ID).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteAlertRulesByIds 批量删除alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService)DeleteAlertRulesByIds(ctx context.Context, IDs []string,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&alertRulePkg.AlertRules{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", IDs).Delete(&alertRulePkg.AlertRules{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateAlertRules 更新alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService)UpdateAlertRules(ctx context.Context, alertRules alertRulePkg.AlertRules) (err error) {
	err = global.GVA_DB.Model(&alertRulePkg.AlertRules{}).Where("id = ?",alertRules.ID).Updates(&alertRules).Error
	return err
}

// GetAlertRules 根据ID获取alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService)GetAlertRules(ctx context.Context, ID string) (alertRules alertRulePkg.AlertRules, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&alertRules).Error
	return
}
// GetAlertRulesInfoList 分页获取alertRules表记录
// Author [yourname](https://github.com/yourname)
func (alertRulesService *AlertRulesService)GetAlertRulesInfoList(ctx context.Context, info alertRulePkgReq.AlertRulesSearch) (list []alertRulePkg.AlertRules, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&alertRulePkg.AlertRules{})
    var alertRuless []alertRulePkg.AlertRules
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

	err = db.Find(&alertRuless).Error
	return  alertRuless, total, err
}
func (alertRulesService *AlertRulesService)GetAlertRulesPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
