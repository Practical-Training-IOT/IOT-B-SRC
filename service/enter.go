package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup     system.ServiceGroup
	ExampleServiceGroup    example.ServiceGroup
	ProductPkgServiceGroup productPkg.ServiceGroup
	DevicePkgServiceGroup  devicePkg.ServiceGroup
	DriverPkgServiceGroup  driverPkg.ServiceGroup
	MyDrivePkgServiceGroup myDrivePkg.ServiceGroup
}
