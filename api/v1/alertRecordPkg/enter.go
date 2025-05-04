package alertRecordPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ AlertsApi }

var alertsService = service.ServiceGroupApp.AlertRecordPkgServiceGroup.AlertsService
