
// 自动生成模板Resources
package resources_iot
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// resources表 结构体  Resources
type Resources struct {
    global.GVA_MODEL
  Name  *string `json:"name" form:"name" gorm:"comment:资源名称;column:name;"size:100; binding:"required"`  //资源名称
  ProtocolType  *string `json:"protocolType" form:"protocolType" gorm:"comment:协议类型;column:protocol_type;"size:20;`  //协议类型
  Status  *string `json:"status" form:"status" gorm:"comment:资源状态;column:status;"size:20;`  //资源状态
}


// TableName resources表 Resources自定义表名 resources
func (Resources) TableName() string {
    return "resources"
}





