package resources_iot

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ResourcesRouter }

var resourcesApi = api.ApiGroupApp.Resources_iotApiGroup.ResourcesApi
