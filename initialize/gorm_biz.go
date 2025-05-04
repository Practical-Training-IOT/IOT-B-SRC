package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/resources_iot"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(productPkg.Products{}, resources_iot.Resources{}, alertRulePkg.Alarms{})
	err = db.AutoMigrate(alertRulePkg.Alarms{})
	if err != nil {
		return err
	}
	return nil
}
