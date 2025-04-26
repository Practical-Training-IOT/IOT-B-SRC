package productPkg

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct{ ProductsApi }

var productsService = service.ServiceGroupApp.ProductPkgServiceGroup.ProductsService
