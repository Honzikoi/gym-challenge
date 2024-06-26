package handlers

// Login route
// func Login(c *fiber.Ctx) error {
// 	// Extract the credentials from the request body
// 	loginRequest := new(models.LoginRequest)
// 	if err := c.BodyParser(loginRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	// Find the user by credentials
// 	user, err := repository.FindByCredentials(loginRequest.Email, loginRequest.Password)
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	day := time.Hour * 24
// 	// Create the JWT claims, which includes the user ID and expiry time
// 	claims := jtoken.MapClaims{
// 		"ID":    user.ID,
// 		"email": user.Email,
// 		"exp":   time.Now().Add(day * 1).Unix(),
// 	}
// 	// Create token
// 	token := jtoken.NewWithClaims(jtoken.SigningMethodHS256, claims)
// 	// Generate encoded token and send it as response.
// 	t, err := token.SignedString([]byte(config.Secret))
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 			"error": err.Error(),
// 		})
// 	}
// 	// Return the token
// 	return c.JSON(models.LoginResponse{
// 		Token: t,
// 	})
// }

// Protected route
// func Protected(c *fiber.Ctx) error {
// 	// Get the user from the context and return it
// 	user := c.Locals("user").(*jtoken.Token)
// 	claims := user.Claims.(jtoken.MapClaims)
// 	email := claims["email"].(string)
// 	favPhrase := claims["fav"].(string)
// 	return c.SendString("Welcome 👋" + email + " " + favPhrase)
// }

// func Singup(c *fiber.Ctx) error {

// }
