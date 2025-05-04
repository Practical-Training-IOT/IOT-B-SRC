package devicePkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ DevicesApi }

var devicesService = service.ServiceGroupApp.DevicePkgServiceGroup.DevicesService
