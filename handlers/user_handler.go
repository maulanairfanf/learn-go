package handlers

import (
	"myapi/models"
	"myapi/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

var userService = services.NewUserService()

func GetUsers(c *gin.Context) {
	users, err := userService.GetAll()
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, users)
}

func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid User ID")
		return
	}

	user, err := userService.GetByID(id)
	if err != nil {
		ErrorResponse(c, 404, "User not found")
		return
	}
	SuccessResponse(c, user)
}

func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid User ID")
		return
	}

	user, err := userService.Delete(id)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, user)
}

func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(c, 400, "Invalid User ID")
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(c, 400, "Invalid Payload")
		return
	}

	user, err := userService.Update(id, req)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}
	SuccessResponse(c, user)
}
