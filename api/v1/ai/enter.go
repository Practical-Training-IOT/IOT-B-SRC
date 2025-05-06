package ai

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ScenesApi }

var scenesService = service.ServiceGroupApp.SceneServiceGroup.ScenesService
