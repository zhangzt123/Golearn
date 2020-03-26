//

package datasource

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	c "github.com/zhangzt123/Golearn/Gin/conf"
)

//Datasource0    .

var datasource0 *gorm.DB

//var dataSourceName = "zzt:Zzt=123456@tcp(192.168.118.130:3306)/nacosdb?charset=utf8&parseTime=True&loc=Local"
var dataSourceName = c.Conf.GetString("db.dsn")

func NewInstance() *gorm.DB {
	datasource0, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		_ = datasource0.Close()
		panic(fmt.Sprintf("database connect fail :%v \n", err))
	}
	datasource0.DB().SetConnMaxLifetime(c.Conf.GetDuration("db.maxLifetime"))
	datasource0.DB().SetMaxIdleConns(c.Conf.GetInt("db.maxIdleConns"))
	datasource0.DB().SetMaxOpenConns(c.Conf.GetInt("db.maxOpenConns"))
	return datasource0
}

func Close() error {
	return datasource0.Close()
}
