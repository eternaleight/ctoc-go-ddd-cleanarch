package usecases

import (
	"fmt"

	"github.com/eternaleight/go-backend/domain/models"
)

// 投稿ストア操作のインターフェースを定義
type PostStoreInterface interface {
	CreatePost(post *models.Post) error
	GetLatestPosts() ([]models.Post, error)
}

type PostUsecases struct {
	PostStore PostStoreInterface
}

// PostUsecasesの新しいインスタンスを初期化
func NewPostUsecases(postStore PostStoreInterface) *PostUsecases {
	return &PostUsecases{
		PostStore: postStore,
	}
}

func (u *PostUsecases) CreatePost(content string, userID uint) (*models.Post, error) {
	// 投稿内容が空の場合のエラーチェック
	if content == "" {
		return nil, fmt.Errorf("投稿内容がありません")
	}

	post := models.Post{
		Content:  content,
		AuthorID: userID,
	}

	if err := u.PostStore.CreatePost(&post); err != nil {
		return nil, fmt.Errorf("サーバーエラーです。")
	}

	return &post, nil
}

func (u *PostUsecases) GetLatestPosts() ([]models.Post, error) {
	posts, err := u.PostStore.GetLatestPosts()
	if err != nil {
		return nil, fmt.Errorf("サーバーエラーです。")
	}

	return posts, nil
}
