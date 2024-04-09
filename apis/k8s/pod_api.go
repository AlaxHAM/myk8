package k8s

import "github.com/gin-gonic/gin"

type PodApi struct {
}

func (*PodApi) GetPodListOrDetail(ctx *gin.Context) {
	namespace := ctx.Params("namespace")
	name := ctx.Query("name")
	keyword := ctx.Query("keyword")

	if name != "" {
		
	} else {

	}
}
