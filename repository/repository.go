package repository

import (
	"github.com/dilmurodov/Book_catalog/models"
	"github.com/jinzhu/gorm"
)

type Books interface {
	Create(book models.Book) error
	GetAll() ([]models.Book, error)
	GetById(bookId string) (models.Book, error)
	Delete(bookId string) error
	Update(bookId string, input models.Book) error
}

type BooksCategory interface {
	Create(category models.BooksCategory) (int, error)
	GetAll() ([]models.BooksCategory, error)
	GetById(categoryId string) (models.BooksCategory, error)
	Delete(categoryId string) error
	Update(categoryId string, input models.BooksCategory) error
}

type Repository struct {
	Books
	BooksCategory
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Books:         NewBooksPostgres(db),
		BooksCategory: NewCategoryPostgres(db),
	}
}
