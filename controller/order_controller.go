package controller

import (
	"net/http"
	"order-sys/models"
	"order-sys/service"
	"time"

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
	Status models.DishStatus `json:"status" binding:"required"`
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

	orderItems := make([]models.OrderItem, 0, len(req.Items))
	for _, item := range req.Items {
		menuItem, err := service.GetMenuItemByID(item.MenuItemID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu item ID"})
			return
		}

		orderItems = append(orderItems, models.OrderItem{
			MenuItemID: menuItem.ID,
			Name:       menuItem.Name,
			Price:      menuItem.Price,
			Quantity:   item.Quantity,
			Status:     models.DishStatusPreparing,
			UpdatedAt:  time.Now(),
		})
	}

	order, err := service.CreateOrder(userID, orderItems)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, order)
}

func GetUserAllOrders(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orders, err := service.GetUserAllOrders(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func GetUserOrderByID(c *gin.Context) {
	userID := c.GetString("userID")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orderID := c.Param("id")
	order, err := service.GetOrderByID(orderID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if order.UserId != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	c.JSON(http.StatusOK, order)
}

// -----------------------------------------------------------------------------------------

func GetPendingOrders(c *gin.Context) {
	userID := c.GetString("userID")
	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.Role != "chef" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	orders, err := service.GetPendingOrders()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func UpdateDishStatus(c *gin.Context) {
	userID := c.GetString("userID")
	user, err := service.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if user.Role != "chef" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
		return
	}

	orderID := c.Param("id")
	menuItemID := c.Param("itemId")

	var req UpdateDishStatusRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = service.UpdateDishStatus(orderID, menuItemID, req.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dish status updated"})
}
