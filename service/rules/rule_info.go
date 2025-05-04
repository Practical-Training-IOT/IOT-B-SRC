package rules

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rules"
	rulesReq "github.com/flipped-aurora/gin-vue-admin/server/model/rules/request"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type RuleInfoService struct{}

// Rule 结构体用于处理HTTP方式
type Rule struct {
	RuleName        string                `json:"ruleName"`
	RuleDescription string                `json:"ruleDescription"`
	URL             string                `json:"url"`
	HTTPMethod      string                `json:"httpMethod"`
	BodyType        string                `json:"bodyType"`
	Headers         []response.HttpHeader `json:"headers"`
	BodyContent     string                `json:"bodyContent"`
	TimeoutMS       int                   `json:"timeoutMS"`
}

// CreateRuleInfo 创建ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) CreateRuleInfo(ctx context.Context, ruleInfo *rules.Rule) (err error) {
	db := global.GVA_DB.Begin()
	sqlStatement := "SELECT" + ruleInfo.QueryFields + "FROM" + ruleInfo.MessageSource + "WHERE" + ruleInfo.Conditions
	info := rules.RuleInfo{
		RuleName:        ruleInfo.RuleName,
		RuleDescription: ruleInfo.RuleDescription,
		IsEnabled:       ruleInfo.IsEnabled,
	}
	err = db.Create(&info).Error
	if err != nil {
		db.Rollback()
		return err
	}
	condition := rules.RuleCondition{
		RuleID:        info.RuleId,
		MessageSource: ruleInfo.MessageSource,
		QueryFields:   ruleInfo.QueryFields,
		Conditions:    ruleInfo.Conditions,
		SqlStatement:  sqlStatement,
	}
	err = db.Create(&condition).Error
	if err != nil {
		db.Rollback()
		return err
	}
	forwarding := rules.RuleForwarding{
		RuleID:         info.RuleId,
		ForwardingType: ruleInfo.ForwardMethod,
		UseResource:    strconv.Itoa(ruleInfo.Resource),
		CreatedAt:      time.Now(),
	}
	err = db.Create(&forwarding).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()

	var messages []rules.Message
	result := db.Raw(sqlStatement).Scan(&messages)
	if result.Error != nil {
		fmt.Printf("执行 SQL 查询失败: %v", result.Error)
	}
	switch ruleInfo.ForwardMethod {
	case "HTTP":
		//获取资源
		var resources resources_iot.Resources
		//获取资源配置
		var resCon resources_iot.ProtocolConfigs
		//返回http的数据
		err = global.GVA_DB.Where("id = ?", ruleInfo.Resource).First(&resources).Error
		if err != nil {
			fmt.Println(err)
		}
		err = global.GVA_DB.Where("resources_id = ?", ruleInfo.Resource).First(&resCon).Error
		if err != nil {
			fmt.Println(err)
		}
		var httpC resources_iot.HTTPConfig
		json.Unmarshal(resCon.Config, &httpC)
		rule := Rule{
			RuleName:        ruleInfo.RuleName,
			RuleDescription: ruleInfo.RuleDescription,
			URL:             httpC.URL,
			HTTPMethod:      httpC.Method,
			BodyType:        httpC.BodyType,
			Headers:         httpC.HttpHeaders,
			TimeoutMS:       httpC.Timeout,
		}
		err = sendHTTPRequest(rule, messages)
		if err != nil {
			return err
		}
	case "Kafka":
		//获取资源
		var resources resources_iot.Resources
		//获取资源配置
		var resCon resources_iot.ProtocolConfigs
		//返回kafka的数据
		err = global.GVA_DB.Where("id = ?", ruleInfo.Resource).First(&resources).Error
		if err != nil {
			fmt.Println(err)
		}
		err = global.GVA_DB.Where("resources_id = ?", ruleInfo.Resource).First(&resCon).Error
		if err != nil {
			fmt.Println(err)
		}
		var kafkaC resources_iot.KafkaConfig
		json.Unmarshal(resCon.Config, &kafkaC)
		var sli []string
		sli = append(sli, kafkaC.Brokers)
		err = sendToKafka(sli, kafkaC.Topic, kafkaC.SaslUserName, kafkaC.SaslPassword, messages)
		if err != nil {
			return err
		}
	case "MQTT":
		//获取资源
		var resources resources_iot.Resources
		//获取资源配置
		var resCon resources_iot.ProtocolConfigs
		//返回mqtt数据
		err = global.GVA_DB.Where("id = ?", ruleInfo.Resource).First(&resources).Error
		if err != nil {
			fmt.Println(err)
		}
		err = global.GVA_DB.Where("resources_id = ?", ruleInfo.Resource).First(&resCon).Error
		if err != nil {
			fmt.Println(err)
		}
		var mqttC resources_iot.MQTTConfig
		json.Unmarshal(resCon.Config, &mqttC)
		err = sendToMQTT(mqttC.BrokerAddress, mqttC.MQTTClient, mqttC.Username, mqttC.Password, byte(mqttC.QoS), messages)
		if err != nil {
			return err
		}
	default:
		//获取资源
		var resources resources_iot.Resources
		//获取资源配置
		var resCon resources_iot.ProtocolConfigs
		//返回http的数据
		err = global.GVA_DB.Where("id = ?", ruleInfo.Resource).First(&resources).Error
		if err != nil {
			fmt.Println(err)
		}
		err = global.GVA_DB.Where("resources_id = ?", ruleInfo.Resource).First(&resCon).Error
		if err != nil {
			fmt.Println(err)
		}
		var httpC resources_iot.HTTPConfig
		json.Unmarshal(resCon.Config, &httpC)
		rule := Rule{
			RuleName:        ruleInfo.RuleName,
			RuleDescription: ruleInfo.RuleDescription,
			URL:             httpC.URL,
			HTTPMethod:      httpC.Method,
			BodyType:        httpC.BodyType,
			Headers:         httpC.HttpHeaders,
			TimeoutMS:       httpC.Timeout,
		}
		err = sendHTTPRequest(rule, messages)
		if err != nil {
			return err
		}
	}
	return err
}

func sendHTTPRequest(rule Rule, messages []rules.Message) error {
	client := &http.Client{
		Timeout: time.Duration(rule.TimeoutMS) * time.Second,
	}

	var bodyReader *bytes.Reader

	switch rule.BodyType {
	case "json":
		bodyBytes, err := json.Marshal(messages)
		if err != nil {
			return fmt.Errorf("构建 JSON 请求体失败: %v", err)
		}
		bodyReader = bytes.NewReader(bodyBytes)
	case "text", "html", "xml", "js":
		bodyStr := fmt.Sprintf("%v", messages)
		bodyReader = bytes.NewReader([]byte(bodyStr))
	case "form":
		formData := url.Values{}

		for i, msg := range messages {
			// 构建表单数据
			formData.Add(fmt.Sprintf("messages[%d].id", i), fmt.Sprintf("%d", msg.ID))
			formData.Add(fmt.Sprintf("messages[%d].device_id", i), msg.DeviceID)
			formData.Add(fmt.Sprintf("messages[%d].message_type", i), msg.MessageType)

			// 将 Content 转换为字符串形式添加到表单中
			formData.Add(fmt.Sprintf("messages[%d].content", i), string(msg.Content))

			// 格式化时间
			formData.Add(fmt.Sprintf("messages[%d].received_at", i), msg.ReceivedAt.Format(time.RFC3339))

			// 添加 Processed 字段
			formData.Add(fmt.Sprintf("messages[%d].processed", i), fmt.Sprintf("%t", msg.Processed))

			// 处理 ProcessedAt 可能为空的情况
			if msg.ProcessedAt != nil {
				formData.Add(fmt.Sprintf("messages[%d].processed_at", i), msg.ProcessedAt.Format(time.RFC3339))
			} else {
				formData.Add(fmt.Sprintf("messages[%d].processed_at", i), "")
			}
		}

		bodyReader = bytes.NewReader([]byte(formData.Encode()))
	case "none":
		bodyReader = bytes.NewReader([]byte{})
	default:
		return fmt.Errorf("不支持的 Body 类型: %s", rule.BodyType)
	}

	req, err := http.NewRequest(rule.HTTPMethod, rule.URL, bodyReader)
	if err != nil {
		return fmt.Errorf("创建 HTTP 请求失败: %v", err)
	}

	// 设置请求头
	for _, value := range rule.Headers {
		req.Header.Set(value.Key, value.Value)
	}

	// 设置 Content-Type
	switch rule.BodyType {
	case "json":
		req.Header.Set("Content-Type", "application/json")
	case "text":
		req.Header.Set("Content-Type", "text/plain")
	case "html":
		req.Header.Set("Content-Type", "text/html")
	case "xml":
		req.Header.Set("Content-Type", "application/xml")
	case "js":
		req.Header.Set("Content-Type", "application/javascript")
	case "form":
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("发送 HTTP 请求失败: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取响应体失败: %v", err)
	}

	fmt.Printf("响应状态码: %s\n", resp.Status)
	fmt.Printf("响应内容: %s\n", string(body))

	return nil
}

func sendToKafka(brokers []string, topic, username, password string, messages []rules.Message) error {
	config := sarama.NewConfig()
	config.Net.SASL.Enable = true
	config.Net.SASL.User = username
	config.Net.SASL.Password = password
	config.Net.SASL.Mechanism = sarama.SASLTypePlaintext
	config.Producer.Return.Successes = true

	client, err := sarama.NewClient(brokers, config)
	if err != nil {
		return fmt.Errorf("连接 Kafka 失败: %v", err)
	}
	defer client.Close()

	producer, err := sarama.NewSyncProducerFromClient(client)
	if err != nil {
		return fmt.Errorf("创建生产者失败: %v", err)
	}
	defer producer.Close()

	for _, msg := range messages {
		jsonData, _ := json.Marshal(msg)
		message := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.ByteEncoder(jsonData),
		}

		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			return fmt.Errorf("发送消息失败: %v", err)
		}
		fmt.Printf("消息已发送至 Kafka 分区 %d 偏移量 %d\n", partition, offset)
	}

	return nil
}

func sendToMQTT(broker, clientID, username, password string, qos byte, messages []rules.Message) error {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return fmt.Errorf("连接 MQTT 失败: %v", token.Error())
	}
	defer client.Disconnect(250)

	topic := "iot/messages"

	for _, msg := range messages {
		jsonData, _ := json.Marshal(msg)
		token := client.Publish(topic, qos, false, jsonData)
		token.Wait()
		if token.Error() != nil {
			return fmt.Errorf("发布消息失败: %v", token.Error())
		}
		fmt.Printf("消息已发布到 MQTT 主题: %s\n", topic)
	}

	return nil
}

// DeleteRuleInfo 删除ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) DeleteRuleInfo(ctx context.Context, ID string) (err error) {
	err = global.GVA_DB.Delete(&rules.RuleInfo{}, "rule_id = ?", ID).Error
	err = global.GVA_DB.Delete(&rules.RuleCondition{}, "rule_id = ?", ID).Error
	err = global.GVA_DB.Delete(&rules.RuleForwarding{}, "rule_id = ?", ID).Error
	return err
}

// DeleteRuleInfoByIds 批量删除ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) DeleteRuleInfoByIds(ctx context.Context, IDs []string) (err error) {
	err = global.GVA_DB.Delete(&[]rules.RuleInfo{}, "rule_id in ?", IDs).Error
	err = global.GVA_DB.Delete(&rules.RuleCondition{}, "rule_id = ?", IDs).Error
	err = global.GVA_DB.Delete(&rules.RuleForwarding{}, "rule_id = ?", IDs).Error
	return err
}

// UpdateRuleInfo 更新ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) UpdateRuleInfo(ctx context.Context, ruleInfo rules.Rule) (err error) {
	db := global.GVA_DB.Begin()
	sqlStatement := "SELECT" + ruleInfo.QueryFields + "FROM" + ruleInfo.MessageSource + "WHERE" + ruleInfo.Conditions
	info := rules.RuleInfo{
		RuleName:        ruleInfo.RuleName,
		RuleDescription: ruleInfo.RuleDescription,
		IsEnabled:       ruleInfo.IsEnabled,
	}
	err = db.Model(&rules.RuleInfo{}).Where("rule_id = ?", ruleInfo.ID).Updates(&info).Error
	if err != nil {
		db.Rollback()
		return err
	}
	condition := rules.RuleCondition{
		RuleID:        int(ruleInfo.ID),
		MessageSource: ruleInfo.MessageSource,
		QueryFields:   ruleInfo.QueryFields,
		Conditions:    ruleInfo.Conditions,
		SqlStatement:  sqlStatement,
	}
	err = db.Model(&rules.RuleCondition{}).Where("rule_id = ?", ruleInfo.ID).Updates(&condition).Error
	if err != nil {
		db.Rollback()
		return err
	}
	forwarding := rules.RuleForwarding{
		RuleID:         int(ruleInfo.ID),
		ForwardingType: ruleInfo.ForwardMethod,
		UseResource:    strconv.Itoa(ruleInfo.Resource),
		CreatedAt:      time.Now(),
	}
	err = db.Model(&rules.RuleForwarding{}).Where("rule_id = ?", ruleInfo.ID).Updates(&forwarding).Error
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return err
}

// GetRuleInfo 根据ID获取ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) GetRuleInfo(ctx context.Context, ID string) (rules.RuleResponse, error) {
	var ruleInfo rules.RuleInfo
	var ruleForwarding rules.RuleForwarding
	var ruleCondition rules.RuleCondition
	err := global.GVA_DB.Where("rule_id = ?", ID).First(&ruleInfo).Error
	if err != nil {
		return rules.RuleResponse{}, err
	}
	err = global.GVA_DB.Where("rule_id = ?", ID).First(&ruleForwarding).Error
	if err != nil {
		return rules.RuleResponse{}, err
	}
	err = global.GVA_DB.Where("rule_id = ?", ID).First(&ruleCondition).Error
	if err != nil {
		return rules.RuleResponse{}, err
	}
	var res resources_iot.Resources
	global.GVA_DB.Where("id = ?", ruleForwarding.UseResource).First(&res)
	sli := rules.RuleResponse{
		GVA_MODEL: global.GVA_MODEL{
			ID:        uint(ruleInfo.RuleId),
			CreatedAt: ruleInfo.CreatedAt,
			UpdatedAt: ruleInfo.UpdatedAt,
			DeletedAt: ruleInfo.DeletedAt,
		},
		RuleName:        ruleInfo.RuleName,
		RuleDescription: ruleInfo.RuleDescription,
		IsEnabled:       ruleInfo.IsEnabled,
		MessageSource:   ruleCondition.MessageSource,
		QueryFields:     ruleCondition.QueryFields,
		Conditions:      ruleCondition.Conditions,
		SQLStatement:    ruleCondition.SqlStatement,
		ForwardMethod:   ruleForwarding.ForwardingType,
		Resource:        res.Name,
	}
	return sli, nil
}

// GetRuleInfoInfoList 分页获取ruleInfo表记录
// Author [yourname](https://github.com/yourname)
func (ruleInfoService *RuleInfoService) GetRuleInfoInfoList(ctx context.Context, info rulesReq.RuleInfoSearch) (list []rules.Rule, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&rules.RuleInfo{})
	var ruleInfos []rules.RuleInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.RuleName != nil && *info.RuleName != "" {
		db = db.Where("rule_name LIKE ?", "%"+*info.RuleName+"%")
	}
	if info.IsEnabled != nil {
		db = db.Where("is_enabled = ?", *info.IsEnabled)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&ruleInfos).Error

	var sli []rules.Rule

	for i := 0; i < len(ruleInfos); i++ {
		var ruleCondition rules.RuleForwarding
		global.GVA_DB.Where("rule_id = ?", ruleInfos[i].RuleId).Find(&ruleCondition)
		one := rules.Rule{
			GVA_MODEL: global.GVA_MODEL{
				ID:        uint(ruleInfos[i].RuleId),
				CreatedAt: ruleInfos[i].CreatedAt,
				UpdatedAt: time.Time{},
				DeletedAt: gorm.DeletedAt{},
			},
			RuleName:        ruleInfos[i].RuleName,
			RuleDescription: ruleInfos[i].RuleDescription,
			IsEnabled:       ruleInfos[i].IsEnabled,
			ForwardMethod:   ruleCondition.ForwardingType,
		}
		sli = append(sli, one)
	}
	return sli, total, err
}
func (ruleInfoService *RuleInfoService) GetRuleInfoPublic(ctx context.Context) {
	// 此方法为获取数据源定义的数据
	// 请自行实现
}

func (ruleInfoService *RuleInfoService) HandleSwitchChange(c context.Context, info rulesReq.HandSearch) error {
	err := global.GVA_DB.Model(&rules.RuleInfo{}).Where("rule_id = ?", info.ID).Update("is_enabled", info.Status).Error
	if err != nil {
		return err
	}
	return nil
}
