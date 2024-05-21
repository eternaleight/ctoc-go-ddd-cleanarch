package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/dtos"
	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/gin-gonic/gin"
)

// 認証関連のリクエストを処理します
type AuthHandler struct {
	AuthUsecases usecases.AuthUsecasesInterface
}

// 新しいAuthHandlerのインスタンスを初期化します
func NewAuthHandler(authUsecases usecases.AuthUsecasesInterface) *AuthHandler {
	return &AuthHandler{
		AuthUsecases: authUsecases,
	}
}

// 新しいユーザーを登録します
func (h *AuthHandler) Register(c *gin.Context) {
	var input dtos.RegisterRequest

	// リクエストからJSONデータをバインドします
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユースケースを呼び出します
	user, emailMd5Hash, gravatarURL, accessTokenString, refreshTokenString, err := h.AuthUsecases.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dtos.AuthResponse{
		User:         user,
		EmailMd5Hash: emailMd5Hash,
		GravatarURL:  gravatarURL,
		Token:        accessTokenString,
	}

	// アクセストークンとリフレッシュトークンをヘッダーに設定
	c.Header("Authorization", "Bearer "+accessTokenString)
	c.Header("Refresh-Token", refreshTokenString)

	c.JSON(http.StatusOK, response)
}

// ユーザーのログインを処理します
func (h *AuthHandler) Login(c *gin.Context) {
	var input dtos.LoginRequest

	// リクエストからJSONデータをバインドします
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// ユースケースを呼び出します
	gravatarURL, accessTokenString, refreshTokenString, err := h.AuthUsecases.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	response := dtos.AuthResponse{
		GravatarURL: gravatarURL,
		Token:       accessTokenString,
	}

	// アクセストークンとリフレッシュトークンをヘッダーに設定
	c.Header("Authorization", "Bearer "+accessTokenString)
	c.Header("Refresh-Token", refreshTokenString)

	c.JSON(http.StatusOK, response)
}

// リフレッシュトークンを使用してアクセストークンをリフレッシュします
func (h *AuthHandler) Refresh(c *gin.Context) {
	refreshTokenString := c.GetHeader("Refresh-Token")
	if refreshTokenString == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token is required"})
		return
	}

	accessTokenString, newRefreshTokenString, err := h.AuthUsecases.RefreshToken(refreshTokenString)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// 新しいアクセストークンとリフレッシュトークンをヘッダーに設定
	c.Header("Authorization", "Bearer "+accessTokenString)
	c.Header("Refresh-Token", newRefreshTokenString)

	c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}
