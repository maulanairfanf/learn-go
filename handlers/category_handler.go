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
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, categories)
}

// GetCategory handles retrieving a single category by ID
func GetCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid Category ID")
		return
	}

	category, err := categoryService.GetByID(id)
	if err != nil {
		ErrorResponse(c, 404, "Category not found")
		return
	}

	SuccessResponse(c, category)
}

// CreateCategory handles the creation of a new category
func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	category, err := categoryService.Create(req)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, category)
}

// UpdateCategory handles the update of a category
func UpdateCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid Category ID")
		return
	}

	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	category, err := categoryService.Update(id, req)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}

	SuccessResponse(c, category)
}

func DeleteCategory(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid Category Id")
		return
	}

	_, err = categoryService.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(c, 404, "Category not found")
			return
		} else {
			ErrorResponse(c, 500, err.Error())
			return
		}
	}
	SuccessResponse(c, "Category deleted successfully")
}
