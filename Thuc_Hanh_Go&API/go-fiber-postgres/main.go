package main

import (
    "log" // Add this line
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "gorm.io/gorm"
)
type Repository struct{
	DB *gorm.DB
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	app.Get("/api/posts", r.GetPosts)
	app.Post("/api/posts", r.CreatePost)
	// Add other routes as needed
	api.Delete("/api/posts/:id", r.DeletePost)
}
func mainthuchanh() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}
	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}