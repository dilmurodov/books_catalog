package repository

import (
	"fmt"

	"github.com/dilmurodov/Book_catalog/models"
	"github.com/jinzhu/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.Book{})
	db.AutoMigrate(&models.BooksCategory{})

	return db, nil
}
