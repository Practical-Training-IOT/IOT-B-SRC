package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type DevicesSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	ProductRange   *string    `json:"product_range" form:"product_range" `
	Platform       *string    `json:"platform" form:"platform" `
	Drive          *string    `json:"drive" form:"drive" `
	Status         *string    `json:"status" form:"status" `
	DriveName      *string    `json:"drive_name" form:"drive_name" `
	ProductName    *string    `json:"product_name" form:"product_name" `
	request.PageInfo
}
