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
	{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 10},
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

// Check using id
func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, book)

	//	Just checking somethings out about the checkout
	/*c.IndentedJSON(http.StatusOK, book.Quantity)*/

	/*if book.Quantity == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"Quantity": "Out of quantity"})
	} else {
		book.Quantity--
	}*/
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

// Checkout
func checkOutById(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
	}

	book, _ := getBookById(id)

	if book.Quantity == 0 {
		c.JSON(http.StatusForbidden, gin.H{"message": "Book not avaiable"})
		// os.Exit(1) // doing this will aslo stop the server instead use the return keyword
		return
	}

	if book.Quantity <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Book not avaiable"})
	}

	book.Quantity -= 1
	c.JSON(http.StatusOK, book)
}

func returnBook(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "missing id query parameter"})
	}

	book, err := getBookById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"messgae": "something went wrong"})
	}

	book.Quantity += 1
//	print(book.Quantity)
	c.JSON(http.StatusOK, book)

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
	router.GET("/checkout", checkOutById)

	//	return quantity
	router.GET("/return", returnBook)

	// run and list server
	router.Run("localhost:8080")
}
