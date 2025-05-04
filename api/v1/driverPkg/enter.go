package driverPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ DriversApi }

var driversService = service.ServiceGroupApp.DriverPkgServiceGroup.DriversService
