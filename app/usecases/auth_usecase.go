package usecases

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/utils"
)

// 認証に関するユースケースのインターフェースを定義します
type AuthUsecasesInterface interface {
	RegisterUser(username, email, password string) (*models.User, string, string, string, error)
	LoginUser(email, password string) (string, string, error)
}

// 認証に関するユースケースの具体的な実装を定義します
type AuthUsecases struct {
	AuthStore stores.AuthStoreInterface
}

// 新しいインスタンスを初期化します
func NewAuthUsecases(authStore stores.AuthStoreInterface) *AuthUsecases {
	return &AuthUsecases{
		AuthStore: authStore,
	}
}

// ユーザーを登録します
func (u AuthUsecases) RegisterUser(username, email, password string) (*models.User, string, string, string, error) {
	// ユーザーを登録
	user, err := u.AuthStore.RegisterUser(username, email, password)
	if err != nil {
		return nil, "", "", "", err
	}

	// GravatarのURLとメールのハッシュを生成
	emailMd5Hash := fmt.Sprintf("%x", utils.GetGravatarURL(email, 800))
	gravatarURL := utils.GetGravatarURL(email, 800)

	// JWTトークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, "", "", "", err
	}

	return user, emailMd5Hash, gravatarURL, tokenString, nil
}

// LoginUserはユーザーのログインを処理します
func (u AuthUsecases) LoginUser(email, password string) (string, string, error) {
	// メールでユーザーを取得
	user, err := u.AuthStore.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}

	// パスワードを比較
	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", err
	}

	// JWTトークンを生成
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
