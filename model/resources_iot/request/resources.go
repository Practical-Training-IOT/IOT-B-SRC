package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ResourcesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

type ResourcesSearchRes struct {
	ID   int    `json:"id" form:"id" query:"id"`
	Name string `json:"name" form:"name" query:"name"`
}
