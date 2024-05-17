package usecases

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/utils"
)

type AuthUsecasesInterface interface {
	RegisterUser(username, email, password string) (*models.User, string, string, string, error)
	LoginUser(email, password string) (string, string, error)
}

type AuthUsecases struct {
	AuthStore stores.AuthStoreInterface
}

// NewPostUsecasesはPostUsecasesの新しいインスタンスを初期化します
func NewAuthUsecases(authStore stores.AuthStoreInterface) *AuthUsecases {
	return &AuthUsecases{
		AuthStore: authStore,
	}
}

func (u AuthUsecases) RegisterUser(username, email, password string) (*models.User, string, string, string, error) {
	// Register the user
	user, err := u.AuthStore.RegisterUser(username, email, password)
	if err != nil {
		return nil, "", "", "", err
	}

	emailMd5Hash := fmt.Sprintf("%x", utils.GetGravatarURL(email, 800))

	gravatarURL := utils.GetGravatarURL(email, 800)

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, "", "", "", err
	}

	return user, emailMd5Hash, gravatarURL, tokenString, nil
}

func (u AuthUsecases) LoginUser(email, password string) (string, string, error) {
	// Retrieve user by email
	user, err := u.AuthStore.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	// Compare password
	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", err
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	gravatarURL := utils.GetGravatarURL(email, 800)
	return gravatarURL, tokenString, nil
}
