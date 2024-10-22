package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
)

// ユーザーストア操作のインターフェースを定義
type UserStoreInterface interface {
	CreateUser(user *models.User) error
	GetUserByID(id uint) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserUsecases struct {
	UserStore UserStoreInterface
}

// UserUsecasesの新しいインスタンスを初期化
func NewUserUsecases(userStore UserStoreInterface) *UserUsecases {
	return &UserUsecases{
		UserStore: userStore,
	}
}

func (u *UserUsecases) GetUserByID(userID uint) (*models.User, error) {
	// 指定されたIDのユーザー情報を取得
	user, err := u.UserStore.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
