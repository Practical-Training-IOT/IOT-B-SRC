
package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ScenesSearch struct{
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    Scenename  *string `json:"scenename" form:"scenename" `
    Enabledstatus  *bool `json:"enabledstatus" form:"enabledstatus" `
    request.PageInfo
}
