package routes

import (
	"github.com/user/test-gorm/internal/handlers"
	"github.com/user/test-gorm/internal/repositories"
	"github.com/user/test-gorm/internal/services"
	"github.com/user/test-gorm/internal/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Setup configures all the routes
func Setup(app *fiber.App, db *gorm.DB, validator *utils.Validator) {
	// Initialize dependencies
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService, validator)

	// Health check endpoint
	app.Get("/health", handlers.HealthCheck)

	// API version 1 routes
	v1 := app.Group("/api/v1")
	{
		// User routes
		users := v1.Group("/users")
		{
			users.Get("", userHandler.GetUsers)
			users.Get("/:id", userHandler.GetUser)
			users.Post("", userHandler.CreateUser)
			users.Put("/:id", userHandler.UpdateUser)
			users.Delete("/:id", userHandler.DeleteUser)
		}
	}
}
