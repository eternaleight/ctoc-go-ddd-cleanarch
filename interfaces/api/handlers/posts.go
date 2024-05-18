package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/gin-gonic/gin"
)

// 投稿関連のリクエストを処理します
type PostHandler struct {
	PostUsecases usecases.PostUsecasesInterface
}

// 新しいPostHandlerのインスタンスを初期化します
func NewPostHandler(postUsecases usecases.PostUsecasesInterface) *PostHandler {
	return &PostHandler{
		PostUsecases: postUsecases,
	}
}

// 新しい投稿を作成します
func (h *PostHandler) CreatePost(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	// リクエストからJSONデータをバインドします
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// isAuthenticatedミドルウェアで設定されたuserIDを取得します
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDが見つかりません"})
		return
	}

	// userIDの型を確認します
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの型が正しくありません"})
		return
	}

	// ユースケースを呼び出します
	post, err := h.PostUsecases.CreatePost(input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// 最新の投稿を取得します
func (h *PostHandler) GetLatestPosts(c *gin.Context) {
	// ユースケースを呼び出します
	posts, err := h.PostUsecases.GetLatestPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
