
// 自动生成模板Alerts
package alertRecordPkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// alerts表 结构体  Alerts
type Alerts struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:告警名称;column:name;"size:100; binding:"required"`  //告警名称
  DeviceId  *int `json:"deviceId" form:"deviceId" gorm:"comment:设备ID;column:device_id;"size:32;`  //设备ID
  AlertLevel  *string `json:"alertLevel" form:"alertLevel" gorm:"comment:告警级别;column:alert_level;"size:20; binding:"required"`  //告警级别
  AlertType  *string `json:"alertType" form:"alertType" gorm:"comment:告警类型;column:alert_type;"size:50; binding:"required"`  //告警类型
  AlertContent  *string `json:"alertContent" form:"alertContent" gorm:"comment:告警内容;column:alert_content;"`  //告警内容
  Status  *string `json:"status" form:"status" gorm:"comment:告警状态;column:status;"size:20; binding:"required"`  //告警状态
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName alerts表 Alerts自定义表名 alerts
func (Alerts) TableName() string {
    return "alerts"
}





