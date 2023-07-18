// making an libray api
package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

// List of books in libary
var books = []book{
	{ID: "1", Title: "In search of lost time", Author: "Marcel Proust", Quantity: 2},
	{ID: "2", Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Quantity: 5},
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 0},
}

// get all the books in json format from the list
func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

// add books to list
func addbooks(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	// append book to slice
	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// get a book by id
/* func getById(c *gin.Context) {
	var id string
	if err := c.GetString(&id); err != nil {
		return
	}

} */

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	//	c.IndentedJSON(http.StatusOK, book)
	c.IndentedJSON(http.StatusOK, book.Quantity)

	if book.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Quantity": "Out of quantity"})
	}
}

// Check if exist
func getBookById(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// Check out
func checkOutById(c *gin.Context) {
	//	lets take id and based id lets return the quantity
	/*id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book does not exist"})
	}

	c.IndentedJSON(http.StatusOK, book.Quantity)*/

}

func checkOutQuantity(id string) (*book, error) {

	return nil, errors.New("not found")
}

func main() {
	// println("Hello, world")
	router := gin.Default()

	// get request to list all books
	router.GET("/books", getBooks)

	// post request to add books
	router.POST("/books", addbooks)

	// Check using id
	router.GET("/books/:id", bookById)

	//	Check quantity
	//	router.GET("/books/:id", checkOutById)

	// run and list server
	router.Run("localhost:8080")
}
