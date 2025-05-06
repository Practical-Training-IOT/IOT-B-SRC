package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type ScenesApi struct{}

type Request struct {
	Message string `json:"message"`
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

func (a ScenesApi) Chats(c *gin.Context) {
	var scenes Request
	err := c.ShouldBindJSON(&scenes)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 构造请求体
	requestBody := map[string]string{
		"model":  "deepseek-r1",
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

	// 返回响应给前端
	response.OkWithDetailed(fullResponse, "获取成功", c)
}
