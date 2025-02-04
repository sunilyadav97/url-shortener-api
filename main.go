package main

import (
	"log"
	"os"

	"url-shortener-api/database"
	"url-shortener-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to MongoDB
	database.ConnectMongo()

	// Create Gin router instance
	router := gin.Default()

	// Setup API routes
	routes.SetupRoutes(router)

	// Get port from environment variables
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Starting server on port", port)
	router.Run(":" + port)
}
