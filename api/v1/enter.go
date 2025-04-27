package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resources"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup        system.ApiGroup
	ExampleApiGroup       example.ApiGroup
	ProductPkgApiGroup    productPkg.ApiGroup
	ResourcesApiGroup     resources.ApiGroup
	Resources_iotApiGroup resources_iot.ApiGroup
}
