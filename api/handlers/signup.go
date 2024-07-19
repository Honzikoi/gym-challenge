package handlers

import (
	"regexp"

	"github.com/Honzikoi/gym-challenge/internal/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// isEmailValid checks if the provided email address is valid.
func isEmailValid(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// SignUp godoc
//
//	@Summary		Sign up a new user
//	@Description	Create a new user with email and password
//	@Tags			Signup & Login
//	@Accept			json
//	@Produce		json
//	@Param			signup	body		SignUpRequest	true	"Sign up request"
//	@Success		201		{string}	User			created	successfully
//	@Failure		400		{string}	Invalid			request	payload
//	@Failure		500		{string}	Failed			to		create	user
//	@Router			/signup [post]
func SignUp(c *fiber.Ctx) error {
	db := database.DB.Db

	var req SignUpRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	// Validate email format
	if !isEmailValid(req.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid email address",
		})
	}

	// Check if user already exists
	var existingUser models.Users
	if err := db.Where("email = ?", req.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "User already exists",
		})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to hash password",
		})
	}

	// Create the user
	user := models.Users{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	if err := db.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create user",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created successfully",
	})
}
