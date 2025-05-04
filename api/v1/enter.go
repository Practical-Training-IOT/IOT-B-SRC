package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	ProductPkgApiGroup productPkg.ApiGroup
	DevicePkgApiGroup  devicePkg.ApiGroup
	DriverPkgApiGroup  driverPkg.ApiGroup
	MyDrivePkgApiGroup myDrivePkg.ApiGroup
}
