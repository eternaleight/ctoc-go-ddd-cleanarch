package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/gin-gonic/gin"
)

// 投稿関連のリクエストを処理
type PostHandler struct {
	PostUsecases usecases.PostUsecasesInterface
}

// 新しいPostHandlerのインスタンスを初期化
func NewPostHandler(postUsecases usecases.PostUsecasesInterface) *PostHandler {
	return &PostHandler{
		PostUsecases: postUsecases,
	}
}

// 新しい投稿を作成
func (h *PostHandler) CreatePost(c *gin.Context) {
	var input struct {
		Content string `json:"content"`
	}

	// リクエストからJSONデータをバインド
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// isAuthenticatedミドルウェアで設定されたuserIDを取得
	userIDValue, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDが見つかりません"})
		return
	}

	// userIDの型を確認
	userID, ok := userIDValue.(uint)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーIDの型が正しくありません"})
		return
	}

	// ユースケースを呼び出す
	post, err := h.PostUsecases.CreatePost(input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// 最新の投稿を取得
func (h *PostHandler) GetLatestPosts(c *gin.Context) {
	// ユースケースを呼び出す
	posts, err := h.PostUsecases.GetLatestPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
