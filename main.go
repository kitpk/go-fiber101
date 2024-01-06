package main

import (
	"os"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/kitpk/go-fiber101/docs"
)

// @title Book API
// @description This is a sample server for a book API.
// @version 1.0
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("load .env error")
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)

	// Apply CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "+",
		AllowMethods: "GET, POST, HEAD, PUT, DELETE, PATCH",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Login route
	app.Post("/login", login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	// Middleware
	bookGroup := app.Group("/book", checkMiddleware)

	// Setup routes
	bookGroup.Get("/", getBooks)
	bookGroup.Get("/:id", getBook)
	bookGroup.Post("/", createBook)
	bookGroup.Put("/:id", updateBook)
	bookGroup.Delete("/:id", deleteBook)

	app.Post("/upload", uploadFile)

	app.Listen(":" + os.Getenv("PORT"))
}
