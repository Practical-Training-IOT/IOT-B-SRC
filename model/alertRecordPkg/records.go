
// 自动生成模板Records
package alertRecordPkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// 告警记录 结构体  Records
type Records struct {
    global.GVA_MODEL
  RuleName  *string `json:"ruleName" form:"ruleName" gorm:"comment:告警规则名称;column:rule_name;"size:50;`  //告警规则名称
  RuleLevel  *string `json:"ruleLevel" form:"ruleLevel" gorm:"comment:告警等级;column:rule_level;"size:50;`  //告警等级
  TriggeringTime  *time.Time `json:"triggeringTime" form:"triggeringTime" gorm:"comment:触发时间;column:triggering_time;"`  //触发时间
  ProcessingTime  *time.Time `json:"processingTime" form:"processingTime" gorm:"comment:处理时间;column:processing_time;"`  //处理时间
  Status  *string `json:"status" form:"status" gorm:"comment:告警状态;column:status;"size:50;`  //告警状态
  ProcessResult  *string `json:"processResult" form:"processResult" gorm:"comment:处理结果;column:process_result;"size:100;`  //处理结果
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName 告警记录 Records自定义表名 records
func (Records) TableName() string {
    return "records"
}





