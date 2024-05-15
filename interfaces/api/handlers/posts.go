package handlers

import (
	"net/http"

	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	PostStore stores.PostStoreInterface
}

// NewPostHandlerはPostHandlerの新しいインスタンスを初期化します
func NewPostHandler(postStore stores.PostStoreInterface) *PostHandler {
	return &PostHandler{
		PostStore: postStore,
	}
}

// CreatePostは新しい投稿を作成します
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
	post, err := usecases.CreatePost(h.PostStore, input.Content, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"post": post})
}

// GetLatestPostsは最新の投稿を取得します
func (h *PostHandler) GetLatestPosts(c *gin.Context) {
	// ユースケースを呼び出します
	posts, err := usecases.GetLatestPosts(h.PostStore)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "サーバーエラーです。"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
