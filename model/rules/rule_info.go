// 自动生成模板RuleInfo
package rules

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"gorm.io/gorm"
	"time"
)

// ruleInfo表 结构体  RuleInfo
type RuleInfo struct {
	RuleId          int            `json:"ruleId" form:"ruleId" gorm:"primarykey;column:rule_id;size:32;"`         //ruleId字段
	RuleName        string         `json:"ruleName" form:"ruleName" gorm:"column:rule_name;"`                      //规则名称
	RuleDescription string         `json:"ruleDescription" form:"ruleDescription" gorm:"column:rule_description;"` //规则描述
	IsEnabled       bool           `json:"isEnabled" form:"isEnabled" gorm:"column:is_enabled;"`                   //是否显示
	CreatedAt       time.Time      // 创建时间
	UpdatedAt       time.Time      // 更新时间
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

type Rule struct {
	global.GVA_MODEL
	RuleName        string `json:"ruleName,omitempty"`
	RuleDescription string `json:"ruleDescription,omitempty"`
	IsEnabled       bool   `json:"isEnabled"`
	MessageSource   string `json:"messageSource,omitempty"`
	QueryFields     string `json:"queryFields,omitempty"`
	Conditions      string `json:"conditions,omitempty"`
	SQLStatement    string `json:"sqlStatement,omitempty"`
	ForwardMethod   string `json:"forwardMethod,omitempty"`
	Resource        int    `json:"resource,omitempty"`
}

type RuleResponse struct {
	global.GVA_MODEL
	RuleName        string `json:"ruleName,omitempty"`
	RuleDescription string `json:"ruleDescription,omitempty"`
	IsEnabled       bool   `json:"isEnabled"`
	MessageSource   string `json:"messageSource,omitempty"`
	QueryFields     string `json:"queryFields,omitempty"`
	Conditions      string `json:"conditions,omitempty"`
	SQLStatement    string `json:"sqlStatement,omitempty"`
	ForwardMethod   string `json:"forwardMethod,omitempty"`
	Resource        string `json:"resource,omitempty"`
}

// TableName ruleInfo表 RuleInfo自定义表名 rule_info
func (RuleInfo) TableName() string {
	return "rule_info"
}

// RuleForwarding 映射到 "public"."rule_forwarding" 表
type RuleForwarding struct {
	ForwardingID   int        `gorm:"column:forwarding_id;primaryKey;autoIncrement" json:"forwarding_id"`
	RuleID         int        `gorm:"column:rule_id" json:"rule_id"` // 使用指针表示允许NULL
	ForwardingType string     `gorm:"column:forwarding_type;size:50" json:"forwarding_type"`
	UseResource    string     `gorm:"column:use_resource;size:255" json:"use_resource"`
	CreatedAt      time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName 返回表名，确保 GORM 知道该结构体对应哪个数据库表
func (RuleForwarding) TableName() string {
	return "public.rule_forwarding"
}

// RuleCondition 映射到 "public"."rule_condition" 表
type RuleCondition struct {
	ConditionID   int        `gorm:"column:condition_id;primaryKey;autoIncrement" json:"condition_id"`
	RuleID        int        `gorm:"column:rule_id" json:"rule_id"` // 外键，可为空
	MessageSource string     `gorm:"column:message_source;size:255" json:"message_source"`
	QueryFields   string     `gorm:"column:query_fields;type:text" json:"query_fields"`
	Conditions    string     `gorm:"column:conditions;type:text" json:"conditions"`
	SqlStatement  string     `gorm:"column:sql_statement;type:text" json:"sql_statement"`
	CreatedAt     time.Time  `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName 返回表名，带 schema 名称（public.rule_condition）
func (RuleCondition) TableName() string {
	return "public.rule_condition"
}

// Message 是对应 message_bus 表的结构体
type Message struct {
	ID          int64           `gorm:"primaryKey" json:"id"` // BIGSERIAL 主键
	DeviceID    string          `gorm:"column:device_id;not null" json:"device_id"`
	MessageType string          `gorm:"column:message_type;not null" json:"message_type"`
	Content     json.RawMessage `gorm:"column:content;type:jsonb;not null" json:"content"`
	ReceivedAt  time.Time       `gorm:"column:received_at;default:CURRENT_TIMESTAMP" json:"received_at"`
	Processed   bool            `gorm:"column:processed;default:false" json:"processed"`
	ProcessedAt *time.Time      `gorm:"column:processed_at" json:"processed_at,omitempty"` // 可为 nil
}

// TableName 显式指定数据库表名
func (Message) TableName() string {
	return "message"
}
