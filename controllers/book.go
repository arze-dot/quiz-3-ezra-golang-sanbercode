package controllers

import (
	"net/http"
	"quiz-3-ezra-golang-sanbercode/database"
	"quiz-3-ezra-golang-sanbercode/repository"
	"quiz-3-ezra-golang-sanbercode/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllBooks retrieves all books from the database
func GetAllBooks(c *gin.Context) {
	var result gin.H

	books, err := repository.GetAllBook(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": books,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertBook inserts a new book into the database
func InsertBook(c *gin.Context) {
	var book structs.Book

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	err = repository.InsertBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}

// UpdateBook updates an existing book in the database
func UpdateBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&book)
	if err != nil {
		panic(err)
	}

	book.ID = id

	err = repository.UpdateBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book from the database
func DeleteBook(c *gin.Context) {
	var book structs.Book
	id, _ := strconv.Atoi(c.Param("id"))

	book.ID = id
	err := repository.DeleteBook(database.DbConnection, book)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, book)
}
