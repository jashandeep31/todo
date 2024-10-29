package main

import (
	"github/jashandeep31/todo/database"
	"github/jashandeep31/todo/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading the .env file

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// database connection

	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()
	// Routes of the application
	routes.SetupRoutes(r)

	// loading and setting and port and other configs
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Starting server on port %s", port)

	if err := r.Run(":" + port); err != nil {
		log.Fatal("Server could not start: ", err)
	}
}
