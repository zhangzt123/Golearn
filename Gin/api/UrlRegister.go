/*
 init url register
*/

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzt123/Golearn/Gin/api/v1"
)

type reqtype int

//
const (
	GET reqtype = iota
	POST
	PUT
	DELETE
	TRACE
	Any
	PATCH
)

type GroupV struct {
	Apiname  string
	Httppool reqtype
	Handlers gin.HandlerFunc
}

type k string

var Mygroup = make(map[k]*GroupV)

func init() {
	Mygroup["/v1"] = &GroupV{Apiname: "/GetName/:name", Httppool: GET, Handlers: v1.Gethelloworld}

}
