package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dilmurodov/Book_catalog/handler"
	"github.com/dilmurodov/Book_catalog/repository"
	"github.com/dilmurodov/Book_catalog/service"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	rest "github.com/dilmurodov/Book_catalog/server"
	
)

func main() {

	// @title           Books Catalog API
	// @version         1.0
	// @description     API Server to Books Catalog.

	// @host localhost:8000
	// @BasePath /
	// @schemes http
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(rest.Server)

	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	logrus.Print("Books App Started")

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("Books App Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

//======

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/dilmurodov/Book_catalog/models"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/lib/pq"
// )

// var db *gorm.DB
// var err error

// var (
// 	book_catalog = &models.BookCatalog{
// 		CatalogName: "Literature",
// 	}
// 	books = []models.Book{
// 		{Name: "Things Fall Apart", Author: "Chinua Achebe", Page: 209, Year: 1958, Link: "https://en.wikipedia.org/wiki/Things_Fall_Apart", BookCatalogID: 1},
// 		{Name: "The Divine Comedy", Author: "Dante Alighieri", Page: 928, Year: 1836, Link: "https://en.wikipedia.org/wiki/Fairy_Tales_Told_for_Children._First_Collection.", BookCatalogID: 1},
// 		{Name: "Fairy tales", Author: "Hans Christian Andersen", Page: 784, Year: 1315, Link: "https://en.wikipedia.org/wiki/Divine_Comedy", BookCatalogID: 1},
// 	}
// )

// func main() {

// 	dialect := os.Getenv("DIALECT")
// 	host := os.Getenv("DB_HOST")
// 	port := os.Getenv("DB_PORT")
// 	user := os.Getenv("USER")
// 	name := os.Getenv("NAME")
// 	password := os.Getenv("PASSWORD")

// 	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, name, password, port)

// 	// Opening Connection to db

// 	fmt.Println(dialect)

// 	db, err = gorm.Open(dialect, dbURI)

// 	if err != nil {
// 		log.Fatal(err)
// 	} else {
// 		fmt.Println("Successfully Connected to db")
// 	}

// 	// Close connection to database

// 	defer db.Close()

// 	db.AutoMigrate(&models.BookCatalog{})
// 	db.AutoMigrate(&models.Book{})

// 	db.Create(&book_catalog)

// 	for indexBks := range books {
// 		db.Create(&books[indexBks])
// 	}
// }
