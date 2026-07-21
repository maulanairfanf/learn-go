package handlers

import (
	"myapi/models"
	"myapi/services"

	"github.com/gin-gonic/gin"
)

var authService = services.NewAuthService()

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,min=6"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func Login(c *gin.Context) {
	var loginReq LoginRequest
	if err := c.ShouldBindJSON(&loginReq); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	user, err := authService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		ErrorResponse(c, 401, err.Error())
		return
	}

	token, err := authService.GenerateToken(user.ID)
	if err != nil {
		ErrorResponse(c, 500, "Failed to generate token")
		return
	}
	SuccessResponse(c, LoginResponse{Token: token})
}

func Register(c *gin.Context) {
	var registerReq RegisterRequest
	if err := c.ShouldBindJSON(&registerReq); err != nil {
		ErrorResponse(c, 400, "Invalid request payload")
		return
	}

	user := models.User{
		Username: registerReq.Username,
		Password: registerReq.Password,
		Email:    registerReq.Email,
	}

	_, err := authService.Register(user)
	if err != nil {
		ErrorResponse(c, 500, err.Error())
		return
	}

	SuccessResponse(c, "Success Registration")
}
