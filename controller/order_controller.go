package controller

import (
	"net/http"
	"order-sys/models"

	"github.com/gin-gonic/gin"
)

type CreateOrderRequest struct {
	Items []OrderItemRequest `json:"items" binding:"required"`
}

type OrderItemRequest struct {
	MenuItemID string `json:"menu_item_id" binding:"required"`
	Quantity   int    `json:"quantity" binding:"required"`
}

type UpdateDishStatusRequest struct {
	OrderID string `json:"order_id" binding:"required"`
	DishID  string `json:"dish_id" binding:"required"`
}

func CreateOrder(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		return
	}

	var req CreateOrderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderItem := make([]models.OrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		orderItem = append(orderItem, models.OrderItem{
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
		})
	}
}
