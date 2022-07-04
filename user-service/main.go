package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("just_a_secret_key")

type claims struct {
	User user `json:"user"`
	jwt.StandardClaims
}

type user struct {
	UserID           string `json:"user_id"`
	UserName         string `json:"user_name"`
	RegistrationDate string `json:"registration_date"`
	Role             string `json:"role"`
}

var users = map[string]user{
	"1": {
		UserID:           "1",
		UserName:         "Tom",
		RegistrationDate: "2020-01-01",
		Role:             "Author",
	},
	"2": {
		UserID:           "2",
		UserName:         "Jerry",
		RegistrationDate: "2020-01-02",
		Role:             "Viewer",
	},
}

func main() {
	app := fiber.New()
	app.Get("/users/:user_id", func(c *fiber.Ctx) error {
		userID := c.Params("user_id")
		return c.JSON(users[userID])
	})

	app.Post("/login/:user_id", func(c *fiber.Ctx) error {
		userID := c.Params("user_id")
		user := users[userID]
		expirationTime := time.Now().Add(30 * time.Minute)
		claims := &claims{
			User: user,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime.Unix(),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		c.Cookie(&fiber.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: expirationTime,
		})
		return c.JSON(fiber.Map{
			"user":   user,
			"token":  tokenString,
			"expire": expirationTime,
		})
	})

	app.Listen(":8081")
}
