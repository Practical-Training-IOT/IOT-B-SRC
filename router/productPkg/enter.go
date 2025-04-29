package productPkg

import api "github.com/flipped-aurora/gin-vue-admin/server/api/v1"

type RouterGroup struct{ ProductsRouter }

var productsApi = api.ApiGroupApp.ProductPkgApiGroup.ProductsApi
