package repository

import (

	"github.com/dilmurodov/Book_catalog/models"
	"github.com/jinzhu/gorm"
)

type BooksPostgres struct {
	db *gorm.DB
}

func NewBooksPostgres(db *gorm.DB) *BooksPostgres {
	return &BooksPostgres{db: db}
}

func (r *BooksPostgres) Create(book models.Book) error {

	err := r.db.Create(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BooksPostgres) GetAll() ([]models.Book, error) {

	var books []models.Book

	err := r.db.Find(&books).Error

	if err != nil {
		return books, err
	}

	return books, nil
}

func (r *BooksPostgres) GetById(bookId string) (models.Book, error) {

	var book models.Book

	err := r.db.Find(&book, bookId).Error

	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}

func (r *BooksPostgres) Delete(bookId string) error {

	var book models.Book

	if err := r.db.Find(&book, bookId).Error; err != nil {
		return err
	}

	err := r.db.Delete(&book, bookId).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BooksPostgres) Update(bookId string, input models.Book) error {

	var book models.UpdateBook

	book.Author = input.Author
	book.Link = input.Link
	book.Title = input.Title
	book.Pages = input.Pages
	book.Year = input.Year

	err := r.db.Model(models.Book{}).Where("id = ?", bookId).Updates(book).Error

	if err != nil {
		return err
	}

	return nil
}
