package controller

import (
	"net/http"
	"order-sys/service"

	"github.com/gin-gonic/gin"
)

type MenuItemRequest struct {
	Name     string  `json:"name" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Category string  `json:"category" binding:"required"`
}

func GetMenu(c *gin.Context) {
	menu, err := service.GetMenu()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, menu)
}

func AddMenuItem(c *gin.Context) {
	var req MenuItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	item, err := service.AddMenuItem(req.Name, req.Price, req.Category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, item)
}
