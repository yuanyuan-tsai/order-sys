package service

import (
	"fmt"
	"order-sys/models"
	"order-sys/repository"

	"github.com/google/uuid"
)

func RegisterUser(name, email, password string) (*models.User, error) {
	existingUser, _ := repository.GetUserByEmail(email)
	if existingUser != nil {
		return nil, fmt.Errorf("email already registered")
	}

	user := models.User{
		ID:       uuid.New().String(),
		Name:     name,
		Email:    email,
		Password: password,
		Role:     "customer",
	}

	err := repository.SaveUser(user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func Login(email, password string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}

	if user.Password != password {
		return nil, fmt.Errorf("invalid email or password")
	}

	return user, nil
}

func GetUserByID(ID string) (*models.User, error) {
	return repository.GetUserByID(ID)
}
