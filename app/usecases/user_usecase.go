package usecases

import (
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
)

func GetUserByID(store stores.UserStoreInterface, userID uint) (*models.User, error) {
	// 指定されたIDのユーザー情報を取得
	user, err := store.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
