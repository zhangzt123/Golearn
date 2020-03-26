package Repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/zhangzt123/Golearn/Gin/conf/datasource"
	"github.com/zhangzt123/Golearn/Gin/entity"
)

var ds *gorm.DB

var rec = func() {
	if e := recover(); e != nil {
		//todo
		fmt.Println(e)
		ds.Rollback()
		return
	}
	ds.Commit()

}

func init() {
	ds = datasource.NewInstance()
}

func Findalluser() (u []entity.User) {
	defer rec()
	ds.Find(&u)
	return
}
