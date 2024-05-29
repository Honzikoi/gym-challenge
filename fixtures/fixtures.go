package fixtures

import (
	"log"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"golang.org/x/crypto/bcrypt"
)

// CreateUserFixture creates a user fixture for testing
func CreateUserFixture() {
	password, err := bcrypt.GenerateFromPassword([]byte("test"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Failed to hash password: ", err)
	}

	user := models.Users{
		Username: "test",
		Email:    "test@test.test",
		Password: string(password),
	}

	result := database.DB.Db.Create(&user)
	if result.Error != nil {
		log.Fatal("Failed to create user fixture: ", result.Error)
	}

	log.Println("User fixture created successfully")
}
