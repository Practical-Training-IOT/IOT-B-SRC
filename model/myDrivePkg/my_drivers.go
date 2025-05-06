// 自动生成模板MyDrivers
package myDrivePkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// myDrivers表 结构体  MyDrivers
type MyDrivers struct {
	global.GVA_MODEL
	DriverName           *string `json:"driverName" form:"driverName" gorm:"comment:驱动名称;column:driver_name;"`                               //驱动名称
	DriverType           *string `json:"driverType" form:"driverType" gorm:"comment:驱动类型;column:driver_type;"`                               //驱动类型
	MirrorNumber         *string `json:"mirrorNumber" form:"mirrorNumber" gorm:"comment:镜像编号;column:mirror_number;"`                         //镜像编号
	Agreement            *string `json:"agreement" form:"agreement" gorm:"comment:协议;column:agreement;"`                                     //协议
	MirrorAddress        *string `json:"mirrorAddress" form:"mirrorAddress" gorm:"comment:驱动镜像地址;column:mirror_address;"`                    //驱动镜像地址
	DriverIdentification *string `json:"driverIdentification" form:"driverIdentification" gorm:"comment:驱动标识;column:driver_identification;"` //驱动标识
	Version              *string `json:"version" form:"version" gorm:"comment:版本;column:version;"`                                           //版本
	Description          *string `json:"description" form:"description" gorm:"comment:描述;column:description;"`                               //描述
}

// TableName myDrivers表 MyDrivers自定义表名 myDrivers
func (MyDrivers) TableName() string {
	return "drivers"
}
