package alertRecordPkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ RecordsRouter }

var recordsApi = api.ApiGroupApp.AlertRecordPkgApiGroup.RecordsApi
