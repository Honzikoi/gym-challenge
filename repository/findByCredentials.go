package repository

import (
	"errors"

	"github.com/Honzikoi/gym-challenge/models"
)

// Simulate a database call
func FindByCredentials(email, password string) (*models.Users, error) {
	// Here you would query your database for the user with the given email
	if email == "test@mail.com" && password == "test12345" {
		return &models.Users{
			Email:    "test@mail.com",
			Password: "test12345",
		}, nil
	}
	return nil, errors.New("user not found")
}
