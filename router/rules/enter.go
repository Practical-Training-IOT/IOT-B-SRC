package rules

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ RuleInfoRouter }

var ruleInfoApi = api.ApiGroupApp.RulesApiGroup.RuleInfoApi
