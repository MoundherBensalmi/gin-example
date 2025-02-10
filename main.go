package main

import (
	"MBFacto/config"
	"MBFacto/database"
	"MBFacto/routes"
	"MBFacto/utils/jwt"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	config.Load()
	database.ConnectToDB()
}

func main() {
	defer database.CloseDB()

	if config.Cfg.APP.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		accessToken, refreshToken, err := jwt_helper.GenerateTokens(1)
		if err != nil {
			c.JSON(403, gin.H{
				"error": "unauthorized access",
			})
			c.Abort()
		}

		c.JSON(200, gin.H{
			"message":      "Hello, MBFacto!",
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
	})

	routes.RegisterAPIRoutes(r)

	log.Printf("Starting server on http://localhost:%s\n", config.Cfg.APP.Port)
	if err := r.Run(":" + config.Cfg.APP.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
