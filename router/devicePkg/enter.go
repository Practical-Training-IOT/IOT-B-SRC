package devicePkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ DevicesRouter }

var devicesApi = api.ApiGroupApp.DevicePkgApiGroup.DevicesApi
