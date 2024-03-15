package routes

import (
	"api/internal/routes/category_routes"
	"api/internal/routes/event_routes"
	"api/internal/routes/user_routes"
	"api/internal/routes/authentication"
	"api/internal/middleware"

	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewServer() error {
	app := fiber.New()

	app.Use(cors.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	// User routes
	user := api.Group("/user", middleware.ValidateJWT)
	user.Get("/", user_routes.GetAllUsers)
	user.Get("/:id", user_routes.GetUser)
	user.Put("/:id", user_routes.EditUser)

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", authentication.Login)
	auth.Post("/register", authentication.CreateUser)

	// Category routes
	category := api.Group("/category")
	category.Get("/", category_routes.GetCategories)

	// Event routes
	event := api.Group("/event")
	event.Post("/", event_routes.CreateEvent)
	event.Get("/:id", event_routes.GetEventById)
	event.Get("/", event_routes.GetAllEvents)

	app.Listen(":" + os.Getenv("PORT"))

	return nil
}