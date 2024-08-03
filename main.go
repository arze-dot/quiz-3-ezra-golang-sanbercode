package main

import (
	"database/sql"
	"fmt"
	"log"
	"quiz-3-ezra-golang-sanbercode/controllers"
	"quiz-3-ezra-golang-sanbercode/database"
	"quiz-3-ezra-golang-sanbercode/middleware"
	"quiz-3-ezra-golang-sanbercode/seeder"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "quiz_3"
)

var (
	DB  *sql.DB
	err error
)

func main() {
	psqlInfo := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", psqlInfo)
	defer DB.Close()
	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	database.DBMigrate(DB)

	psqlSeeder := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname,
	)

	DB, err = sql.Open("postgres", psqlSeeder)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer DB.Close()

	// Seed the database
	seeder.SeedUsers(DB)
	fmt.Println("Successfully connected!")

	router := gin.Default()

	router.POST("/login", controllers.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/book", controllers.GetAllBooks)
		authorized.POST("/book", controllers.InsertBook)
		authorized.PUT("/book/:id", controllers.UpdateBook)
		authorized.DELETE("/book/:id", controllers.DeleteBook)

		authorized.GET("/category", controllers.GetAllCategories)
		authorized.POST("/category", controllers.InsertCategory)
		authorized.PUT("/category/:id", controllers.UpdateCategory)
		authorized.DELETE("/category/:id", controllers.DeleteCategory)

		authorized.GET("/user", controllers.GetAllUsers)
		authorized.POST("/user", controllers.InsertUser)
		authorized.PUT("/user/:id", controllers.UpdateUser)
		authorized.DELETE("/user/:id", controllers.DeleteUser)
	}

	router.Run(":8080")
}
