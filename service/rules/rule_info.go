
package rules

import (
	"context"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rules"
    rulesReq "github.com/flipped-aurora/gin-vue-admin/server/model/rules/request"
)

type RuleInfoService struct {}
// CreateRuleInfo 创建ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) CreateRuleInfo(ctx context.Context, ruleInfo *rules.RuleInfo) (err error) {
	err = global.GVA_DB.Create(ruleInfo).Error
	return err
}

// DeleteRuleInfo 删除ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService)DeleteRuleInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&rules.RuleInfo{},"id = ?",ID).Error
	return err
}

// DeleteRuleInfoByIds 批量删除ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService)DeleteRuleInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]rules.RuleInfo{},"id in ?",IDs).Error
	return err
}

// UpdateRuleInfo 更新ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService)UpdateRuleInfo(ctx context.Context, ruleInfo rules.RuleInfo) (err error) {
	err = global.GVA_DB.Model(&rules.RuleInfo{}).Where("id = ?",ruleInfo.ID).Updates(&ruleInfo).Error
	return err
}

// GetRuleInfo 根据ID获取ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService)GetRuleInfo(ctx context.Context, ID string) (ruleInfo rules.RuleInfo, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&ruleInfo).Error
	return
}
// GetRuleInfoInfoList 分页获取ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService)GetRuleInfoInfoList(ctx context.Context, info rulesReq.RuleInfoSearch) (list []rules.RuleInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&rules.RuleInfo{})
    var ruleInfos []rules.RuleInfo
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    
    if info.RuleName != nil && *info.RuleName != "" {
        db = db.Where("rule_name LIKE ?", "%"+ *info.RuleName+"%")
    }
    if info.IsEnabled != nil {
        db = db.Where("is_enabled = ?", *info.IsEnabled)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	if limit != 0 {
       db = db.Limit(limit).Offset(offset)
    }

	err = db.Find(&ruleInfos).Error
	return  ruleInfos, total, err
}
func (ruleInfoService *RuleInfoService)GetRuleInfoPublic(ctx context.Context) {
    // 此方法为获取数据源定义的数据
    // 请自行实现
}
