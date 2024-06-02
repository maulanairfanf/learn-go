package handlers

import (
	"encoding/json"
	"myapi/db"
	"myapi/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// APIResponse represents a standard API response format
type APIResponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

// SuccessResponse creates a successful API response with data
func SuccessResponse(w http.ResponseWriter, data interface{}) {
	response := APIResponse{
		Status: http.StatusOK,
		Data:   data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// ErrorResponse creates an error API response with a message
func ErrorResponse(w http.ResponseWriter, status int, message string) {
	response := APIResponse{
		Status: status,
		Data: map[string]string{
			"error": message,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    if result := db.DB.Find(&users); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// Return the product as a successful API response
	SuccessResponse(w, users)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
    var products []models.Product
    if result := db.DB.Find(&products); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// Return the product as a successful API response
	SuccessResponse(w, products)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
    // Get the product ID from the request URL parameters
    vars := mux.Vars(r)
    id, err := strconv.Atoi(vars["id"])
    if err != nil {
        http.Error(w, "Invalid product ID", http.StatusBadRequest)
        return
    }

    // Query the database for the product with the given ID
    var product models.Product
    if result := db.DB.First(&product, id); result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
            http.NotFound(w, r)
            return
        }
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

	// Return the product as a successful API response
	SuccessResponse(w, product)
}