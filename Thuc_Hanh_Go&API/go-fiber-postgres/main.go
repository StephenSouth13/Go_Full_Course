package main

import (
	"log" // Add this line
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
	Publisher string `json:"publisher"`
	}
type Repository struct{
	DB *gorm.DB
}
func (r *Repository) SetupRoutes(app *fiber.App) {
	api := app.Group("/api")
	app.Post("/create_book", r.CreateBook)
	api.Delete("/delete_book/:id", r.DeleteBook)
	api.Get("/get_book/:id", r.GetBookByID)
	api.Get("/books", r.GetBooks)
	api.Put("/update_book/:id", r.UpdateBook)
}
// Create Book
func (r *Repository) CreateBook(context *fiber.Ctx) error {
	book := Book{}
	err := context.BodyParser(&book)
	if err != nil {
		context.Status(fiber.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "Invalid request body"})
			return err
		}
		err = r.DB.Create(&book).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Failed to create book"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Book created successfully",})
	return nil	

}
// Get Book
func (r *Repository) GetBooks(context *fiber.Ctx) error {
	bookModel := &[]models.Book{}

	err := r.DB.Find(bookModel).Error
	if err != nil {
		context.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Failed to retrieve books"})
		return err
	}
	context.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Books retrieved successfully",
		"data":    bookModels,
	})
	return nil
}
// Delete Book
func (r *Repository) DeleteBook(context *fiber.Ctx) error {

}
// Update Book
func (r *Repository) UpdateBook(context *fiber.Ctx) error {
		
	

func main() {
	err := godotenv.Load(".env")
	if err != nil{
		log.Fatal(err)
	}
	db, err := storage.NewConnection(config)
	if err != nil{
		log.Fatal("Failed to connect to database:", err)
	}

	r := Repository{
		DB: db,
	}
	app := fiber.New()
	r.SetupRoutes(app)
	app.Listen(":8080")
}