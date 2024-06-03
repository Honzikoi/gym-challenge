package handlers

import (
	"context"
	"encoding/json"
	"os"

	"github.com/Honzikoi/gym-challenge/database"
	"github.com/Honzikoi/gym-challenge/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/microsoft"
	"google.golang.org/api/oauth2/v2"
)

var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback",
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

var microsoftOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/microsoft/callback",
	ClientID:     os.Getenv("MICROSOFT_CLIENT_ID"),
	ClientSecret: os.Getenv("MICROSOFT_CLIENT_SECRET"),
	Scopes:       []string{"User.Read"},
	Endpoint:     microsoft.OAuth2Endpoint,
}

func GoogleLogin(c *fiber.Ctx) error {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not get token")
	}

	service, err := oauth2.New(googleOauthConfig.Client(context.Background(), token))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create OAuth2 service")
	}

	userInfo, err := service.Userinfo.Get().Do()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not get user info")
	}

	var user models.Users
	database.DB.Db.Where("email = ?", userInfo.Email).First(&user)

	if user.ID == 0 {
		user.Email = userInfo.Email
		user.Username = userInfo.Email // or use userInfo.Name
		user.Password = ""             // OAuth users don't have a password
		if err := database.DB.Db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not create user")
		}
	}

	tokenString, err := createToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create token")
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}

func MicrosoftLogin(c *fiber.Ctx) error {
	url := microsoftOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func MicrosoftCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := microsoftOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not get token")
	}

	client := microsoftOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://graph.microsoft.com/v1.0/me")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not get user info")
	}

	defer resp.Body.Close()
	userInfo := struct {
		Email    string `json:"mail"`
		Username string `json:"displayName"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not parse user info")
	}

	var user models.Users
	database.DB.Db.Where("email = ?", userInfo.Email).First(&user)

	if user.ID == 0 {
		user.Email = userInfo.Email
		user.Username = userInfo.Username // or use userInfo.Name
		user.Password = ""                // OAuth users don't have a password
		if err := database.DB.Db.Create(&user).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Could not create user")
		}
	}

	tokenString, err := createToken(user.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not create token")
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})
}
