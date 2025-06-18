package service

import (
	"fmt"
	"order-sys/models"
	"order-sys/repository"

	"github.com/google/uuid"
)

func GetMenu() (*models.Menu, error) {
	items, err := repository.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	return &models.Menu{
		Items: items,
	}, nil
}

func GetMenuItemByID(ID string) (*models.MenuItem, error) {
	items, err := repository.GetAllMenuItems()
	if err != nil {
		return nil, err
	}

	for _, item := range items {
		if item.ID == ID {
			return &item, nil
		}
	}

	return nil, fmt.Errorf("menu item not found")
}

func AddMenuItem(name string, price float64, category string) (*models.MenuItem, error) {
	item := models.MenuItem{
		ID:       uuid.New().String(),
		Name:     name,
		Price:    price,
		Category: category,
	}
	err := repository.AddMenuItem(item)
	if err != nil {
		return nil, err
	}
	return &item, nil
}
