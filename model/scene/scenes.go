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

// TriggerCondition 表示数据库中 "trigger_condition" 表对应的结构体。
type TriggerCondition struct {
	ID int `gorm:"primaryKey;column:id" json:"id"`
	// SceneID 是触发条件关联的场景ID，允许为空（可选）。
	SceneID int `gorm:"column:sceneid" json:"sceneid"`
	// TriggerType 表示触发类型，例如设备触发、时间触发等，不能为空。
	TriggerType string `gorm:"column:triggertype;size:50;not null" json:"triggertype"`
	// Product 是触发条件涉及的产品信息，允许为空（可选）。
	Product string `gorm:"column:product;size:255" json:"product"`
	// Device 是触发条件涉及的设备信息，允许为空（可选）。
	Device string `gorm:"column:device;size:255" json:"device"`
	// Function 是触发条件涉及的功能信息，允许为空（可选）。
	Function string `gorm:"column:function;size:255" json:"function"`
	// ValueType 描述值类型的信息，允许为空（可选）。
	ValueType string `gorm:"column:valuetype;size:50" json:"valuetype"`
	// JudgmentCondition 存储判断条件，允许为空（可选）。
	JudgmentCondition string `gorm:"column:judgmentcondition;size:255" json:"judgmentcondition"`
	// Time 用于存储与触发条件相关的时间信息，允许为空（可选）。
	Time *time.Time `gorm:"column:time;size:6" json:"time"`
	// DaysOfWeek 存储一周中的哪几天作为触发条件，允许为空（可选）。
	DaysOfWeek string `gorm:"column:daysofweek;size:50" json:"daysofweek"`
	// TriggerMode 定义了触发模式，例如定时触发、事件触发等，允许为空（可选）。
	TriggerMode string `gorm:"column:triggermode;size:50" json:"triggermode"`
}

// TableName 设置表名为 "trigger_condition".
func (TriggerCondition) TableName() string {
	return "triggerconditions"
}

// ExecutionAction 表示数据库中 "executionactions" 表对应的结构体。
type ExecutionAction struct {
	ID int `gorm:"primaryKey;column:id;not null" json:"id"`
	// SceneID 是关联的场景ID，允许为空（可选）。
	SceneID int `gorm:"column:sceneid" json:"sceneid"`
	// ActionOrder 定义动作的执行顺序，不能为空。
	ActionOrder int `gorm:"column:actionorder;not null" json:"actionorder"`
	// ActionType 表示动作类型（如DeviceAction），不能为空。
	ActionType string `gorm:"column:actiontype;size:50;not null" json:"actiontype"`
	// Product 表示产品名称，允许为空（可选）。
	Product string `gorm:"column:product;size:255" json:"product"`
	// Device 表示设备名称，允许为空（可选）。
	Device string `gorm:"column:device;size:255" json:"device"`
	// Function 表示功能名称，允许为空（可选）。
	Function string `gorm:"column:function;size:255" json:"function"`
	// Value 表示数值，允许为空（可选）。
	Value string `gorm:"column:value;size:255" json:"value"`
}

// TableName 设置表名为 "executionactions".
func (ExecutionAction) TableName() string {
	return "executionactions"
}
