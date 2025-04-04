package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/middleware"
	"github.com/shopnest/user-service/config"
)

// SetupRouter configures the Gin router
func SetupRouter(handler *handler.UserHandler, cfg config.Config) *gin.Engine {
	r := gin.Default()

	// User Public routes
	r.POST("/register", handler.Register)
	r.GET("/verify-email", handler.VerifyEmail)
	r.POST("/login", handler.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Add your protected routes here
		// auth.GET("/profile", h.GetProfile)
	}

	return r
}
