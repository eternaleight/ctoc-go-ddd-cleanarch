package stores

import (
	"github.com/eternaleight/go-backend/domain/models"
	"gorm.io/gorm"
)

// 投稿ストア操作のインターフェースを定義
type PostStoreInterface interface {
	CreatePost(post *models.Post) error
	GetLatestPosts() ([]models.Post, error)
}

// 投稿に関するデータベース操作を管理
type PostStore struct {
	DB *gorm.DB
}

// 新しいPostStoreのインスタンスを作成
func NewPostStore(db *gorm.DB) *PostStore {
	return &PostStore{DB: db}
}

// 新しい投稿をデータベースに保存
func (s *PostStore) CreatePost(post *models.Post) error {
	return s.DB.Create(post).Error
}

// 最新の投稿を取得
func (s *PostStore) GetLatestPosts() ([]models.Post, error) {
	var posts []models.Post
	// 投稿日時の降順に10件の投稿を取得し、それらの投稿者も同時に取得
	if err := s.DB.Order("\"createdAt\" desc").Limit(10).Preload("Author").Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}
