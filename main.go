package main

import (
	"MBFacto/config"
	"MBFacto/database"
	"MBFacto/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	config.Load()
	database.ConnectToDB()
}

func main() {
	cfg := config.Get()
	defer database.CloseDB()

	if cfg.APP.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, MBFacto!",
		})
	})

	routes.RegisterAPIRoutes(r)

	log.Printf("Starting server on http://localhost:%s\n", cfg.APP.Port)
	if err := r.Run(":" + cfg.APP.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
