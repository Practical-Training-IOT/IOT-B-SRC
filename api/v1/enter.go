package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	AlertRecordPkgApiGroup alertRecordPkg.ApiGroup
	AlterLevelPkgApiGroup  alterLevelPkg.ApiGroup
	AlertRulePkgApiGroup   alertRulePkg.ApiGroup
}
