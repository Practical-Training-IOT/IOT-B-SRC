package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(productPkg.Products{}, devicePkg.Devices{}, driverPkg.Drivers{}, myDrivePkg.MyDrivers{})
	if err != nil {
		return err
	}
	return nil
}
