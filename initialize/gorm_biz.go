package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/productPkg"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(productPkg.Products{}, alertRulePkg.Alarms{}, alertRecordPkg.Records{})
	err = db.AutoMigrate(alertRulePkg.Alarms{}, alertRecordPkg.Records{})
	if err != nil {
		return err
	}
	return nil
}
