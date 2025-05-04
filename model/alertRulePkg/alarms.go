
// 自动生成模板Alarms
package alertRulePkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 告警规则 结构体  Alarms
type Alarms struct {
    global.GVA_MODEL
  RuleName  *string `json:"ruleName" form:"ruleName" gorm:"comment:规则名称;column:rule_name;"size:50; binding:"required"`  //规则名称
  AlarmType  *string `json:"alarmType" form:"alarmType" gorm:"comment:告警类型;column:alarm_type;"size:50; binding:"required"`  //告警类型
  AlarmLevel  *string `json:"alarmLevel" form:"alarmLevel" gorm:"comment:告警级别;column:alarm_level;"size:50; binding:"required"`  //告警级别
  RuleDescription  *string `json:"ruleDescription" form:"ruleDescription" gorm:"comment:规则描述;column:rule_description;"size:50;type:text;"`  //规则描述
  TriggerMode  *string `json:"triggerMode" form:"triggerMode" gorm:"comment:触发方式;column:trigger_mode;"size:50; binding:"required"`  //触发方式
  ValueType  *string `json:"valueType" form:"valueType" gorm:"comment:取值类型;column:value_type;"size:50; binding:"required"`  //取值类型
  ValuePeriod  *string `json:"valuePeriod" form:"valuePeriod" gorm:"comment:取值周期;column:value_period;"size:50; binding:"required"`  //取值周期
  JudgingCondition  *string `json:"judgingCondition" form:"judgingCondition" gorm:"comment:判断条件;column:judging_condition;"size:50; binding:"required"`  //判断条件
  Value  *string `json:"value" form:"value" gorm:"comment:取值;column:value;"size:50; binding:"required"`  //取值
  SilencePeriod  *string `json:"silencePeriod" form:"silencePeriod" gorm:"comment:静默时间;column:silence_period;"size:100;`  //静默时间
  MeanNotification  *string `json:"meanNotification" form:"meanNotification" gorm:"comment:通知方式;column:mean_notification;"size:50;`  //通知方式
  Status  *string `json:"status" form:"status" gorm:"comment:状态;column:status;"size:50;`  //状态
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 告警规则 Alarms自定义表名 alarms
func (Alarms) TableName() string {
    return "alarms"
}





