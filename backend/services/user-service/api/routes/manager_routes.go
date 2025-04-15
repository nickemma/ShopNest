package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/middleware"
	"github.com/shopnest/user-service/config"
)

// SetupRouter configures the Gin router
func RegisterManagerRoutes(r *gin.RouterGroup, handler *handler.ManagerHandler, cfg config.Config) {

	// User Public routes

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Add your protected routes here
		auth.POST("/register", handler.RegisterManager)

		auth.PATCH("/approve", handler.ApproveRegistration)
		auth.GET("/profile", handler.GetManagerProfile)
	}

}
