/*
 init url register
*/

package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzt123/Golearn/Gin/api/v1"
)

type Reqtype int

//
const (
	GET Reqtype = iota
	POST
	PUT
	DELETE
	TRACE
	Any
	PATCH
)

type Subgrouptype struct {
	Subgroup string
	Httptype Reqtype
	Handler  gin.HandlerFunc
}

type Grouptype struct {
	Group         string
	Subgrouptypes []Subgrouptype
}

func Register(fun func(group []*Grouptype)) {
	k := []*Grouptype{
		{
			Group: "/v1",
			Subgrouptypes: []Subgrouptype{
				{Subgroup: "/GetName/:name", Httptype: GET, Handler: v1.Gethelloworld},
				{Subgroup: "/FindAllUser", Httptype: GET, Handler: v1.FindAllUser} /**/}}}
	fun(k)

}
