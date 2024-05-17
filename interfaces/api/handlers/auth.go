package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	AuthUsecases usecases.AuthUsecasesInterface
}

// NewAuthHandler initializes a new instance of AuthHandler
func NewAuthHandler(authStore stores.AuthStoreInterface) *AuthHandler {
	return &AuthHandler{
		AuthUsecases: usecases.NewAuthUsecases(authStore),
	}
}

// Register a new user
func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON data from the request
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the usecase
	user, emailMd5Hash, gravatarURL, tokenString, err := h.AuthUsecases.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Set JWT token as HTTP-only cookie with a 90-day expiration
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"user": user, "emailMd5Hash": emailMd5Hash, "gravatarURL": gravatarURL})
}

// Handle user login
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind JSON data from the request
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the usecase
	gravatarURL, tokenString, err := h.AuthUsecases.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// Set JWT token as HTTP-only cookie with a 90-day expiration
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{"gravatarURL": gravatarURL})
}
