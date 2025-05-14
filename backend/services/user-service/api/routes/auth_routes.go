package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/api/handler"
	"github.com/shopnest/user-service/api/middleware"
	"github.com/shopnest/user-service/config"
)

// SetupRouter configures the Gin router
func RegisterAuthRoutes(r *gin.RouterGroup, handler *handler.AuthHandler, cfg config.Config) {

	// User Public routes
	r.POST("/register", handler.Register)
	r.GET("/verify-email", handler.VerifyEmail)
	r.POST("/login", handler.Login)

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware(cfg.JWTSecret))
	{
		// Add your protected routes here
		auth.GET("/account", handler.GetAccount)
		auth.GET("/refresh", handler.RefreshToken)
	}

}
