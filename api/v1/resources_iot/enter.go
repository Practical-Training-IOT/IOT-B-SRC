package resources_iot

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ResourcesApi }

var resourcesService = service.ServiceGroupApp.Resources_iotServiceGroup.ResourcesService
