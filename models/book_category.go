package models

import "github.com/jinzhu/gorm"

type BooksCategory struct {
	gorm.Model

	CategoryName string `json:"categoryname" gorm:"typevarchar(100);uniqueIndex"`
	Books        []Book `json:"books" gorm:"foreignKey:BooksCategoryID;constraint:OnDelete:CASCADE;OnUpdate:CASCADE;"`
}

type UpdateCategory struct {
	CategoryName string `json:"categoryname" gorm:"typevarchar(100);"`
}

type CreateCategory struct {
	CategoryName string `json:"categoryname" gorm:"typevarchar(100);"`
}