package alertRulePkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AlertRulesRouter }

var alertRulesApi = api.ApiGroupApp.AlertRulePkgApiGroup.AlertRulesApi
