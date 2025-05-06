package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup         system.ServiceGroup
	ExampleServiceGroup        example.ServiceGroup
	ProductPkgServiceGroup     productPkg.ServiceGroup
	AlertRecordPkgServiceGroup alertRecordPkg.ServiceGroup
	AlterLevelPkgServiceGroup  alterLevelPkg.ServiceGroup
	AlertRulePkgServiceGroup   alertRulePkg.ServiceGroup
}
