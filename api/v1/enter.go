package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/rules"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	DevicePkgApiGroup  devicePkg.ApiGroup
	DriverPkgApiGroup  driverPkg.ApiGroup
	MyDrivePkgApiGroup myDrivePkg.ApiGroup
	SystemApiGroup         system.ApiGroup
	ExampleApiGroup        example.ApiGroup
	ProductPkgApiGroup     productPkg.ApiGroup
	ResourcesApiGroup      resources.ApiGroup
	Resources_iotApiGroup  resources_iot.ApiGroup
	AlertRecordPkgApiGroup alertRecordPkg.ApiGroup
	AlterLevelPkgApiGroup  alterLevelPkg.ApiGroup
	AlertRulePkgApiGroup   alertRulePkg.ApiGroup
	RulesApiGroup          rules.ApiGroup
}
