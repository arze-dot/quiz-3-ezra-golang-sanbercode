package controllers

import (
	"net/http"
	"quiz-3-ezra-golang-sanbercode/database"
	"quiz-3-ezra-golang-sanbercode/repository"
	"quiz-3-ezra-golang-sanbercode/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllCategories retrieves all categories from the database
func GetAllCategories(c *gin.Context) {
	var result gin.H

	categories, err := repository.GetAllCategory(database.DbConnection)

	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
	} else {
		result = gin.H{
			"result": categories,
		}
	}

	c.JSON(http.StatusOK, result)
}

// InsertCategory inserts a new category into the database
func InsertCategory(c *gin.Context) {
	var category structs.Category

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	err = repository.InsertCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

// UpdateCategory updates an existing category in the database
func UpdateCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.BindJSON(&category)
	if err != nil {
		panic(err)
	}

	category.ID = id

	err = repository.UpdateCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}

// DeleteCategory deletes a category from the database
func DeleteCategory(c *gin.Context) {
	var category structs.Category
	id, _ := strconv.Atoi(c.Param("id"))

	category.ID = id
	err := repository.DeleteCategory(database.DbConnection, category)
	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, category)
}
