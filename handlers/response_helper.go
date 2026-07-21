package handlers

import (
	"github.com/gin-gonic/gin"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type SuccessParams struct {
	C       *gin.Context
	Data    interface{}
	Message string
}

type ErrorParams struct {
	C       *gin.Context
	Status  int
	Message string
}

func SuccessResponse(p SuccessParams) {
	p.C.JSON(200, APIResponse{
		Status:  200,
		Data:    p.Data,
		Message: p.Message,
	})
}

func ErrorResponse(p ErrorParams) {
	p.C.JSON(p.Status, APIResponse{
		Status:  p.Status,
		Data:    nil,
		Message: p.Message,
	})
}
