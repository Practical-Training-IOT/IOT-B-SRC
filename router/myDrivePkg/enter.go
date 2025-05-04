package myDrivePkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ MyDriversRouter }

var myDriversApi = api.ApiGroupApp.MyDrivePkgApiGroup.MyDriversApi
