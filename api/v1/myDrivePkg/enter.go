package myDrivePkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ MyDriversApi }

var myDriversService = service.ServiceGroupApp.MyDrivePkgServiceGroup.MyDriversService
