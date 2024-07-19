package middlewares

import (
	"encoding/base64"
	"os"
	"strings"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func RequireAuth(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing authorization header"})
	}

	authType := strings.Fields(authHeader)[0]
	authValue := strings.Fields(authHeader)[1]

	if authType == "Bearer" {
		token, err := jwt.Parse(authValue, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("userID", claims["user_id"])
			c.Locals("role", claims["role"])
			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token"})
		}
	} else if authType == "Basic" {
		decoded, err := base64.StdEncoding.DecodeString(authValue)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Basic Auth"})
		}

		credentials := strings.SplitN(string(decoded), ":", 2)
		if len(credentials) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Basic Auth"})
		}

		email := credentials[0]
		password := credentials[1]

		// Check the user's credentials
		db := database.DB.Db
		var user models.Users
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid email or password"})
		}

		c.Locals("userID", user.ID)
		c.Locals("role", user.Role.Name)
		return c.Next()
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid authorization method"})
}
