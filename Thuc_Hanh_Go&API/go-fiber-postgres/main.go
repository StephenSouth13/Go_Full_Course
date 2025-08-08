package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	// IMPORT CÁC PACKAGE NỘI BỘ
	// (Giả sử tên module của bạn là go-fiber-postgres)
	"github.com/stephensouth13/go-fire-postgres/models"
	"github.com/stephensouth13/go-fire-postgres/storage"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

// Repository struct to hold the database connection
type Repository struct {
	DB *gorm.DB
}

// SetupRoutes registers all API routes
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	api.Post("/create_book", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
	api.Put("/update_book/:id", r.UpdateBook)
}

// CreateBook handles the creation of a new book
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := models.Book{}
	err := context.BodyParser(&book)
	if err != nil {
		return context.Status(fiber.StatusUnprocessableEntity).JSON(
			fiber.Map{"message": "Invalid request body"})
	}

	err = r.DB.Create(&book).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Failed to create book"})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book created successfully",
	})
}

// GetBooks retrieves all books from the database
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModels := &[]models.Book{}

	err := r.DB.Find(bookModels).Error
	if err != nil {
		return context.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Failed to retrieve books"})
	}
	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Books retrieved successfully",
		"data":    bookModels,
	})
}

// DeleteBook deletes a book by its ID
func (r *Repository) DeleteBook(context *fiber.Ctx) error {
	bookModel := &models.Book{}
	id := context.Params("id")
	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Book ID is required"})
	}

	// GORM's Delete method returns a *gorm.DB object, not an error directly
	result := r.DB.Delete(bookModel, id)
	if result.Error != nil {
		return context.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Failed to delete book"})
	}

	// GORM sẽ trả về result.RowsAffected == 0 nếu không tìm thấy bản ghi
	if result.RowsAffected == 0 {
		return context.Status(http.StatusNotFound).JSON(
			fiber.Map{"message": "Book not found"})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book deleted successfully",
	})
}

// GetBookByID retrieves a single book by its ID
func (r *Repository) GetBookByID(context *fiber.Ctx) error {
	id := context.Params("id")
	bookModel := &models.Book{}
	if id == "" {
		return context.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Book ID is required"})
	}

	fmt.Println("The ID is:", id)
	err := r.DB.Where("id = ?", id).First(bookModel).Error
	if err != nil {
		return context.Status(http.StatusNotFound).JSON(
			fiber.Map{"message": "Book not found"})
	}

	return context.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Book retrieved successfully",
		"data":    bookModel,
	})
}

// UpdateBook updates an existing book
func (r *Repository) UpdateBook(context *fiber.Ctx) error {
	// Implement logic for updating a book here
	return context.Status(http.StatusNotImplemented).JSON(
		fiber.Map{"message": "Update functionality not yet implemented"})
}

// main function is the entry point of the application
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}
	// Get the port from the environment variable as a string
    portStr := os.Getenv("DB_PORT")

    // Convert the string port to an integer
    port, err := strconv.Atoi(portStr)
    if err != nil {
        log.Fatalf("Invalid DB_PORT in .env file: %v", err)
    }


	config := &storage.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DbName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	db, err := storage.NewConnection(config)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	

	// Ensure database migration is done
	err = models.MigrateBooks(db)
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
