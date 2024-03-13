package category_routes

import (
	"api/internal/database"

	"github.com/gofiber/fiber/v2"
)

func GetCategories(c *fiber.Ctx) error {
	categories := []Category{}

	rows, err := database.DBCon.Query("SELECT * FROM categories")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	for rows.Next() {
		var category Category
		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		categories = append(categories, category)
	}

	return c.JSON(categories)
}