//

package datasource

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Datasource0    .

var datasource0 *gorm.DB

var dataSourceName = "zzt:Zzt=123456@tcp(192.168.118.130:3306)/nacosdb?charset=utf8&parseTime=True&loc=Local"

func NewInstance() *gorm.DB {
	datasource0, err := gorm.Open("mysql", dataSourceName)
	if err != nil {
		datasource0.Close()
		panic(fmt.Sprintf("database connect fail :%v \n", err))
	}
	return datasource0
}

func Close() error {
	return datasource0.Close()
}
