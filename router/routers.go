package router

import (
	v1 "mygin/router/v1"

	"github.com/gin-gonic/gin"
)

// type RouterGroup struct {
// 	V1Routers v1.V1RouterGroup
// }

// var RouterGroupAPP = new(RouterGroup)

func InitRouter() *gin.Engine {
	engine := gin.Default()

	new(v1.V1RouterGroup).InitV1Router(engine)
	// RouterGroupAPP.V1Routers.InitV1Router(engine)
	// K8sGroup := RouterGroupAPP.V1Routers.K8SRouterGroup
	// K8sGroup.InitK8SRouter(engine)
	return engine
}
