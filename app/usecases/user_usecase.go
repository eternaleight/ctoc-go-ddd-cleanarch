package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

type UserUsecasesInterface interface {
	GetUserByID(userID uint) (*models.User, error)
}

type UserUsecases struct {
	UserStore stores.UserStoreInterface
}

// UserUsecasesの新しいインスタンスを初期化します
func NewUserUsecases(userStore stores.UserStoreInterface) *UserUsecases {
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
