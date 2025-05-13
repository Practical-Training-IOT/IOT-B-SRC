// 自动生成模板Products
package productPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// products表 结构体  Products
type Products struct {
	global.GVA_MODEL
	Platform           *string `json:"platform" form:"platform" gorm:"default:cloud;comment:存储平台;column:platform;"`                              //存储平台
	ProductName        *string `json:"productName" form:"productName" gorm:"comment:产品名称;column:product_name;"`                                  //产品名称
	Category           *string `json:"category" form:"category" gorm:"comment:所属品类;column:category;" binding:"required"`                         //所属品类
	SelectCategory     *string `json:"selectCategory" form:"selectCategory" gorm:"comment:选择标准品类;column:select_category;" binding:"required"`    //选择标准品类
	NodeType           *string `json:"nodeType" form:"nodeType" gorm:"comment:节点类型;column:node_type;" binding:"required"`                        //节点类型
	GatewayProtocol    *string `json:"gatewayProtocol" form:"gatewayProtocol" gorm:"comment:接入网关协议;column:gateway_protocol;" binding:"required"` //接入网关协议
	DataFormat         *string `json:"dataFormat" form:"dataFormat" gorm:"comment:数据格式;column:data_format;" binding:"required"`                  //数据格式
	NetworkType        *string `json:"networkType" form:"networkType" gorm:"comment:网络类型;column:network_type;" binding:"required"`               //网络类型
	Factory            *string `json:"factory" form:"factory" gorm:"comment:工厂;column:factory;"`                                                 //工厂
	ProductDescription *string `json:"productDescription" form:"productDescription" gorm:"comment:产品描述;column:product_description;"type:text;"`  //产品描述
	ProductCode        *string `json:"productCode" form:"productCode" gorm:"comment:产品编号;column:product_code;"`                                  //产品编号
	ProductStatus      *string `json:"productStatus" form:"productStatus" gorm:"comment:产品状态;column:product_status;"`                            //产品状态
}

// TableName products表 Products自定义表名 products
func (Products) TableName() string {
	return "products"
}

// Category 表示一个标准品类
type Category struct {
	ID           string    `gorm:"column:id;type:varchar(20);primary_key;not_null"`
	CategoryName string    `gorm:"column:category_name;type:varchar(100);not_null"`
	CategoryKey  string    `gorm:"column:category_key;type:varchar(50);not_null"`
	Scene        string    `gorm:"column:scene;type:varchar(50);not_null"`
	CreatedAt    time.Time `gorm:"column:created_at;type:timestamp"`
}

// 如果需要自定义表名，可以实现 gorm.Model 接口的方法 TableName
func (Category) TableName() string {
	return "categories" // 假设你的表名为 categories
}

// Property represents a row in the properties table.
type Property struct {
	ID          string   `gorm:"column:id;primaryKey" json:"id"`
	CategoryKey string   `gorm:"column:category_key" json:"category_key"`
	Name        string   `gorm:"column:name" json:"name"`
	Code        string   `gorm:"column:code" json:"code"`
	AccessMode  string   `gorm:"column:access_mode" json:"access_mode"`
	DataType    string   `gorm:"column:data_type" json:"data_type"`
	Unit        *string  `gorm:"column:unit" json:"unit,omitempty"` // Use pointer to allow NULL values
	MinValue    *float64 `gorm:"column:min_value" json:"min_value,omitempty"`
	MaxValue    *float64 `gorm:"column:max_value" json:"max_value,omitempty"`
}

// TableName returns the table name for the Property model.
func (Property) TableName() string {
	return "properties"
}
