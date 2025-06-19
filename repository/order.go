package repository

import (
	"encoding/json"
	"fmt"
	"order-sys/models"
	"time"
)

func SaveOrder(order models.Order) error {
	orderJSON, err := json.Marshal(order)
	if err != nil {
		return err
	}

	err = RedisClient.Set(Ctx, fmt.Sprintf("order:%s", order.ID), orderJSON, 0).Err()
	if err != nil {
		return err
	}

	err = RedisClient.LPush(Ctx, fmt.Sprintf("user:%s:orders", order.UserId), order.ID).Err()
	if err != nil {
		return err
	}

	existingOrder, err := GetOrderByID(order.ID)
	isNewOrder := err != nil
	if isNewOrder {
		err = RedisClient.LPush(Ctx, fmt.Sprintf("user:%s:orders", order.UserId), order.ID).Err()
		if err != nil {
			return err
		}

		if order.Status == models.OrderStatusPending {
			err = RedisClient.LPush(Ctx, "pending_orders", order.ID).Err()
			if err != nil {
				return err
			}
		}
	} else if existingOrder.Status != models.OrderStatusCompleted && order.Status == models.OrderStatusCompleted {
		err = RedisClient.LRem(Ctx, "pending_orders", 0, order.ID).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func GetUserOrderIDs(userID string) ([]string, error) {
	orders, err := RedisClient.LRange(Ctx, fmt.Sprintf("user:%s:orders", userID), 0, -1).Result()
	if err != nil {
		return nil, err
	}

	var orderIDs []string
	for _, orderID := range orders {
		orderIDs = append(orderIDs, orderID)
	}

	return orderIDs, nil
}

func GetOrderByID(ID string) (*models.Order, error) {
	result, err := RedisClient.Get(Ctx, fmt.Sprintf("order:%s", ID)).Result()
	if err != nil {
		return nil, err
	}

	var order models.Order
	if err := json.Unmarshal([]byte(result), &order); err != nil {
		return nil, err
	}

	return &order, nil
}

func GetPendingOrders() ([]models.Order, error) {
	orderIDs, err := RedisClient.LRange(Ctx, "pending_orders", 0, -1).Result()
	if err != nil {
		return nil, err
	}

	orders := make([]models.Order, 0, len(orderIDs))
	for _, id := range orderIDs {
		order, err := GetOrderByID(id)
		if err != nil {
			return nil, err
		}
		if order.Status != models.OrderStatusCompleted {
			orders = append(orders, *order)
		}
	}

	return orders, nil
}

func UpdateOrderStatus(orderID string, status models.OrderStatus) error {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return err
	}

	order.Status = status
	order.UpdatedAt = time.Now()

	return SaveOrder(*order)
}

func UpdateDishStatus(orderID string, dishID string, status models.DishStatus) error {
	order, err := GetOrderByID(orderID)
	if err != nil {
		return err
	}

	var allReady = true
	for _, dish := range order.Items {
		if dish.MenuItemID == dishID {
			dish.Status = status
			order.UpdatedAt = time.Now()
		}

		if dish.Status != models.DishStatusReady {
			allReady = false
		}
	}

	if allReady {
		order.Status = models.OrderStatusCompleted
	}

	order.UpdatedAt = time.Now()
	return SaveOrder(*order)
}
