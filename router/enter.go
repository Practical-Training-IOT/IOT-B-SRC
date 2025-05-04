package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	ProductPkg productPkg.RouterGroup
	DevicePkg  devicePkg.RouterGroup
	DriverPkg  driverPkg.RouterGroup
	MyDrivePkg myDrivePkg.RouterGroup
}
