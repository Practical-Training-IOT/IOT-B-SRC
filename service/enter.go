package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/service/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup        system.ServiceGroup
	ExampleServiceGroup       example.ServiceGroup
	ProductPkgServiceGroup    productPkg.ServiceGroup
	ResourcesServiceGroup     resources.ServiceGroup
	Resources_iotServiceGroup resources_iot.ServiceGroup
}
