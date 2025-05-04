
// 自动生成模板Devices
package devicePkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// devices表 结构体  Devices
type Devices struct {
    global.GVA_MODEL
  ProductRange  *string `json:"product_range" form:"product_range" gorm:"comment:产品范围;column:product_range;"`  //产品范围
  Platform  *string `json:"platform" form:"platform" gorm:"comment:平台;column:platform;"`  //平台
  Drive  *string `json:"drive" form:"drive" gorm:"comment:驱动;column:drive;"`  //驱动
  Status  *string `json:"status" form:"status" gorm:"comment:状态;column:status;"`  //状态
  AddDeviceMethod  *string `json:"add_device_method" form:"add_device_method" gorm:"default:1;comment:设备添加方式;column:add_device_method;" binding:"required"`  //设备添加方式
  DriveName  *string `json:"drive_name" form:"drive_name" gorm:"comment:设备名称;column:drive_name;" binding:"required"`  //设备名称
  BelongingProducts  *string `json:"belonging_products" form:"belonging_products" gorm:"comment:所属产品;column:belonging_products;" binding:"required"`  //所属产品
  AssociationDriven  *string `json:"association_driven" form:"association_driven" gorm:"column:association_driven;"`  //关联驱动
  DeviceDescription  *string `json:"device_description" form:"device_description" gorm:"comment:设备描述;column:device_description;"type:text;"`  //设备描述
    CreatedBy  uint   `gorm:"column:created_by;comment:创建者"`
    UpdatedBy  uint   `gorm:"column:updated_by;comment:更新者"`
    DeletedBy  uint   `gorm:"column:deleted_by;comment:删除者"`
}


// TableName devices表 Devices自定义表名 devices
func (Devices) TableName() string {
    return "devices"
}





