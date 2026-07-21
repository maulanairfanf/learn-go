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
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: products, Message: "success"})
}

func GetProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid product ID"})
		return
	}

	product, err := productService.GetByID(id)
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Product not found"})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: product, Message: "success"})
}

func CreateProduct(c *gin.Context) {
	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid request payload"})
		return
	}

	product, err := productService.Create(req)
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: product, Message: "success"})
}

func UpdateProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid product ID"})
		return
	}

	var req models.CreateProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid request payload"})
		return
	}

	product, err := productService.Update(id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Product not found"})
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		}
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: product, Message: "success"})
}

func DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid product ID"})
		return
	}

	if err := productService.Delete(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "Product not found"})
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		}
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: "Product deleted successfully", Message: "success"})
}
