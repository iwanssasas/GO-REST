package main

import (
	"API-SIMPLE/controllers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRouter()
	_ = r.Run(":8080")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "gin connected")
	})

	userRepo := controllers.NewUser()
	r.POST("/users", userRepo.CreateUser)
	r.GET("/users", userRepo.GetUsers)
	r.GET("/users/:id", userRepo.GetUser)
	r.PUT("/users/:id", userRepo.UpdateUser)
	r.DELETE("/users/:id", userRepo.DeleteUser)

	bookRepo := controllers.NewBook()
	r.POST("/books", bookRepo.CreateBook)
	r.GET("/books", bookRepo.GetBooks)
	r.GET("/books/:id", bookRepo.GetBook)
	r.PUT("/books/:id", bookRepo.UpdateBook)
	r.DELETE("/books/:id", bookRepo.DeleteBook)

	return r
}
