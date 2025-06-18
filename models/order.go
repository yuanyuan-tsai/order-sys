package models

import "time"

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "Pending"
	OrderStatusCompleted OrderStatus = "Completed"
)

type DishStatus string

const (
	DishStatusPreparing DishStatus = "Preparing"
	DishStatusReady     DishStatus = "Ready"
)

type OrderItem struct {
	MenuItemID string
	Name       string
	Price      float64
	Quantity   int
	Status     DishStatus
	UpdatedAt  time.Time
}

type Order struct {
	ID        string
	UserId    string
	Items     []OrderItem
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt time.Time
	Total     float64
}
