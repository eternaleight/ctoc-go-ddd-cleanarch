package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/gin-gonic/gin"
)

// UserHandlerはユーザー関連のハンドラを管理します
type UserHandler struct {
	UserStore stores.UserStoreInterface
}

// NewUserHandlerはUserHandlerの新しいインスタンスを初期化します
func NewUserHandler(userStore stores.UserStoreInterface) *UserHandler {
	return &UserHandler{
		UserStore: userStore,
	}
}

// GetUserはユーザー情報を取得します
func (h *UserHandler) GetUser(c *gin.Context) {
	// ミドルウェアからuserIDを取得
	userID := c.MustGet("userID").(uint)

	// ユースケースの呼び出し
	user, err := usecases.GetUserByID(h.UserStore, userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つからないか、データベースエラー"})
		return
	}

	// ユーザー情報をレスポンスとして返す
	c.JSON(http.StatusOK, gin.H{"user": user})
}
