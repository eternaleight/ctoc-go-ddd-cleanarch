package stores

import (
	"github.com/eternaleight/go-backend/domain/models"
	"gorm.io/gorm"
)

// ユーザーストア操作のインターフェースを定義
type UserStoreInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// ユーザーに関するデータベース操作を管理
type UserStore struct {
	DB *gorm.DB
}

// 新しいUserStoreのインスタンスを生成
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{DB: db}
}

// ユーザーをデータベースに保存
func (s *UserStore) CreateUser(user *models.User) error {
	return s.DB.Create(user).Error
}

// IDに基づいてユーザー情報を取得
func (s *UserStore) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// メールアドレスに基づいてユーザー情報を取得
func (s *UserStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
