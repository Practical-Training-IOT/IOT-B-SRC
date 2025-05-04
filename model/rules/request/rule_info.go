package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type RuleInfoSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	RuleName       *string    `json:"ruleName" form:"ruleName" `
	IsEnabled      *bool      `json:"isEnabled" form:"isEnabled" `
	request.PageInfo
}

type HandSearch struct {
	ID     int  `json:"id" form:"id" uri:"id"`
	Status bool `json:"status" form:"status"`
}
