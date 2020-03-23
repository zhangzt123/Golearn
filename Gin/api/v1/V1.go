package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Get
//
var Gethelloworld = func(c *gin.Context) {
	str := c.Param("name")
	c.String(http.StatusOK, fmt.Sprintln("hello "+str))
}
