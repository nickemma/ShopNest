package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/internal/application"
)

// UserHandler handles HTTP requests for user operations
type CustomerHandler struct {
	service application.CustomerService
}

// NewUserHandler creates a new handler instance
func NewCustomerHandler(service application.CustomerService) *CustomerHandler {
	return &CustomerHandler{service: service}
}

// Register handles POST /register
func (h *CustomerHandler) RegisterCustomer(c *gin.Context) {
	var req struct {
		Name   string `json:"name" binding:"required"`
		Email  string `json:"email"`
		AuthId string `json:"authId"`
	}

	authId, ok := c.Get("authId")
	if !ok || authId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token, missing sub")})
		return
	}
	email, ok := c.Get("email")
	if !ok || authId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token, missing email")})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	req.AuthId = authId.(string)
	req.Email = email.(string)

	customerID, err := h.service.RegisterCustomer(c.Request.Context(), req.AuthId, req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"customerId": customerID})
}

// Register handles POST /activate
func (h *CustomerHandler) ActivateCustomer(c *gin.Context) {
	userType, ok := c.Get("userType")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing userId")})
		return
	}

	if userType.(string) != "MANAGER" {
		c.JSON(http.StatusForbidden, gin.H{"message": "only a manager can activate user"})
		return
	}
	var req struct {
		AuthID string `json:"authId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customerID, err := h.service.ActivateCustomer(c.Request.Context(), req.AuthID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"customerId": customerID, "message": "customer activated"})
}

func (h *CustomerHandler) GetCustomerProfile(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing userId")})
		return
	}

	profile, err := h.service.GetCustomerProfile(c.Request.Context(), userId.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": profile})
	return
}
