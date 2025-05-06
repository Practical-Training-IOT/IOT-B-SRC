// 自动生成模板Scenes
package scene

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// scenes表 结构体  Scenes
type Scenes struct {
	global.GVA_MODEL
	Scenename        *string    `json:"scenename" form:"scenename" gorm:"comment:场景的名称。;column:scenename;"size:255; binding:"required"`     //场景的名称。
	Scenedescription *string    `json:"scenedescription" form:"scenedescription" gorm:"comment:场景的描述。;column:scenedescription;"type:text;"` //场景的描述。
	Creationtime     *time.Time `json:"creationtime" form:"creationtime" gorm:"comment:场景创建的时间。;column:creationtime;"size:6;`               //场景创建的时间。
	Enabledstatus    *bool      `json:"enabledstatus" form:"enabledstatus" gorm:"comment:场景是否启用的状态。;column:enabledstatus;"`                 //场景是否启用的状态。
}

// TableName scenes表 Scenes自定义表名 scenes
func (Scenes) TableName() string {
	return "scenes"
}

type Category struct {
	ID           string    `gorm:"column:id;type:varchar(20);primary_key" json:"id"`
	CategoryName string    `gorm:"column:category_name;type:varchar(100);not null" json:"category_name"`
	CategoryKey  string    `gorm:"column:category_key;type:varchar(50);unique;not null" json:"category_key"`
	Scene        string    `gorm:"column:scene;type:varchar(50);not null" json:"scene"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName 指定数据库中的表名
func (Category) TableName() string {
	return "categories"
}

type Property struct {
	ID          string   `gorm:"column:id;type:varchar(50);primary_key" json:"id"`
	CategoryKey string   `gorm:"column:category_key;type:varchar(50);not null" json:"category_key"`
	Name        string   `gorm:"column:name;type:varchar(100);not null" json:"name"`
	Code        string   `gorm:"column:code;type:varchar(50);not null" json:"code"`
	AccessMode  string   `gorm:"column:access_mode;type:varchar(10);not null" json:"access_mode"`
	DataType    string   `gorm:"column:data_type;type:varchar(20);not null" json:"data_type"`
	Unit        *string  `gorm:"column:unit;type:varchar(20)" json:"unit,omitempty"`      // 可为空
	MinValue    *float64 `gorm:"column:min_value;type:float8" json:"min_value,omitempty"` // 可为空
	MaxValue    *float64 `gorm:"column:max_value;type:float8" json:"max_value,omitempty"` // 可为空
}

// TableName 指定数据库中的表名
func (Property) TableName() string {
	return "properties"
}
