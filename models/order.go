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
	Name       string     `json:"name"`
	Price      float64    `json:"price"`
	Quantity   int        `json:"quantity"`
	Status     DishStatus `json:"status"`
	UpdatedAt  time.Time
}

type Order struct {
	ID        string
	UserId    string
	Items     []OrderItem `json:"items"`
	Status    OrderStatus `json:"status"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Total     float64
}
