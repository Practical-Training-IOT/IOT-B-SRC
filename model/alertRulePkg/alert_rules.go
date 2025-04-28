
// 自动生成模板AlertRules
package alertRulePkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// alertRules表 结构体  AlertRules
type AlertRules struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:规则名称;column:name;"size:100; binding:"required"`  //规则名称
  TriggerType  *string `json:"triggerType" form:"triggerType" gorm:"comment:触发类型：设备数据、设备状态等;column:trigger_type;"size:50; binding:"required"`  //触发类型：设备数据、设备状态等
  TriggerCondition  *string `json:"triggerCondition" form:"triggerCondition" gorm:"comment:触发条件，JSON格式;column:trigger_condition;"type:text;" binding:"required"`  //触发条件，JSON格式
  ActionType  *string `json:"actionType" form:"actionType" gorm:"comment:动作类型：通知、设备控制等;column:action_type;"size:50; binding:"required"`  //动作类型：通知、设备控制等
  ActionConfig  *string `json:"actionConfig" form:"actionConfig" gorm:"comment:动作配置，JSON格式;column:action_config;"type:text;" binding:"required"`  //动作配置，JSON格式
  Status  *string `json:"status" form:"status" gorm:"comment:规则状态：启用、禁用;column:status;"size:20; binding:"required"`  //规则状态：启用、禁用
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName alertRules表 AlertRules自定义表名 alert_rules
func (AlertRules) TableName() string {
    return "alert_rules"
}





