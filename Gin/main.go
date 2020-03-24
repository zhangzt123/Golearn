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
	engine.Use()
	//api.Register(regfun)
	api.Register(regfun)
}

var regfun = func(grouptypes []*api.Grouptype) {
	for _, grouptype := range grouptypes {
		groupRoute := engine.Group(grouptype.Group)
		for _, sub := range grouptype.Subgrouptypes {
			switch sub.Httptype {
			case api.GET:
				groupRoute.GET(sub.Subgroup, sub.Handler)
			case api.POST:
				groupRoute.POST(sub.Subgroup, sub.Handler)
			case api.PUT:
				groupRoute.PUT(sub.Subgroup, sub.Handler)
			case api.DELETE:
				groupRoute.DELETE(sub.Subgroup, sub.Handler)
			case api.PATCH:
				groupRoute.PATCH(sub.Subgroup, sub.Handler)
			case api.Any:
				groupRoute.Any(sub.Subgroup, sub.Handler)
			default:
				groupRoute.Any(sub.Subgroup, sub.Handler)
			}
		}
	}

}

func main() {

	err = engine.Run(":8080")
	if err != nil {
		//todo
	}

}
