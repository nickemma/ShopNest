package handler

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shopnest/user-service/internal/application"
)

// ManagerHandler handles HTTP requests for user operations
type ManagerHandler struct {
	service application.ManagerService
}

// NewManagerHandler creates a new handler instance
func NewManagerHandler(service application.ManagerService) *ManagerHandler {
	return &ManagerHandler{service: service}
}

// Register handles POST /register
func (h *ManagerHandler) RegisterManager(c *gin.Context) {
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

	userID, err := h.service.Register(c.Request.Context(), req.AuthId, req.Name, req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"userId": userID})
}

// Register handles POST /activate
func (h *ManagerHandler) ApproveRegistration(c *gin.Context) {
	userType, ok := c.Get("userType")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing userId")})
		return
	}

	email, ok := c.Get("email")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing email")})
		return
	}

	// TODO: this will be extended with roles or super admin only
	// workaround: encode the super admin
	if userType.(string) != "MANAGER" && email != "superadmin@gmail.com"  {
		c.JSON(http.StatusForbidden, gin.H{"message": "only a manager can approve other manager"})
		return
	}
	var req struct {
		AuthID string `json:"authId" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := h.service.ApproveRegistration(c.Request.Context(), req.AuthID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"userId": userID, "message": "user activated"})
}

func (h *ManagerHandler) GetManagerProfile(c *gin.Context) {
	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("malformed access token. missing userId")})
		return
	}

	profile, err := h.service.GetManagerProfile(c.Request.Context(), userId.(string))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": profile})
	return
}
