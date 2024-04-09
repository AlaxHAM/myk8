package harbor

import (
	"github.com/gin-gonic/gin"
)

func RegisterHarborRouter(rg *gin.RouterGroup) {
	// 定义 harbor 相关操作的路由
	rg.POST("/push/image", PushImageTest)

}

func PushImageTest(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "push image success"})
}
