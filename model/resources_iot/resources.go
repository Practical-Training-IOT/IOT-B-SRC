// 自动生成模板Resources
package resources_iot

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/response"
)

// resources表 结构体  Resources
type Resources struct {
	global.GVA_MODEL
	Name         string `json:"name" form:"name" gorm:"comment:资源名称;column:name;"size:100; binding:"required"`      //资源名称
	ProtocolType string `json:"protocolType" form:"protocolType" gorm:"comment:协议类型;column:protocol_type;"size:20;` //协议类型
	Status       string `json:"status" form:"status" gorm:"comment:资源状态;column:status;"size:20;`                    //资源状态
}

// TableName resources表 Resources自定义表名 resources
func (Resources) TableName() string {
	return "resources"
}

// ProtocolConfigs 对应 protocol_configs 表的结构体
type ProtocolConfigs struct {
	ID           int64  `gorm:"primaryKey;column:id;default:nextval('protocol_configs_id_seq'::regclass)" json:"id"`
	Config       []byte `gorm:"column:config;type:jsonb" json:"config"`
	ResourcesID  int32  `gorm:"column:resources_id" json:"resources_id"`
	ProtocolType string `gorm:"column:protocol_type;type:varchar(20)" json:"protocol_type"`
}

// TableName 指定表名
func (ProtocolConfigs) TableName() string {
	return "protocol_configs"
}

// ProtocolConfig 通用配置结构体
type ProtocolConfig struct {
	InstanceName string          `json:"instance_name" binding:"required"` // 实例名称
	Config       json.RawMessage `json:"config" binding:"required"`        // 协议特定配置 (动态 JSON)
	ProtocolType string          `json:"protocol_type" binding:"required"` // 协议类型 (HTTP, Kafka, MQTT)
}

// ProtocolConfig 更新通用配置结构体
type ProtocolUpdateConfig struct {
	ID           int64           `json:"id" binding:"required"`
	InstanceName string          `json:"instance_name" binding:"required"` // 实例名称
	Config       json.RawMessage `json:"config" binding:"required"`        // 协议特定配置 (动态 JSON)
	ProtocolType string          `json:"protocol_type" binding:"required"` // 协议类型 (HTTP, Kafka, MQTT)
}

type HTTPConfig struct {
	URL         string                `json:"url" binding:"required"`         // URL
	Method      string                `json:"http_method" binding:"required"` // HTTP Method
	BodyType    string                `json:"body_type" binding:"required"`   // Body Type
	Timeout     int                   `json:"timeout" binding:"required"`     // Timeout (ms)
	HttpHeaders []response.HttpHeader `json:"httpHeaders" binding:"required"`
}

type KafkaConfig struct {
	Brokers      string `json:"brokers" binding:"required"`        // Kafka Brokers
	Topic        string `json:"topic" binding:"required"`          // Topic
	SaslAuthType string `json:"sasl_auth_type" binding:"required"` // Sasl Auth Type
	SaslUserName string `json:"sasl_user_name" binding:"required"` // Sasl User Name
	SaslPassword string `json:"sasl_password" binding:"required"`  // Sasl Password
}

type MQTTConfig struct {
	BrokerAddress   string `json:"broker_address" binding:"required"`   // Broker Address
	MQTTTopic       string `json:"mqtt_topic" binding:"required"`       // MQTT Topic
	MQTTClient      string `json:"mqtt_client" binding:"required"`      // MQTT Client
	ProtocolVersion string `json:"protocol_version" binding:"required"` // Protocol Version
	QoS             int    `json:"qos" binding:"required"`              // QoS
	Username        string `json:"username" binding:"required"`         // Username
	Password        string `json:"password" binding:"required"`         // Password
}
