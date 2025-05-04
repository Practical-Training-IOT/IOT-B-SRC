package driverPkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ DriversRouter }

var driversApi = api.ApiGroupApp.DriverPkgApiGroup.DriversApi
