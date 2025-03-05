package router_v1

import (
	"web-api/internal/api/controllers"

	"github.com/gin-gonic/gin"
)

// API - LOGIN
func RegisterLoginRouter(router *gin.RouterGroup) {

	router.POST("/login", controllers.LG.Login)
}
