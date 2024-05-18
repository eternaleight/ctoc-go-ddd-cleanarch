package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/gin-gonic/gin"
)

// 認証関連のリクエストを処理します
type AuthHandler struct {
	AuthUsecases usecases.AuthUsecasesInterface
}

// 新しいAuthHandlerのインスタンスを初期化します
func NewAuthHandler(authStore stores.AuthStoreInterface) *AuthHandler {
	return &AuthHandler{
		AuthUsecases: usecases.NewAuthUsecases(authStore),
	}
}

// 新しいユーザーを登録します
func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	c.JSON(http.StatusOK, gin.H{"user": user, "emailMd5Hash": emailMd5Hash, "gravatarURL": gravatarURL})
}

// ユーザーのログインを処理します
func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

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

	c.JSON(http.StatusOK, gin.H{"gravatarURL": gravatarURL})
}
