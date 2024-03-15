package authentication

import (
	"api/internal/database"
	"api/internal/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Logins a user
func Login(c *fiber.Ctx) error {
	var u *User = new(User)

	if err := c.BodyParser(u); err != nil {
		return err
	}

	// Check if user exists
	userSQL := database.DBCon.QueryRow("SELECT id, username, password FROM users WHERE username = $1", u.Username)
	if userSQL.Err() != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}
	var user *User = new(User)
	err := userSQL.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Println(err)
	}
	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Invalid credentials",
		})
	}

	// Generate JWT
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to generate token",
		})
	}


	// Return user
	return c.JSON(fiber.Map{
		"message": "Logged in",
		"user": fiber.Map{
			"ID": user.ID,
			"Username": user.Username,
		},
		"tokens": fiber.Map{
			"access_token": token,
		},
	})
}