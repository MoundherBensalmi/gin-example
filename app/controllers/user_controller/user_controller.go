package user_controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"users": "users",
	})
}
