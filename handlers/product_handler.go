package handlers

import (
	"errors"
	"myapi/models"
	"myapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var productService = services.NewProductService()

func GetProducts(c *gin.Context) {
	products, err := productService.GetAll()
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, products)
}

func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	product, err := productService.GetByID(id)
	if err != nil {
		ErrorResponse(c, 404, "Product not found")
		return
	}
	SuccessResponse(c, product)
}

func CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	product, err := productService.Create(req)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, product)
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	product, err := productService.Update(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(c, 404, "Product not found")
		} else {
			ErrorResponse(c, 500, err.Error())
		}
		return
	}
	SuccessResponse(c, product)
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid product ID")
		return
	}

	if err := productService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(c, 404, "Product not found")
		} else {
			ErrorResponse(c, 500, err.Error())
		}
		return
	}
	SuccessResponse(c, "Product deleted successfully")
}
