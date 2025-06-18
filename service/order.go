package service

import (
	"order-sys/models"
	"order-sys/repository"
	"time"

	"github.com/google/uuid"
)

func CreateOrder(userID string, items []models.OrderItem) error {
	var total float64
	for _, item := range items {
		total += item.Price * float64(item.Quantity)
	}

	order := models.Order{
		ID:        uuid.New().String(),
		UserId:    userID,
		Items:     items,
		Total:     total,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    models.OrderStatusPending,
	}

	err := repository.SaveOrder(order)
	if err != nil {
		return nil
	}

	// go processOrder(order.ID)

	return nil
}

func GetOrder(orderID string) (*models.Order, error) {
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}
	return order, err
}

func GetUserAllOrders(userID string) ([]models.Order, error) {
	orderIDs, err := repository.GetUserOrderIDs(userID)
	if err != nil {
		return nil, err
	}

	orders := []models.Order{}
	for _, orderID := range orderIDs {
		order, err := GetOrder(orderID)
		if err != nil {
			return nil, err
		}
		orders = append(orders, *order)
	}

	return orders, nil
}

func GetPreparingOrders() ([]models.Order, error) {
	return nil, nil
}
