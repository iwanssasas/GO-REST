package controllers

import (
	"API-SIMPLE/database"
	"API-SIMPLE/models"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BookRepo struct {
	Db *gorm.DB
}

func NewBook() *BookRepo {
	db := database.InitDb()
	db.AutoMigrate(&models.Book{})
	return &BookRepo{Db: db}
}

//create user
func (repository *BookRepo) CreateBook(c *gin.Context) {
	var book models.Book
	c.BindJSON(&book)
	err := models.CreateBook(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

//get users
func (repository *BookRepo) GetBooks(c *gin.Context) {
	var book []models.Book
	err := models.GetBooks(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

func (repository *BookRepo) GetBook(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var book models.Book
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// update user
func (repository *BookRepo) UpdateBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.GetBook(repository.Db, &book, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.BindJSON(&book)
	err = models.UpdateBook(repository.Db, &book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, book)
}

// delete user
func (repository *BookRepo) DeleteBook(c *gin.Context) {
	var book models.Book
	id, _ := c.Params.Get("id")
	err := models.DeleteBook(repository.Db, &book, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
