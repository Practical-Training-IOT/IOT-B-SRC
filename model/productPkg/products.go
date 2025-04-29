
// 自动生成模板Products
package productPkg
import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// products表 结构体  Products
type Products struct {
    global.GVA_MODEL
  Platform  *string `json:"platform" form:"platform" gorm:"default:cloud;comment:存储平台;column:platform;"`  //存储平台
  ProductName  *string `json:"productName" form:"productName" gorm:"comment:产品名称;column:product_name;"`  //产品名称
  Category  *string `json:"category" form:"category" gorm:"comment:所属品类;column:category;" binding:"required"`  //所属品类
  SelectCategory  *string `json:"selectCategory" form:"selectCategory" gorm:"comment:选择标准品类;column:select_category;" binding:"required"`  //选择标准品类
  NodeType  *string `json:"nodeType" form:"nodeType" gorm:"comment:节点类型;column:node_type;" binding:"required"`  //节点类型
  GatewayProtocol  *string `json:"gatewayProtocol" form:"gatewayProtocol" gorm:"comment:接入网关协议;column:gateway_protocol;" binding:"required"`  //接入网关协议
  DataFormat  *string `json:"dataFormat" form:"dataFormat" gorm:"comment:数据格式;column:data_format;" binding:"required"`  //数据格式
  NetworkType  *string `json:"networkType" form:"networkType" gorm:"comment:网络类型;column:network_type;" binding:"required"`  //网络类型
  Factory  *string `json:"factory" form:"factory" gorm:"comment:工厂;column:factory;"`  //工厂
  ProductDescription  *string `json:"productDescription" form:"productDescription" gorm:"comment:产品描述;column:product_description;"type:text;"`  //产品描述
}


// TableName products表 Products自定义表名 products
func (Products) TableName() string {
    return "products"
}





