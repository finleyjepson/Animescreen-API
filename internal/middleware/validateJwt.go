package middleware

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// ValidateJWT validates the JWT token
func ValidateJWT(c *fiber.Ctx) error {

	// Get the bearer token from the header
	accessToken := c.Get("Authorization")

	if accessToken == "" {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Strip the bearer prefix from the token
	accessToken = accessToken[7:]

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	if token.Valid {
		return c.Next()
	}
	
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"message": "Unauthorized",
	})
}
