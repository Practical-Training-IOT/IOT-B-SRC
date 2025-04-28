package initialize

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/alertRulePkg"
)

func bizModel() error {
	db := global.GVA_DB
	err := db.AutoMigrate(alertRulePkg.AlertRules{}, alertRecordPkg.Alerts{})
	if err != nil {
		return err
	}
	return nil
}
