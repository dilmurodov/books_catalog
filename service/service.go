package service

import (
	"github.com/dilmurodov/Book_catalog/models"
	"github.com/dilmurodov/Book_catalog/repository"
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

type Service struct {
	Books
	BooksCategory
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Books:         NewBooksService(repos.Books),
		BooksCategory: NewCategoryService(repos.BooksCategory),
	}
}
