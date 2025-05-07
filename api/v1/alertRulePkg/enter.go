package alertRulePkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AlarmsApi }

var alarmsService = service.ServiceGroupApp.AlertRulePkgServiceGroup.AlarmsService
