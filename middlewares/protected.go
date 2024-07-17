package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the Authorization header
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Missing or malformed JWT",
			})
		}

		// Validate the token
		token, err := jwt.Parse(strings.TrimPrefix(tokenString, "Bearer "), func(token *jwt.Token) (interface{}, error) {
			// Make sure that the token method conform to "SigningMethodHMAC"
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			// Return the secret signing key
			return []byte("your_secret_key"), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		}

		// Extract roles from the token claims
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid JWT claims",
			})
		}

		roles, ok := claims["roles"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid JWT claims",
			})
		}

		// Convert roles to a string slice
		rolesStr := make([]string, len(roles))
		for i, role := range roles {
			rolesStr[i], ok = role.(string)
			if !ok {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid role in JWT claims",
				})
			}
		}

		// Store roles and token in the context
		c.Locals("roles", rolesStr)
		c.Locals("user", token)

		return c.Next()
	}
}
