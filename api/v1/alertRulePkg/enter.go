package alertRulePkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AlertRulesApi }

var alertRulesService = service.ServiceGroupApp.AlertRulePkgServiceGroup.AlertRulesService
