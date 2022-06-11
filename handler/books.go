package handler

import (
	"github.com/dilmurodov/Book_catalog/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

// @Summary Creat Book
// @Description create book
// @Accept json
// @Produce json
// @Tags books
// @ID creat-book
// @Param input body models.CreateBook false "book"
// @Success 201 {object} map[string]interface{}
// @Failure 500 {object}  map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/books [post]
func (h *Handler) createBook(c *gin.Context) {

	var input models.Book

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err = h.services.Books.Create(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "new book successfully added!",
	})

}

// @Summary Get All Books
// @Description get books
// @Accept json
// @Produce json
// @Tags books
// @ID get-books
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/books [get]
func (h *Handler) getAllBooks(c *gin.Context) {
	entities, err := h.services.Books.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": "error",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   entities,
	})

}

// @Summary Get books by Id
// @Description get books by id
// @Accept json
// @Produce json
// @Tags books
// @ID get-book
// @Param id path string true "book_id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/books/{id} [get]
func (h *Handler) getBookById(c *gin.Context) {
	id := c.Param("id")
	entity, err := h.services.Books.GetById(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   entity,
	})
}

// @Summary Update Book by Id
// @Description update book by id
// @Accept json
// @Produce json
// @Tags books
// @ID update-book
// @Param id path string true "books_id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Param input body models.UpdateBook true "book"
// @Router /api/books/{id} [patch]
func (h *Handler) updateBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	c.ShouldBindJSON(&book)

	err := h.services.Books.Update(id, book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err = h.services.Books.Update(id, book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Book successfully updated!",
	})
}

// @Summary delete Book by Id
// @Description delete book by id
// @Accept json
// @Produce json
// @Tags books
// @ID delete-book
// @Param id path string true "book_id"
// @Success 202 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/books/{id} [delete]
func (h *Handler) deleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.services.Books.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "Book successfully deleted!",
	})
}
