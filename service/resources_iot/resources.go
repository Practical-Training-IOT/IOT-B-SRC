package resources_iot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
	resources_iotReq "github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/response"
	"log"
	"net/http"
	"strings"
	"time"
)

type ResourcesService struct{}

// CreateResources 创建resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) CreateResources(ctx context.Context, resources *resources_iot.ProtocolConfig) (err error) {
	db := global.GVA_DB.Begin()
	resource := resources_iot.Resources{
		Name:         resources.InstanceName,
		ProtocolType: resources.ProtocolType,
		Status:       "VerificationFailed",
	}
	err = db.Create(&resource).Error
	if err != nil {
		db.Rollback()
		return err
	}
	resCon := resources_iot.ProtocolConfigs{
		Config:       resources.Config,
		ResourcesID:  int32(resource.ID),
		ProtocolType: resources.ProtocolType,
	}
	err = db.Create(&resCon).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return err
}

// DeleteResources 删除resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) DeleteResources(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&resources_iot.Resources{}, "id = ?", ID).Error
	global.GVA_DB.Delete(&resources_iot.ProtocolConfigs{}, "resources_id = ?", ID)
	return err
}

// DeleteResourcesByIds 批量删除resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) DeleteResourcesByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]resources_iot.Resources{}, "id in ?", IDs).Error
	global.GVA_DB.Delete(&resources_iot.ProtocolConfigs{}, "resources_id = ?", IDs)
	return err
}

// UpdateResources 更新resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) UpdateResources(ctx context.Context, resources resources_iot.ProtocolUpdateConfig) (err error) {
	resource := resources_iot.Resources{
		Name:         resources.InstanceName,
		ProtocolType: resources.ProtocolType,
	}
	err = global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", resources.ID).Updates(&resource).Error
	if err != nil {
		return err
	}
	resCon := resources_iot.ProtocolConfigs{
		Config:       resources.Config,
		ResourcesID:  int32(resource.ID),
		ProtocolType: resources.ProtocolType,
	}
	err = global.GVA_DB.Model(&resources_iot.ProtocolConfigs{}).Where("resources_id = ?", resources.ID).Updates(&resCon).Error
	if err != nil {
		return err
	}
	return err
}

// GetResources 根据ID获取resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) GetResources(ctx context.Context, ID string) (interface{}, error) {
	//获取资源
	var resources resources_iot.Resources
	//获取资源配置
	var resCon resources_iot.ProtocolConfigs
	//返回http的数据
	var httpCon response.HTTPConfigResponse
	//返回kafka的数据
	var kafkaCon response.KafkaConfigResponse
	//返回mqtt数据
	var mqttCon response.MQTTConfigResponse
	err := global.GVA_DB.Where("id = ?", ID).First(&resources).Error
	if err != nil {
		return nil, err
	}
	err = global.GVA_DB.Where("resources_id = ?", ID).First(&resCon).Error
	if err != nil {
		return nil, err
	}
	switch resCon.ProtocolType {
	case "HTTP":
		var httpC resources_iot.HTTPConfig
		json.Unmarshal(resCon.Config, &httpC)
		httpCon = response.HTTPConfigResponse{
			InstanceName: resources.Name,
			URL:          httpC.URL,
			Method:       httpC.Method,
			BodyType:     httpC.BodyType,
			Timeout:      httpC.Timeout,
			HttpHeaders:  httpC.HttpHeaders,
		}
		return httpCon, nil
	case "Kafka":
		var kafkaC resources_iot.KafkaConfig
		json.Unmarshal(resCon.Config, &kafkaC)
		kafkaCon = response.KafkaConfigResponse{
			InstanceName: resources.Name,
			Brokers:      kafkaC.Brokers,
			Topic:        kafkaC.Topic,
			SaslAuthType: kafkaC.SaslAuthType,
			SaslUserName: kafkaC.SaslUserName,
			SaslPassword: kafkaC.SaslPassword,
		}
		return kafkaCon, nil
	case "MQTT":
		var mqttC resources_iot.MQTTConfig
		json.Unmarshal(resCon.Config, &mqttC)
		mqttCon = response.MQTTConfigResponse{
			InstanceName:    resources.Name,
			BrokerAddress:   mqttC.BrokerAddress,
			MQTTTopic:       mqttC.MQTTTopic,
			MQTTClient:      mqttC.MQTTClient,
			ProtocolVersion: mqttC.ProtocolVersion,
			QoS:             mqttC.QoS,
			Username:        mqttC.Username,
			Password:        mqttC.Password,
		}
		return mqttCon, nil
	}
	return nil, nil
}

// GetResourcesInfoList 分页获取resources表记录
// Author [yourname](https://github.com/yourname)
func (resourcesService *ResourcesService) GetResourcesInfoList(ctx context.Context, info resources_iotReq.ResourcesSearch) (list []resources_iot.Resources, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&resources_iot.Resources{})
	var resourcess []resources_iot.Resources
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt).Where("protocol_type=?", info.Name)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&resourcess).Error
	return resourcess, total, err
}
func (resourcesService *ResourcesService) GetResourcesPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (resourcesService *ResourcesService) CheckHTTP(ctx context.Context, id int) error {
	//获取资源
	var resources resources_iot.Resources
	//获取资源配置
	var resCon resources_iot.ProtocolConfigs
	//返回http的数据
	var httpCon response.HTTPConfigResponse
	err := global.GVA_DB.Where("id = ?", id).First(&resources).Error
	if err != nil {
		fmt.Println(err)
	}
	err = global.GVA_DB.Where("resources_id = ?", id).First(&resCon).Error
	if err != nil {
		fmt.Println(err)
	}
	var httpC resources_iot.HTTPConfig
	json.Unmarshal(resCon.Config, &httpC)
	httpCon = response.HTTPConfigResponse{
		InstanceName: resources.Name,
		URL:          httpC.URL,
		Method:       httpC.Method,
		BodyType:     httpC.BodyType,
		Timeout:      httpC.Timeout,
		HttpHeaders:  httpC.HttpHeaders,
	}
	request, err := sendHTTPRequest(httpCon)
	if err != nil {
		global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationFailed")
		return err
	} else if request {
		global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationPassed")
	}
	return nil
}

func sendHTTPRequest(config response.HTTPConfigResponse) (bool, error) {
	// 创建 HTTP 客户端，设置超时时间
	client := &http.Client{
		Timeout: time.Duration(config.Timeout) * time.Second, // 超时时间从配置中读取
	}

	// 构建请求
	req, err := http.NewRequest(config.Method, config.URL, nil)
	if err != nil {
		return false, fmt.Errorf("failed to create request: %v", err)
	}

	// 如果 HttpHeaders 不为空，则添加到请求中
	if config.HttpHeaders != nil {
		for _, value := range config.HttpHeaders {
			req.Header.Set(value.Key, value.Value)
		}
	}

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return false, fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	// 检查响应状态码
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return true, nil
	}

	return false, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
}

func (resourcesService *ResourcesService) CheckKafka(ctx context.Context, id int) error {
	//获取资源
	var resources resources_iot.Resources
	//获取资源配置
	var resCon resources_iot.ProtocolConfigs
	//返回kafka的数据
	var kafkaCon response.KafkaConfigResponse
	err := global.GVA_DB.Where("id = ?", id).First(&resources).Error
	if err != nil {
		fmt.Println(err)
	}
	err = global.GVA_DB.Where("resources_id = ?", id).First(&resCon).Error
	if err != nil {
		fmt.Println(err)
	}
	var kafkaC resources_iot.KafkaConfig
	json.Unmarshal(resCon.Config, &kafkaC)
	kafkaCon = response.KafkaConfigResponse{
		InstanceName: resources.Name,
		Brokers:      kafkaC.Brokers,
		Topic:        kafkaC.Topic,
		SaslAuthType: kafkaC.SaslAuthType,
		SaslUserName: kafkaC.SaslUserName,
		SaslPassword: kafkaC.SaslPassword,
	}
	err = validateKafka(kafkaCon)
	if err != nil {
		global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationFailed")
		return err
	}
	global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationPassed")
	return nil
}

func validateKafka(config response.KafkaConfigResponse) error {
	// 分割 Kafka brokers 字符串
	kafkaBrokers := strings.Split(config.Brokers, ",")
	if config.InstanceName == "" {
		return fmt.Errorf("instance name is required")
	}
	if len(kafkaBrokers) == 0 || kafkaBrokers[0] == "" {
		return fmt.Errorf("kafka brokers are required")
	}
	if config.Topic == "" {
		return fmt.Errorf("topic is required")
	}

	// 创建 Sarama 配置
	saramaConfig := sarama.NewConfig() // 使用不同的变量名避免冲突
	if config.SaslAuthType != "" {
		switch config.SaslAuthType {
		case "plain":
			saramaConfig.Net.SASL.Enable = true
			saramaConfig.Net.SASL.User = config.SaslUserName
			saramaConfig.Net.SASL.Password = config.SaslPassword
		default:
			log.Printf("Unsupported SASL auth type: %s", config.SaslAuthType)
		}
	}

	// 创建 Kafka 客户端
	client, err := sarama.NewClient(kafkaBrokers, saramaConfig)
	if err != nil {
		return fmt.Errorf("failed to create kafka client: %v", err)
	}
	defer client.Close()

	// 获取主题分区信息
	partitions, err := client.Partitions(config.Topic)
	if err != nil {
		return fmt.Errorf("failed to get partitions for topic %s: %v", config.Topic, err)
	}

	// 打印成功信息
	fmt.Printf("Successfully connected to Kafka instance '%s' with brokers '%v' and topic '%s'. Partitions: %v\n",
		config.InstanceName, kafkaBrokers, config.Topic, partitions)

	return nil
}

func (resourcesService *ResourcesService) CheckMQTT(ctx context.Context, id int) error {
	//获取资源
	var resources resources_iot.Resources
	//获取资源配置
	var resCon resources_iot.ProtocolConfigs
	//返回mqtt数据
	var mqttCon response.MQTTConfigResponse
	err := global.GVA_DB.Where("id = ?", id).First(&resources).Error
	if err != nil {
		fmt.Println(err)
	}
	err = global.GVA_DB.Where("resources_id = ?", id).First(&resCon).Error
	if err != nil {
		fmt.Println(err)
	}
	var mqttC resources_iot.MQTTConfig
	json.Unmarshal(resCon.Config, &mqttC)
	mqttCon = response.MQTTConfigResponse{
		InstanceName:    resources.Name,
		BrokerAddress:   mqttC.BrokerAddress,
		MQTTTopic:       mqttC.MQTTTopic,
		MQTTClient:      mqttC.MQTTClient,
		ProtocolVersion: mqttC.ProtocolVersion,
		QoS:             mqttC.QoS,
		Username:        mqttC.Username,
		Password:        mqttC.Password,
	}
	err = validateMQTT(mqttCon)
	if err != nil {
		global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationFailed")
		return err
	}
	global.GVA_DB.Model(&resources_iot.Resources{}).Where("id = ?", id).Update("status", "VerificationPassed")
	return nil
}

func validateMQTT(config response.MQTTConfigResponse) error {
	if config.InstanceName == "" {
		return fmt.Errorf("instance name is required")
	}

	if config.BrokerAddress == "" {
		return fmt.Errorf("mqtt broker address is required")
	}

	if config.MQTTTopic == "" {
		return fmt.Errorf("mqtt topic is required")
	}

	if config.MQTTClient == "" {
		return fmt.Errorf("mqtt client ID is required")
	}

	if config.ProtocolVersion == "" {
		config.ProtocolVersion = "3.1.1" // 默认协议版本为 3.1.1
	}

	if config.QoS < 0 || config.QoS > 2 {
		return fmt.Errorf("invalid QoS value, must be 0, 1, or 2")
	}

	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", config.BrokerAddress))
	opts.SetClientID(config.MQTTClient)
	opts.SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
	})

	if config.Username != "" && config.Password != "" {
		opts.SetUsername(config.Username)
		opts.SetPassword(config.Password)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to connect to mqtt broker: %v", token.Error())
	}
	defer client.Disconnect(250)

	if token := client.Subscribe(config.MQTTTopic, byte(config.QoS), nil); token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %v", config.MQTTTopic, token.Error())
	}

	fmt.Printf("Successfully connected to MQTT instance '%s' with broker '%s', topic '%s', and client ID '%s'.\n",
		config.InstanceName, config.BrokerAddress, config.MQTTTopic, config.MQTTClient)

	return nil
}
