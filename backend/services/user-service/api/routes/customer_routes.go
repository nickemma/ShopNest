package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/middleware"
	"github.com/shopnest/user-service/config"
)

// SetupRouter configures the Gin router
func RegisterCustomerRoutes(r *gin.RouterGroup, handler *handler.CustomerHandler, cfg config.Config) {

	// User Public routes

	// Protected routes
	auth := r.Group("/")

	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Add your protected routes here
		auth.POST("/register", handler.RegisterCustomer)

		auth.PATCH("/activate", handler.ActivateCustomer)
		auth.GET("/profile", handler.GetCustomerProfile)
	}

}
