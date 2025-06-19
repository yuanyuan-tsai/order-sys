package service

import (
	"fmt"
	"order-sys/models"
	"order-sys/repository"
	"time"
)

func StartKitchen(chefCount int) {
	for i := 0; i < chefCount; i++ {
		go chef(i)
	}
}

type DishTask struct {
	OrderID    string
	MenuItemID string
	Quantity   int
	Name       string
}

var kitchenChannel = make(chan DishTask, 100)

func processOrder(orderID string) {
	order, err := repository.GetOrderByID(orderID)
	if err != nil {
		fmt.Printf("Error Getting Order:%v\n", err)
		return
	}

	for _, item := range order.Items {
		kitchenChannel <- DishTask{
			OrderID:    orderID,
			MenuItemID: item.MenuItemID,
			Quantity:   item.Quantity,
			Name:       item.Name,
		}
	}
}

func chef(chefID int) {
	for task := range kitchenChannel {
		fmt.Printf("Chef %d is preparing %s for order %s\n", chefID, task.Name, task.OrderID)
		err := repository.UpdateDishStatus(task.OrderID, task.MenuItemID, models.DishStatusPreparing)
		if err != nil {
			fmt.Printf("Error updating dish status: %v\n", err)
			continue
		}
		// Process item in the order
		time.Sleep(time.Duration(task.Quantity) * time.Second)
		// Complete a specific dish
		err = repository.UpdateDishStatus(task.OrderID, task.MenuItemID, models.DishStatusReady)
		if err != nil {
			fmt.Printf("Error updating dish status: %v\n", err)
			continue
		}

		fmt.Printf("Chef %d finished preparing %s for order %s\n", chefID, task.Name, task.OrderID)

	}
}
