package routes

import (
	"MBFacto/routes/api_routes"
	"github.com/gin-gonic/gin"
)

func RegisterAPIRoutes(router *gin.Engine) {
	api := router.Group("/api")
	api_routes.RegisterAPIUserRoutes(api)
}
