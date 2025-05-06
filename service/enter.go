package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/service/rules"
	"github.com/flipped-aurora/gin-vue-admin/server/service/scene"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	DevicePkgServiceGroup      devicePkg.ServiceGroup
	DriverPkgServiceGroup      driverPkg.ServiceGroup
	MyDrivePkgServiceGroup     myDrivePkg.ServiceGroup
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	ProductPkgServiceGroup     productPkg.ServiceGroup
	ResourcesServiceGroup      resources.ServiceGroup
	Resources_iotServiceGroup  resources_iot.ServiceGroup
	AlertRecordPkgServiceGroup alertRecordPkg.ServiceGroup
	AlterLevelPkgServiceGroup  alterLevelPkg.ServiceGroup
	AlertRulePkgServiceGroup   alertRulePkg.ServiceGroup
	RulesServiceGroup          rules.ServiceGroup
	SceneServiceGroup          scene.ServiceGroup
}
