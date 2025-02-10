package api_routes

import (
	"MBFacto/app/controllers/user_controller"
	"github.com/gin-gonic/gin"
)

func RegisterAPIUserRoutes(api *gin.RouterGroup) {
	users := api.Group("/users")
	{
		users.GET("/", user_controller.GetAll)
	}
}
