package main

import (
	"fmt"
	"social_media_sever/config"
	"social_media_sever/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to database
	config.ConnectDatabase()

	// Set GIN mode
	gin.SetMode(config.Config.GinMode)

	// Setup router
	router := routes.SetupRouter()

	// Start server
	port := fmt.Sprintf(":%s", config.Config.ServerPort)
	router.Run(port)
}
