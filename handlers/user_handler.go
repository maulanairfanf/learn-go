package handlers

import (
	"errors"
	"myapi/models"
	"myapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var userService = services.NewUserService()

func GetUsers(c *gin.Context) {
	users, err := userService.GetAll()
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: users, Message: "success"})
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid User ID"})
		return
	}

	user, err := userService.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "User not found"})
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		}
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: user, Message: "success"})
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid User ID"})
		return
	}

	user, err := userService.Delete(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "User not found"})
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		}
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: user, Message: "success"})
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid User ID"})
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{C: c, Status: 400, Message: "Invalid Payload"})
		return
	}

	user, err := userService.Update(id, req)
	if err != nil {
		if errors.Is(err, services.ErrConflict) {
			ErrorResponse(ErrorParams{C: c, Status: 409, Message: "Username already taken"})
		} else if errors.Is(err, gorm.ErrRecordNotFound) {
			ErrorResponse(ErrorParams{C: c, Status: 404, Message: "User not found"})
		} else {
			ErrorResponse(ErrorParams{C: c, Status: 500, Message: err.Error()})
		}
		return
	}
	SuccessResponse(SuccessParams{C: c, Data: user, Message: "success"})
}
