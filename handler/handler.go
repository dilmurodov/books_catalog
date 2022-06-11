package handler

import (
	"github.com/dilmurodov/Book_catalog/service"
	"github.com/gin-gonic/gin"

	_ "github.com/dilmurodov/Book_catalog/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger(), gin.Recovery())

	// router.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{
		category := api.Group("/categories")
		{
			category.POST("/", h.createCategory)
			category.GET("/", h.getAllCategory)
			category.GET("/:id", h.getCategoryById)
			category.PATCH("/:id", h.updateCategory)
			category.DELETE("/:id", h.deleteCategory)
		}
	}
	{
		books := api.Group("/books")
		{
			books.POST("/", h.createBook)
			books.GET("/", h.getAllBooks)
			books.GET("/:id", h.getBookById)
			books.PATCH("/:id", h.updateBook)
			books.DELETE("/:id", h.deleteBook)
		}
	}

	return router
}
