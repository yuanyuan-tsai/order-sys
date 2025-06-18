package repository

import (
	"encoding/json"
	"order-sys/models"
)

func GetAllMenuItems() ([]models.MenuItem, error) {
	result, err := RedisClient.HGetAll(Ctx, "menu_items").Result()
	if err != nil {
		return nil, err
	}

	items := make([]models.MenuItem, 0, len(result))
	for _, v := range result {
		var item models.MenuItem
		if err := json.Unmarshal([]byte(v), &item); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

func AddMenuItem(item models.MenuItem) error {
	itemJSON, err := json.Marshal(item)
	if err != nil {
		return err
	}

	err = RedisClient.HSet(Ctx, "menu_items", item.ID, itemJSON).Err()
	if err != nil {
		return err
	}

	return nil
}
