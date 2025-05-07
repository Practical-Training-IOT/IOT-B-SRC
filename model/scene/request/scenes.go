package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ScenesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	Scenename      *string    `json:"scenename" form:"scenename" `
	Enabledstatus  *bool      `json:"enabledstatus" form:"enabledstatus" `
	request.PageInfo
}

type SceneRequest struct {
	ID               int           `gorm:"primaryKey;column:id;not null" json:"id"`
	SceneName        string        `json:"scenename"`
	SceneDescription string        `json:"scenedescription"`
	CreationTime     time.Time     `json:"creationtime"`
	EnabledStatus    bool          `json:"enabledstatus"`
	Device           int           `json:"device"`
	Function         string        `json:"function"`
	JudgeCondition   string        `json:"judgeCondition"`
	JudgeValue       string        `json:"judgeValue"`
	Product          string        `json:"product"`
	TriggerConfig    TriggerConfig `json:"triggerConfig"`
	TriggerMethod    string        `json:"triggerMethod"`
	TriggerTime      string        `json:"triggerTime"`
	TriggerType      string        `json:"triggerType"`
	TriggerWeekdays  []string      `json:"triggerWeekdays"`
	ValuePeriod      string        `json:"valuePeriod"`
	ValueType        string        `json:"valueType"`
	Actions          []ActionItem  `json:"actions"`
	HTTPHeaders      []HTTPHeader  `json:"httpHeaders"`
}

type TriggerConfig struct {
	Method   string   `json:"method"`
	Type     string   `json:"type"`
	Time     string   `json:"time"`
	Weekdays []string `json:"weekdays"`
	Product  string   `json:"product"`
	// 根据你的实际数据补充其他字段
}

type ActionItem struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Product string `json:"product"`
	Device  int    `json:"device"`
}

type HTTPHeader struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	Product  string `json:"product"`
	Device   int    `json:"device"`
	Function string `json:"function"`
}
