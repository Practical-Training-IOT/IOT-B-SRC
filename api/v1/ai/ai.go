package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"net/http"
	"time"
)

type ScenesApi struct{}

type Request struct {
	ID      int64  `json:"id"`
	Message string `json:"message"`
}

type Ai struct {
	ID       int64  `gorm:"primaryKey" json:"id"`                // 主键id
	Model    string `gorm:"column:model;size:50" json:"model"`   // 使用的哪个ai
	Req      string `gorm:"column:req;type:text" json:"req"`     // 发送的请求
	Res      string `gorm:"column:res;type:text" json:"res"`     // 回答的话语
	Title    string `gorm:"column:title;type:text" json:"title"` // 标题
	AiScenId int64  `gorm:"column:ai_scen_id" json:"ai_scen_id"`
}

// TableName 指定数据库表名
func (Ai) TableName() string {
	return "ai"
}

// AiScene 对应数据库表 public.ai_scene
type AiScene struct {
	ID        int64     `gorm:"primaryKey" json:"id"`                // 主键id
	Title     string    `gorm:"column:title;type:text" json:"title"` // 场景标题
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"` // 创建时间
}

// TableName 指定数据库表名
func (AiScene) TableName() string {
	return "ai_scene"
}

type OllamaResponse struct {
	Model              string `json:"model"`
	CreatedAt          string `json:"created_at"`
	Response           string `json:"response"`
	Done               bool   `json:"done"`
	DoneReason         string `json:"done_reason,omitempty"`
	Context            []int  `json:"context,omitempty"`
	TotalDuration      int64  `json:"total_duration,omitempty"`
	LoadDuration       int64  `json:"load_duration,omitempty"`
	PromptEvalCount    int    `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int    `json:"prompt_eval_duration,omitempty"`
	EvalCount          int    `json:"eval_count,omitempty"`
	EvalDuration       int64  `json:"eval_duration,omitempty"`
}

type ChatResponse struct {
	Message string `json:"message"`
	Model   string `json:"model"`
	ID      int64  `json:"id"`
}

type OneHistoryRequest struct {
	ID int64 `json:"id"`
}

type ChatOneResponse struct {
	Message string `json:"message"`
	Req     string `json:"req"`
	Model   string `json:"model"`
	ID      int64  `json:"id"`
}

type OneHistoryResponse struct {
	Chat []ChatOneResponse
	ID   int64 `json:"id"`
}

var Model = "deepseek-r1"

func (a ScenesApi) Chats(c *gin.Context) {
	var scenes Request
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println(Model)
	// 构造请求体
	requestBody := map[string]string{
		"model":  Model,
		"prompt": scenes.Message,
	}

	jsonBody, _ := json.Marshal(requestBody)

	// 发起请求
	resp, err := http.Post("http://117.27.231.112:11434/api/generate", "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	var fullResponse string
	decoder := json.NewDecoder(resp.Body)

	for {
		var r OllamaResponse
		if err := decoder.Decode(&r); err == io.EOF {
			break
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		fullResponse += r.Response
		fmt.Println(r.Response)
		if r.Done {
			break
		}
	}

	db := global.GVA_DB.Begin()

	var aiScene AiScene

	var aiSceneId int64

	if scenes.ID == 0 {
		aiScene = AiScene{
			Title:     scenes.Message,
			CreatedAt: time.Now(),
		}
		err = db.Create(&aiScene).Error
		if err != nil {
			db.Rollback()
			global.GVA_LOG.Error("获取失败!", zap.Error(err))
			response.FailWithMessage("获取失败:"+err.Error(), c)
			return
		}
		aiSceneId = aiScene.ID
	} else {
		aiSceneId = scenes.ID
	}

	ai := Ai{
		Model:    Model,
		Req:      scenes.Message,
		Res:      fullResponse,
		Title:    scenes.Message,
		AiScenId: aiSceneId,
	}

	err = global.GVA_DB.Create(&ai).Error

	if err != nil {
		db.Rollback()
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}

	db.Commit()

	chatResponse := ChatResponse{
		Message: fullResponse,
		Model:   Model,
		ID:      aiSceneId,
	}

	// 返回响应给前端
	response.OkWithDetailed(chatResponse, "获取成功", c)
}

func (a ScenesApi) Change(c *gin.Context) {
	var scenes Request
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	Model = scenes.Message
}

func (a ScenesApi) History(c *gin.Context) {
	var ai []AiScene
	err := global.GVA_DB.Find(&ai).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	// 返回响应给前端
	response.OkWithDetailed(ai, "获取成功", c)
}

func (a ScenesApi) OneHistory(c *gin.Context) {
	var scenes OneHistoryRequest
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var aiScen AiScene
	err = global.GVA_DB.Where("id = ?", scenes.ID).Find(&aiScen).Error
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败:"+err.Error(), c)
		return
	}
	var ai []Ai
	global.GVA_DB.Where("ai_scen_id = ?", scenes.ID).Find(&ai)
	var sli []ChatOneResponse
	for _, v := range ai {
		one := ChatOneResponse{
			Message: v.Res,
			Model:   v.Model,
			Req:     v.Req,
		}
		sli = append(sli, one)
	}
	one := OneHistoryResponse{
		Chat: sli,
		ID:   aiScen.ID,
	}
	response.OkWithDetailed(one, "获取成功", c)
}
