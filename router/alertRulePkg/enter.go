package alertRulePkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AlarmsRouter }

var alarmsApi = api.ApiGroupApp.AlertRulePkgApiGroup.AlarmsApi
