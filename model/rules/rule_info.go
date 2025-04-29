
// 自动生成模板RuleInfo
package rules
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ruleInfo表 结构体  RuleInfo
type RuleInfo struct {
    global.GVA_MODEL
  RuleId  *int `json:"ruleId" form:"ruleId" gorm:"primarykey;column:rule_id;"size:32;`  //ruleId字段
  RuleName  *string `json:"ruleName" form:"ruleName" gorm:"column:rule_name;"`  //规则名称
  RuleDescription  *string `json:"ruleDescription" form:"ruleDescription" gorm:"column:rule_description;"`  //规则描述
  IsEnabled  *bool `json:"isEnabled" form:"isEnabled" gorm:"column:is_enabled;"`  //是否显示
}


// TableName ruleInfo表 RuleInfo自定义表名 rule_info
func (RuleInfo) TableName() string {
    return "rule_info"
}





