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
	user, emailMd5Hash, gravatarURL, tokenString, err := h.AuthUsecases.RegisterUser(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンをHTTPオンリークッキーとして設定（90日間の有効期限）
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)

	response := dtos.AuthResponse{
		User:         user,
		EmailMd5Hash: emailMd5Hash,
		GravatarURL:  gravatarURL,
		Token:        tokenString,
	}

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
	gravatarURL, tokenString, err := h.AuthUsecases.LoginUser(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// JWTトークンをHTTPオンリークッキーとして設定（90日間の有効期限）
	c.SetCookie("authToken", tokenString, 60*60*24*90, "/", "", false, true)

	response := dtos.AuthResponse{
		GravatarURL: gravatarURL,
		Token:       tokenString,
	}

	c.JSON(http.StatusOK, response)
}
