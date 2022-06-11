package repository

import (
	"errors"

	"github.com/dilmurodov/Book_catalog/models"
	"github.com/jinzhu/gorm"
)

type CategoryPostgres struct {
	db *gorm.DB
}

func NewCategoryPostgres(db *gorm.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (r *CategoryPostgres) Create(category models.BooksCategory) (int, error) {

	var ctgs models.BooksCategory

	if err := r.db.Find(&ctgs, category.CategoryName); err == nil {
		return 1, errors.New("with this ID Category already exists")
	}

	err := r.db.Create(&category).Error

	if err != nil {
		return 1, err
	}

	return 0, nil
}

func (r *CategoryPostgres) GetAll() ([]models.BooksCategory, error) {
	var result []models.BooksCategory

	err := r.db.Set("gorm:auto_preload", true).Find(&result).Error

	if err != nil {
		return result, err
	}

	return result, nil
}

func (r *CategoryPostgres) GetById(categoryId string) (models.BooksCategory, error) {

	var category models.BooksCategory

	err := r.db.Find(&category, categoryId).Error

	if err != nil {
		return models.BooksCategory{}, err
	}

	r.db.Set("gorm:auto_preload", true).Find(&category, categoryId)

	return category, nil
}

func (r *CategoryPostgres) Delete(categoryId string) error {

	var category models.BooksCategory

	err := r.db.Where("id = ?", categoryId).Delete(&category).Error

	if err != nil {
		return err
	}

	err = r.db.Where("books_category_id = ?", categoryId).Delete(&models.Book{}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *CategoryPostgres) Update(categoryId string, input models.BooksCategory) error {

	var category models.BooksCategory

	err := r.db.Find(&category, categoryId).Error

	if err != nil {
		return err
	}

	category.CategoryName = input.CategoryName

	err = r.db.Save(&category).Error

	if err != nil {
		return err
	}

	return nil
}
