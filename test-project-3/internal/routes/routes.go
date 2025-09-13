package routes

import (
	"net/http"

	"github.com/test/test-project-3/internal/config"
	"github.com/test/test-project-3/internal/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
// Dependencies holds all the dependencies needed by the routes
type Dependencies struct {
	DB     *gorm.DB
	Config *config.Config
}

// Setup configures all the routes
func Setup(router *gin.Engine, deps *Dependencies) {
	// Health check endpoint
	router.GET("/health", healthCheck)

	// API version 1 routes
	v1 := router.Group("/api/v1")
	{
		// Initialize handlers
		healthHandler := handlers.NewHealthHandler()
		userHandler := handlers.NewUserHandler(deps.DB, deps.Config)

		// Health routes
		v1.GET("/health", healthHandler.Check)
		v1.GET("/ready", healthHandler.Ready)

		// User routes
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}

		// Example protected routes (if authentication is implemented)
		// protected := v1.Group("/protected")
		// protected.Use(middleware.AuthRequired())
		// {
		//     protected.GET("/profile", userHandler.GetProfile)
		// }
	}

	// Serve static files (optional)
	// router.Static("/static", "./web/static")
	// router.StaticFile("/favicon.ico", "./web/static/favicon.ico")

	// Catch-all route for undefined endpoints
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Route not found",
			"path":  c.Request.URL.Path,
		})
	})
}

// healthCheck is a simple health check endpoint
func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"service": "test-project-3",
		"version": "1.0.0",
	})
}
