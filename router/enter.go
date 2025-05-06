package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/ai"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/router/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/router/rules"
	"github.com/flipped-aurora/gin-vue-admin/server/router/scene"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	DevicePkg      devicePkg.RouterGroup
	DriverPkg      driverPkg.RouterGroup
	MyDrivePkg     myDrivePkg.RouterGroup
	System         system.RouterGroup
	Example        example.RouterGroup
	ProductPkg     productPkg.RouterGroup
	Resources      resources.RouterGroup
	Resources_iot  resources_iot.RouterGroup
	AlertRecordPkg alertRecordPkg.RouterGroup
	AlterLevelPkg  alterLevelPkg.RouterGroup
	AlertRulePkg   alertRulePkg.RouterGroup
	Rules          rules.RouterGroup
	Scene          scene.RouterGroup
	AI             ai.RouterGroup
}
