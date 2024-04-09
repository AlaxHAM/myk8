package v1

import (
	k8sApi "mygin/router/v1/k8s"

	harborApi "mygin/router/v1/harbor"

	"github.com/gin-gonic/gin"
)

type V1RouterGroup struct {
}

func (*V1RouterGroup) InitV1Router(engine *gin.Engine) {
	// 初始化路由，根据 版本 和 资源注册路由组
	V1ApiGroup := engine.Group("api/v1")
	{
		k8sGroup := V1ApiGroup.Group("k8s")
		{
			k8sApi.RegisterK8SRouter(k8sGroup)
		}
		harborGroup := V1ApiGroup.Group("harbor")
		{
			harborApi.RegisterHarborRouter(harborGroup)
		}
	}

}
