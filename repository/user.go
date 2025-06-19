package repository

import (
	"encoding/json"
	"fmt"
	"order-sys/models"
)

func SaveUser(user models.User) error {
	userJSON, err := json.Marshal(user)
	if err != nil {
		return err
	}

	err = RedisClient.Set(Ctx, fmt.Sprintf("user:%s", user.ID), userJSON, 0).Err()
	if err != nil {
		return err
	}

	err = RedisClient.Set(Ctx, fmt.Sprintf("email:%s", user.Email), user.ID, 0).Err()
	if err != nil {
		return err
	}

	return nil
}

func GetUserByID(ID string) (*models.User, error) {
	result, err := RedisClient.Get(Ctx, fmt.Sprintf("user:%s", ID)).Result()
	if err != nil {
		return nil, err
	}

	var user models.User
	if err := json.Unmarshal([]byte(result), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	userID, err := RedisClient.Get(Ctx, fmt.Sprintf("email:%s", email)).Result()
	if err != nil {
		return nil, err
	}

	return GetUserByID(userID)
}
