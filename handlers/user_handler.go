package handlers

import (
	"myapi/db"
	"myapi/models"
	"net/http"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	if result := db.DB.Find(&users); result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	SuccessResponse(w, users)
}