package handler

import (
	"fmt"

	"github.com/dilmurodov/Book_catalog/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

// @Summary Creat Category
// @Description create category
// @Accept json
// @Produce json
// @Tags categories
// @ID creat-category
// @Param input body models.CreateCategory false "category"
// @Success 201 {object} map[string]interface{}
// @Failure 500 {object}  map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/categories [post]
func (h *Handler) createCategory(c *gin.Context) {

	var input models.BooksCategory

	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	_, err = h.services.BooksCategory.Create(input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "new book category successfully added!",
	})
}

// @Summary Get All Category
// @Description get categories
// @Accept json
// @Produce json
// @Tags categories
// @ID get-categories
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/categories [get]
func (h *Handler) getAllCategory(c *gin.Context) {
	entities, err := h.services.BooksCategory.GetAll()

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

// @Summary Get Category by Id
// @Description get category by id
// @Accept json
// @Produce json
// @Tags categories
// @ID get-category
// @Param id path string true "category _"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object}  map[string]interface{}
// @Router /api/categories/{id} [get]
func (h *Handler) getCategoryById(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	entity, err := h.services.BooksCategory.GetById(id)

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

// @Summary Update Category by Id
// @Description update category by id
// @Accept json
// @Produce json
// @Tags categories
// @ID update-category
// @Param id path string true "category_id"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Param input body models.UpdateCategory true "category"
// @Router /api/categories/{id} [patch]
func (h *Handler) updateCategory(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	var category models.BooksCategory

	c.ShouldBindJSON(&category)

	err := h.services.BooksCategory.Update(id, category)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	err = h.services.BooksCategory.Update(id, category)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "category successfully updated!",
	})
}

// @Summary delete Category by Id
// @Description delete category by id
// @Accept json
// @Produce json
// @Tags categories
// @ID delete-category
// @Param id path string true "category_id"
// @Success 202 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/categories/{id} [delete]
func (h *Handler) deleteCategory(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)

	err := h.services.BooksCategory.Delete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"status":  "success",
		"message": "category successfully deleted!",
	})
}
