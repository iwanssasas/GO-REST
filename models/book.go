package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID       uint64 `gorm:"primary_key:auto_increment" json:"id"`
	Tittle   string `gorm:"type:varchar(255)" json:"tittle"`
	Quantity int    `gorm:"type:int(255)" json:"quantity"`
}

//create a user
func CreateBook(db *gorm.DB, Book *Book) (err error) {
	err = db.Create(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetBooks(db *gorm.DB, Book *[]Book) (err error) {
	err = db.Find(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetBook(db *gorm.DB, Book *Book, id string) (err error) {
	err = db.Where("id = ?", id).First(Book).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateBook(db *gorm.DB, Book *Book) (err error) {
	db.Save(Book)
	return nil
}

//delete user
func DeleteBook(db *gorm.DB, Book *Book, id string) (err error) {
	db.Where("id = ?", id).Delete(Book)
	return nil
}
