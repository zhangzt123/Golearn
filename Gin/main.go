package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zhangzt123/Golearn/Gin/api"
	_ "github.com/zhangzt123/Golearn/Gin/conf"
)

var engine *gin.Engine
var err error

func init() {
	gin.SetMode(gin.DebugMode)
	engine = gin.Default()
	mp := api.Mygroup
	for k, v := range mp {
		groupRoute := engine.Group(string(k))
		switch v.Httppool {
		case api.GET:
			groupRoute.GET(v.Apiname, v.Handlers)
		case api.POST:
			groupRoute.POST(v.Apiname, v.Handlers)
		case api.PUT:
			groupRoute.PUT(v.Apiname, v.Handlers)
		case api.DELETE:
			groupRoute.DELETE(v.Apiname, v.Handlers)
		case api.PATCH:
			groupRoute.PATCH(v.Apiname, v.Handlers)
		case api.Any:
			groupRoute.Any(v.Apiname, v.Handlers)
		default:
			groupRoute.Any(v.Apiname, v.Handlers)
		}
	}

}

func main() {

	err = engine.Run(":8080")
	if err != nil {
		//todo
	}

}
