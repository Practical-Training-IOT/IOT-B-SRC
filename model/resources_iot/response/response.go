package response

import "encoding/json"

// ProtocolConfig 通用配置结构体
type ProtocolConfig struct {
	InstanceName string          `json:"instance_name" binding:"required"` // 实例名称
	Config       json.RawMessage `json:"config" binding:"required"`        // 协议特定配置 (动态 JSON)
	ProtocolType string          `json:"protocol_type" binding:"required"` // 协议类型 (HTTP, Kafka, MQTT)
}

type HTTPConfigResponse struct {
	InstanceName string       `json:"instanceName" binding:"required"` // 实例名称
	URL          string       `json:"url" binding:"required"`          // URL
	Method       string       `json:"httpMethod" binding:"required"`   // HTTP Method
	BodyType     string       `json:"bodyType" binding:"required"`     // Body Type
	Timeout      int          `json:"timeout" binding:"required"`      // Timeout (ms)
	HttpHeaders  []HttpHeader `json:"httpHeaders" binding:"required"`
}

type HttpHeader struct {
	Key   string `json:"key" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type KafkaConfigResponse struct {
	InstanceName string `json:"instanceName" binding:"required"` // 实例名称
	Brokers      string `json:"brokers" binding:"required"`      // Kafka Brokers
	Topic        string `json:"topic" binding:"required"`        // Topic
	SaslAuthType string `json:"saslAuthType" binding:"required"` // Sasl Auth Type
	SaslUserName string `json:"saslUserName" binding:"required"` // Sasl User Name
	SaslPassword string `json:"saslPassword" binding:"required"` // Sasl Password
}

type MQTTConfigResponse struct {
	InstanceName    string `json:"instanceName" binding:"required"`    // 实例名称
	BrokerAddress   string `json:"brokerAddress" binding:"required"`   // Broker Address
	MQTTTopic       string `json:"mqttTopic" binding:"required"`       // MQTT Topic
	MQTTClient      string `json:"mqttClient" binding:"required"`      // MQTT Client
	ProtocolVersion string `json:"protocolVersion" binding:"required"` // Protocol Version
	QoS             int    `json:"qos" binding:"required"`             // QoS
	Username        string `json:"username" binding:"required"`        // Username
	Password        string `json:"password" binding:"required"`        // Password
}
