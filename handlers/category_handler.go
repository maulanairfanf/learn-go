package handlers

import (
	"errors"
	"myapi/models"
	"myapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var categoryService = services.NewCategoryService()

// GetCategories handles retrieving all categories
func GetCategories(c *gin.Context) {
	categories, err := categoryService.GetAll()
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: categories, Message: "success"})
}

// GetCategory handles retrieving a single category by ID
func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid Category ID"})
		return
	}

	category, err := categoryService.GetByID(id)
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Category not found"})
		return
	}

	SuccessResponse(SuccessParams{C: c, Data: category, Message: "success"})
}

// CreateCategory handles the creation of a new category
func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid request payload"})
		return
	}

	category, err := categoryService.Create(req)
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: category, Message: "success"})
}

// UpdateCategory handles the update of a category
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid Category ID"})
		return
	}

	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid request payload"})
		return
	}

	category, err := categoryService.Update(id, req)
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}

	SuccessResponse(SuccessParams{C: c, Data: category, Message: "success"})
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid Category Id"})
		return
	}

	_, err = categoryService.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Category not found"})
			return
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
			return
		}
	}
	SuccessResponse(SuccessParams{C: c, Data: "Category deleted successfully", Message: "success"})
}
