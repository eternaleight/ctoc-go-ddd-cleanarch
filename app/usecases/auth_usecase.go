package usecases

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/domain/rules"
	"github.com/eternaleight/go-backend/infra/stores"
)

func RegisterUser(store stores.AuthStoreInterface, username, email, password string) (*models.User, string, string, string, error) {
	// Register the user
	user, err := store.RegisterUser(username, email, password)
	if err != nil {
		return nil, "", "", "", err
	}

	emailMd5Hash := fmt.Sprintf("%x", rules.GetGravatarURL(email, 800))

	gravatarURL := rules.GetGravatarURL(email, 800)

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

func LoginUser(store stores.AuthStoreInterface, email, password string) (string, string, error) {
	// Retrieve user by email
	user, err := store.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	// Compare password
	err = rules.ComparePassword(user.Password, password)
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

	gravatarURL := rules.GetGravatarURL(email, 800)
	return gravatarURL, tokenString, nil
}
