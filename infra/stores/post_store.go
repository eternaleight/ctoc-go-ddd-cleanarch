package stores

import (
	"github.com/eternaleight/go-backend/domain/models"
	"gorm.io/gorm"
)

// PostStoreInterfaceは投稿ストア操作のインターフェースを定義します
type PostStoreInterface interface {
	CreatePost(post *models.Post) error
	GetLatestPosts() ([]models.Post, error)
}

// PostStoreは投稿に関するデータベース操作を管理します
type PostStore struct {
	DB *gorm.DB
}

// NewPostStoreはPostStoreの新しいインスタンスを作成します
func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{DB: db}
}

// CreatePostは新しい投稿をデータベースに保存します
func (s *PostStore) CreatePost(post *models.Post) error {
	return s.DB.Create(post).Error
}

// GetLatestPostsは最新の投稿を取得します
func (s *PostStore) GetLatestPosts() ([]models.Post, error) {
	var posts []models.Post
	// 投稿日時の降順に10件の投稿を取得し、それらの投稿者も同時に取得します
	if err := s.DB.Order("\"createdAt\" desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
