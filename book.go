package main

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// Setup mock data using struct
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

// Initialize in-memory data
var books []Book = []Book{
	{ID: 1, Title: "1984", Author: "Kit"},
	{ID: 2, Title: "MM", Author: "Kit"},
}

// Handler functions
// getBooks godoc
// @Summary Get all books
// @Description Get details of all books
// @Tags books
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {array} Book
// @Router /book [get]
func getBooks(c *fiber.Ctx) error {
	return c.JSON(books)
}

func getBook(c *fiber.Ctx) error {
	// handle error string convert to int
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// find book by id
	for _, book := range books {
		if book.ID == bookID {
			return c.JSON(book)
		}
	}

	// handle not found book
	return c.Status(fiber.StatusNotFound).SendString("Not found")
}

func createBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	books = append(books, *book)
	return c.JSON(book)
}

func updateBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	bookUpdate := new(Book)
	if err := c.BodyParser(bookUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == bookID {
			books[index].Title = bookUpdate.Title
			books[index].Author = bookUpdate.Author
			return c.JSON(books[index])
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not found")
}

func deleteBook(c *fiber.Ctx) error {
	bookID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for index, book := range books {
		if book.ID == bookID {
			books = append(books[:index], books[index+1:]...)
			return c.SendStatus(fiber.StatusNoContent)
		}
	}

	return c.Status(fiber.StatusNotFound).SendString("Not found")
}
