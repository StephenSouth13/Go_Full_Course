package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"my-go-backend/models"
	"my-go-backend/routes"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate database schema
	db.AutoMigrate(&models.Post{})

	// Initialize Gin router
	r := gin.Default()

	// Setup routes
	routes.SetupRoutes(r, db)

	// Run server
	r.Run() // listen and serve on 0.0.0.0:8080
}