package scene

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ScenesRouter }

var scenesApi = api.ApiGroupApp.SceneApiGroup.ScenesApi
