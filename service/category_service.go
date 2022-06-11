package service

import (
	"github.com/dilmurodov/Book_catalog/models"
	"github.com/dilmurodov/Book_catalog/repository"
)

type CategoryService struct {
	repoCategory repository.BooksCategory
}

func NewCategoryService(repoCategory repository.BooksCategory) *CategoryService {
	return &CategoryService{repoCategory: repoCategory}
}

func (r *CategoryService) Create(category models.BooksCategory) (int, error) {

	_, err := r.repoCategory.Create(category)

	if err != nil {
		return 0, err
	}

	return 0, nil
}

func (r *CategoryService) GetAll() ([]models.BooksCategory, error) {

	entities, err := r.repoCategory.GetAll()

	if err != nil {
		return []models.BooksCategory{}, err
	}

	return entities, nil
}

func (r *CategoryService) GetById(categoryId string) (models.BooksCategory, error) {

	var entity models.BooksCategory

	entity, err := r.repoCategory.GetById(categoryId)

	if err != nil {
		return models.BooksCategory{}, err
	}

	return entity, nil
}

func (r *CategoryService) Delete(categoryId string) error {

	err := r.repoCategory.Delete(categoryId)

	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryService) Update(categoryId string, input models.BooksCategory) error {

	err := r.repoCategory.Update(categoryId, input)

	if err != nil {
		return err
	}

	return nil
}
