package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/internal/application"
)

// AuthHandler handles HTTP requests for user operations
type AuthHandler struct {
	service application.AuthService
}

// NewAuthHandler creates a new handler instance
func NewAuthHandler(service application.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

// Register handles POST /register
func (h *AuthHandler) Register(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
		Role     string `json:"role" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	authID, err := h.service.CreateAccount(c.Request.Context(), req.Email, req.Password, req.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"authId": authID})
}

// Login handles POST /login
func (h *AuthHandler) Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.service.Login(c.Request.Context(), req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// VerifyEmail handles GET /verify-email?token=<token>
func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required....."})
		return
	}

	if err := h.service.VerifyEmail(c.Request.Context(), token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "email verified"})
}

// GET /account
func (h *AuthHandler) GetAccount(c *gin.Context) {
	authId, ok := c.Get("authId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing sub")})
		return
	}

	if authId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token required"})
		return
	}

	auth, err := h.service.GetAccount(c.Request.Context(), authId.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": auth})
}
