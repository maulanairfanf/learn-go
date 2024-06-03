package handlers

import (
	"encoding/json"
	"myapi/db"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCategories handles retrieving all categories
func GetCategories(w http.ResponseWriter, r *http.Request) {
    var categories []models.Category

    if err := db.DB.Find(&categories).Error; err != nil {
        ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    SuccessResponse(w, categories)
}

// GetCategory handles retrieving a single category by ID
func GetCategory(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        ErrorResponse(w, http.StatusBadRequest, "Invalid category ID")
        return
    }

    var category models.Category
    if err := db.DB.First(&category, id).Error; err != nil {
        ErrorResponse(w, http.StatusNotFound, "Category not found")
        return
    }

    SuccessResponse(w, category)
}

// CreateCategory handles the creation of a new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category

    // Decode the JSON request body into the category struct
    if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
        ErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
        return
    }

    // Create the category in the database
    if err := db.DB.Create(&category).Error; err != nil {
        ErrorResponse(w, http.StatusInternalServerError, err.Error())
        return
    }

    SuccessResponse(w, category)
}
