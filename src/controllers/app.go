package controllers

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
	service "github.com/igudgz/desafioMoneri/src/services"
	"github.com/joho/godotenv"
)

func StartServer() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Port := os.Getenv("PORT")

	app := fiber.New()
	app.Get("/api/books/:id?", service.GetBook)
	app.Get("/api/books/title/:title?", service.GetBookAuthor)
	app.Get("/api/books/author/:author?", service.GetBookTitle)
	app.Post("/api/books", service.CreateBook)
	app.Put("/api/books/:id", service.UpdateBook)
	app.Delete("/api/books/:id", service.DeleteBook)
	app.Listen(Port)
}
