package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/gin-gonic/gin"
)

type UserUsecasesInterface interface {
	GetUserByID(userID uint) (*models.User, error)
}

// ユーザー関連のハンドラを管理
type UserHandler struct {
	UserUsecases UserUsecasesInterface
}

// UserHandlerの新しいインスタンスを初期化
func NewUserHandler(userUsecases UserUsecasesInterface) *UserHandler {
	return &UserHandler{
		UserUsecases: userUsecases,
	}
}

// ユーザー情報を取得
func (h *UserHandler) GetUser(c *gin.Context) {
	// ミドルウェアからuserIDを取得
	userID := c.MustGet("userID").(uint)

	// ユースケースの呼び出し
	user, err := h.UserUsecases.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つからないか、データベースエラー"})
		return
	}

	// ユーザー情報をレスポンスとして返す
	c.JSON(http.StatusOK, gin.H{"user": user})
}
