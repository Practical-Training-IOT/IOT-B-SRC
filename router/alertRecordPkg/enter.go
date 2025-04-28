package alertRecordPkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AlertsRouter }

var alertsApi = api.ApiGroupApp.AlertRecordPkgApiGroup.AlertsApi
