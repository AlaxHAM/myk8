package k8sApi

import (
	"github.com/gin-gonic/gin"
)

func RegisterK8SRouter(rg *gin.RouterGroup) {
	// 定义 k8s 相关操作的路由

	rg.GET("/namespace", )


  //****************** Pod ******************//
	rg.GET("/pod", GetPodTest)
	rg.POST("/pod", GetPodTest)
	rg.DELETE("/pod/:namespace/:name", GetPodTest)
	
	
	rg.GET("/deployment", GetDeploymentTest)

}

func GetPodTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "get k8s pod success"})
}

func GetDeploymentTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "get k8s deployment success"})
}
