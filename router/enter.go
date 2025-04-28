package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/alertRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alertRulePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/alterLevelPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System         system.RouterGroup
	Example        example.RouterGroup
	AlertRecordPkg alertRecordPkg.RouterGroup
	AlterLevelPkg  alterLevelPkg.RouterGroup
	AlertRulePkg   alertRulePkg.RouterGroup
}
