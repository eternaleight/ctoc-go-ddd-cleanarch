package usecases

import (
	"fmt"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

func CreatePost(store stores.PostStoreInterface, content string, userID uint) (*models.Post, error) {
	// 投稿内容が空の場合のエラーチェック
	if content == "" {
		return nil, fmt.Errorf("投稿内容がありません")
	}

	post := models.Post{
		Content:  content,
		AuthorID: userID,
	}

	if err := store.CreatePost(&post); err != nil {
		return nil, fmt.Errorf("サーバーエラーです。")
	}

	return &post, nil
}

func GetLatestPosts(store stores.PostStoreInterface) ([]models.Post, error) {
	posts, err := store.GetLatestPosts()
	if err != nil {
		return nil, fmt.Errorf("サーバーエラーです。")
	}

	return posts, nil
}
