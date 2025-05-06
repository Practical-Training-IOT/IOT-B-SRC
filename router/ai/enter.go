package ai

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ AiRouter }

var aiApi = api.ApiGroupApp.AIApiGroup.ScenesApi
