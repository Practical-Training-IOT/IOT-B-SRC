package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/devicePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/driverPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/myDrivePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
	"github.com/flipped-aurora/gin-vue-admin/server/model/rules"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(productPkg.Products{}, devicePkg.Devices{}, driverPkg.Drivers{}, myDrivePkg.MyDrivers{})
	err = db.AutoMigrate(productPkg.Products{}, resources_iot.Resources{}, rules.RuleInfo{})
	err = db.AutoMigrate(alertRulePkg.Alarms{}, alertRecordPkg.Records{})
	if err != nil {
		return err
	}
	return nil
}
