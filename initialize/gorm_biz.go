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
	"github.com/flipped-aurora/gin-vue-admin/server/model/scene"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(productPkg.Products{}, devicePkg.Devices{}, driverPkg.Drivers{}, myDrivePkg.MyDrivers{}, scene.Scenes{})
	err = db.AutoMigrate(productPkg.Products{}, resources_iot.Resources{}, rules.RuleInfo{}, scene.Scenes{})
	err = db.AutoMigrate(alertRulePkg.AlertRules{}, alertRecordPkg.Alerts{}, rules.RuleInfo{}, scene.Scenes{})
	if err != nil {
		return err
	}
	return nil
}
