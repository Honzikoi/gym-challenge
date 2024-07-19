package controllers

import (
	"regexp"
	"strconv"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// isEmailValid checks if the provided email address is valid.
func isEmailValid(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

// GetUsers godoc
//
//	@Summary		Get all users
//	@Description	Retrieve all users
//	@Tags			Users
//	@Produce		json
//	@Success		200	{array}		models.Users
//	@Failure		500	{string}	Internal Server Error
//	@Router			/users [get]
func GetUsers(c *fiber.Ctx) error {
	var users []models.Users
	if err := database.DB.Db.Preload("About").Preload("Groups").Preload("Role").Find(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

// CreateUser godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.Users	true	"User"
//	@Success		201	{object}	models.Users
//	@Failure		400	{string}	Cannot parse JSON
//	@Failure		400	{string}	Invalid email address
//	@Failure		400	{string}	User already exists
//	@Failure		500	{string}	Internal Server Error
//	@Router			/users [post]
func CreateUser(c *fiber.Ctx) error {
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate email format
	if !isEmailValid(users.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email address"})
	}

	// Check if user already exists
	var existingUser models.Users
	if err := database.DB.Db.Where("email = ?", users.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	users.Password = string(hashedPassword)

	// Ensure the RoleID is set to default if not provided
	if users.RoleID == 0 {
		users.RoleID = 1
	}

	if err := database.DB.Db.Create(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Retrieve the user with preloaded Role
	var user models.Users
	if err := database.DB.Db.Preload("Role").First(&user, users.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user with role"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// CreateCoach godoc
//
//	@Summary		Create a new coach
//	@Description	Create a new coach with role ID 2
//	@Tags			Coaches
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.Users	true	"Coach"
//	@Success		201	{object}	models.Users
//	@Failure		400	{string}	Cannot parse JSON
//	@Failure		400	{string}	Invalid email address
//	@Failure		400	{string}	User already exists
//	@Failure		500	{string}	Internal Server Error
//	@Router			/coaches [post]
func CreateCoach(c *fiber.Ctx) error {
	users := new(models.Users)
	if err := c.BodyParser(users); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	// Validate email format
	if !isEmailValid(users.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email address"})
	}

	// Check if user already exists
	var existingUser models.Users
	if err := database.DB.Db.Where("email = ?", users.Email).First(&existingUser).Error; err == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "User already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	users.Password = string(hashedPassword)

	// Ensure the RoleID is set to 2 for coaches
	users.RoleID = 2

	if err := database.DB.Db.Create(&users).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}

	// Retrieve the coach with preloaded Role
	var user models.Users
	if err := database.DB.Db.Preload("Role").First(&user, users.ID).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to retrieve user with role"})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetUser godoc
//
//	@Summary		Get a user by ID
//	@Description	Retrieve a user by ID
//	@Tags			Users
//	@Produce		json
//	@Param			id	path		int	true	"User ID"
//	@Success		200	{object}	models.Users
//	@Failure		400	{string}	Invalid user ID
//	@Failure		404	{string}	User not found
//	@Router			/users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user models.Users
	result := database.DB.Db.Preload("About").Preload("Groups").Preload("Role").First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// UpdateUser godoc
//
//	@Summary		Update a user by ID
//	@Description	Update a user's details by ID
//	@Tags			Users
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int			true	"User ID"
//	@Param			user	body	models.Users	true	"User"
//	@Success		200	{object}	models.Users
//	@Failure		400	{string}	Invalid user ID
//	@Failure		404	{string}	User not found
//	@Router			/users/{id} [put]
func UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var user models.Users
	result := database.DB.Db.Preload("About").Preload("Groups").Preload("Role").First(&user, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if err := database.DB.Db.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Internal Server Error"})
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

// DeleteUser godoc
//
//	@Summary		Delete a user by ID
//	@Description	Delete a user by ID
//	@Tags			Users
//	@Param			id	path	int	true	"User ID"
//	@Success		204	{string}	Successfully deleted
//	@Failure		400	{string}	Invalid user ID
//	@Failure		404	{string}	User not found
//	@Router			/users/{id} [delete]
func DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	result := database.DB.Db.Delete(&models.Users{}, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
