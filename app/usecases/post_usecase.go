package usecases

import (
	"fmt"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

type PostUsecasesInterface interface {
	CreatePost(content string, userID uint) (*models.Post, error)
	GetLatestPosts() ([]models.Post, error)
}

type PostUsecases struct {
	PostStore stores.PostStoreInterface
}

// NewPostUsecasesはPostUsecasesの新しいインスタンスを初期化します
func NewPostUsecases(postStore stores.PostStoreInterface) *PostUsecases {
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
