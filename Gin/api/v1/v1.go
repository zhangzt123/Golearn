package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zhangzt123/Golearn/Gin/Repository"
	"net/http"
)

//Get
//
var Gethelloworld = func(c *gin.Context) {
	str := c.Param("name")
	c.String(http.StatusOK, fmt.Sprintln("hello "+str))
}

var FindAllUser = func(c *gin.Context) {

	userslice := Repository.Findalluser()
	//bt,_:=json.Marshal(userslice)
	c.String(http.StatusOK, fmt.Sprintln(userslice))
}
