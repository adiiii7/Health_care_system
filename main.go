// main.go

package main

import (
	"fmt"
	"healthcareapp/config" // Import your configuration package
	"healthcareapp/routes" // Import your routes package

	"github.com/gin-gonic/gin"
)

func main() {
	// Load application configuration
	appConfig, err := config.LoadConfig()
	if err != nil {
		panic("Failed to load configuration: " + err.Error())
	}

	// Initialize the Gin router
	router := gin.Default()

	// Access configuration settings
	fmt.Println("Application Name:", appConfig.AppName)
	fmt.Println("App Port:", appConfig.AppPort)
	fmt.Println("Database Host:", appConfig.Database.Host)

	// Set up application routes
	routes.SetupRoutes(router)

	// Start the server
	serverAddr := fmt.Sprintf(":%d", appConfig.AppPort)
	router.Run(serverAddr)
}
