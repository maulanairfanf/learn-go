package handlers

import (
	"encoding/json"
	"myapi/db"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetProducts handles retrieving all products along with their categories
func GetProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product

    if err := db.DB.Preload("Categories").Find(&products).Error; err != nil {
        ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    SuccessResponse(w, products)
}

// GetProduct handles retrieving a single product by ID along with its categories
func GetProduct(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        ErrorResponse(w, http.StatusBadRequest, "Invalid product ID")
        return
    }

    var product models.Product
    if err := db.DB.Preload("Categories").First(&product, id).Error; err != nil {
        ErrorResponse(w, http.StatusNotFound, "Product not found")
        return
    }

    SuccessResponse(w, product)
}

// CreateProduct handles the creation of a new product
func CreateProduct(w http.ResponseWriter, r *http.Request) {
    var product models.Product

    // Decode the JSON request body into the product struct
    if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
        ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Create the product in the database
    if err := db.DB.Create(&product).Error; err != nil {
        ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    SuccessResponse(w, product)
}
