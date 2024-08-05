package main

import (
	"deathtiny_encounters/config"
	"deathtiny_encounters/repositories"
	"deathtiny_encounters/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize PostgreSQL connection by gorm
	db, err := gorm.Open(postgres.Open(config.GetDatabaseURL()), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	repositories.InitDB(db)

	// Create a new Gin router
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router)

	// Start the server
	router.Run(":8080")
}
