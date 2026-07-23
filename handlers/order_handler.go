package handlers

import (
	"myapi/models"
	"myapi/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var orderService = services.NewOrderService()

func GetOrders(c *gin.Context) {
	orders, err := orderService.GetAll()
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  500,
			Message: err.Error(),
		})
		return
	}

	responses := make([]models.OrderResponse, len(orders))
	for i, order := range orders {
		responses[i] = order.ToResponse()
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    responses,
		Message: "Success Get Orders",
	})
}

func GetOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  400,
			Message: "Invalid Order ID",
		})
		return
	}

	order, err := orderService.GetByID(id)
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  404,
			Message: "Order Not Found",
		})
		return
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    order.ToResponse(),
		Message: "Success Get Order",
	})
}

func CreateOrder(c *gin.Context) {
	var req models.OrderPayload
	if err := c.ShouldBindJSON(&req); err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusBadRequest,
			Message: "Invalid request payload",
		})
		return
	}

	userID, _ := c.Get("userID")

	order, err := orderService.Create(req, userID.(uint))
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    order.ToResponse(),
		Message: "Success Create Order",
	})
}

func PayOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusBadRequest,
			Message: "Invalid order ID",
		})
		return
	}

	order, err := orderService.Pay(id)
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    order,
		Message: "Success update to paid",
	})
}

func CancelOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusBadRequest,
			Message: "Invalid order ID",
		})
		return
	}

	order, err := orderService.Cancel(id)
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    order,
		Message: "Success update to cancel",
	})
}

func DeleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusBadRequest,
			Message: "Invalid Order ID",
		})
		return
	}

	order, err := orderService.Delete(id)
	if err != nil {
		ErrorResponse(ErrorParams{
			C:       c,
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
		return
	}

	SuccessResponse(SuccessParams{
		C:       c,
		Data:    order.ToResponse(),
		Message: "Success delete order",
	})
}
