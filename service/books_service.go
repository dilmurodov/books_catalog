package service

import (
	"github.com/dilmurodov/Book_catalog/models"
	"github.com/dilmurodov/Book_catalog/repository"
)

type BooksService struct {
	repo repository.Books
}

func NewBooksService(repo repository.Books) *BooksService {
	return &BooksService{repo: repo}
}

func (r *BooksService) Create(book models.Book) error {

	err := r.repo.Create(book)

	if err != nil {
		return err
	}

	return nil
}

func (r *BooksService) GetAll() ([]models.Book, error) {

	books, err := r.repo.GetAll()

	if err != nil {
		return []models.Book{}, err
	}

	return books, nil
}

func (r *BooksService) GetById(bookId string) (models.Book, error) {

	books, err := r.repo.GetById(bookId)

	if err != nil {
		return models.Book{}, err
	}

	return books, nil
}

func (r *BooksService) Delete(bookId string) error {

	err := r.repo.Delete(bookId)

	if err != nil {
		return err
	}

	return nil
}

func (r *BooksService) Update(bookId string, input models.Book) error {

	err := r.repo.Update(bookId, input)

	if err != nil {
		return err
	}

	return nil
}
