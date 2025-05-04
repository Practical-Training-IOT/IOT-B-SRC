package rules

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ RuleInfoApi }

var ruleInfoService = service.ServiceGroupApp.RulesServiceGroup.RuleInfoService
