package router_v1

import (
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	RegisterPFCModelRouter(v1.Group(""))

	RegisterCommonRouter(v1.Group(""))

	RegisterLoginRouter(v1.Group(""))

}
