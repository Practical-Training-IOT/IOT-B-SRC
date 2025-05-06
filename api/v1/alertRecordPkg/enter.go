package alertRecordPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ RecordsApi }

var recordsService = service.ServiceGroupApp.AlertRecordPkgServiceGroup.RecordsService
