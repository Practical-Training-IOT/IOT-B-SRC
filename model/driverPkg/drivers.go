
// 自动生成模板Drivers
package driverPkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// drivers表 结构体  Drivers
type Drivers struct {
    global.GVA_MODEL
  DriverName  *string `json:"driver_name" form:"driver_name" gorm:"comment:驱动名称;column:driver_name;" binding:"required"`  //驱动名称
  DriverType  *string `json:"driver_type" form:"driver_type" gorm:"comment:驱动类型;column:driver_type;"`  //驱动类型
  MirrorNumber  *string `json:"mirror_number" form:"mirror_number" gorm:"comment:镜像编号;column:mirror_number;"`  //镜像编号
  Agreement  *string `json:"agreement" form:"agreement" gorm:"comment:协议;column:agreement;" binding:"required"`  //协议
  MirrorAddress  *string `json:"mirror_address" form:"mirror_address" gorm:"comment:驱动镜像地址;column:mirror_address;" binding:"required"`  //镜像仓库地址
  DriverIdentification  *string `json:"driver_identification" form:"driver_identification" gorm:"comment:驱动标识;column:driver_identification;" binding:"required"`  //驱动标识
  Version  *string `json:"version" form:"version" gorm:"comment:版本;column:version;" binding:"required"`  //版本
  Description  *string `json:"description" form:"description" gorm:"comment:描述;column:description;"type:text;"`  //描述
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName drivers表 Drivers自定义表名 drivers
func (Drivers) TableName() string {
    return "drivers"
}





