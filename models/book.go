package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model

	Title           string `gorm:"type:varchar(100);uniqueIndex;" json:"title"`
	Author          string `json:"author"`
	BooksCategoryID uint   `json:"categoryId"`
	Pages           int    `json:"pages"`
	Year            int    `json:"year"`
	Link            string `json:"link"`
}

type UpdateBook struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	BooksCategoryID int    `json:"categoryId"`
	Pages           int    `json:"pages"`
	Year            int    `json:"year"`
	Link            string `json:"link"`
}

type CreateBook struct {
	Title           string `json:"title"`
	Author          string `json:"author"`
	BooksCategoryID int    `json:"categoryId"`
	Pages           int    `json:"pages"`
	Year            int    `json:"year"`
	Link            string `json:"link"`
}
