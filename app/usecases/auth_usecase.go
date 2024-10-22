package usecases

import (
	"fmt"
	"os"
	"time"

	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/utils"
	"github.com/golang-jwt/jwt/v4"
)

// 認証ストア操作のインターフェースを定義
type AuthStoreInterface interface {
	RegisterUser(username, email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// 認証に関するユースケースの具体的な実装を定義
type AuthUsecases struct {
	AuthStore AuthStoreInterface
}

// 新しいインスタンスを初期化
func NewAuthUsecases(authStore AuthStoreInterface) *AuthUsecases {
	return &AuthUsecases{
		AuthStore: authStore,
	}
}

// ユーザーを登録
func (u AuthUsecases) RegisterUser(username, email, password string) (*models.User, string, string, string, string, error) {
	// ユーザーを登録
	user, err := u.AuthStore.RegisterUser(username, email, password)
	if err != nil {
		return nil, "", "", "", "", err
	}

	// GravatarのURLとメールのハッシュを生成
	emailMd5Hash := fmt.Sprintf("%x", utils.GetGravatarURL(email, 800))
	gravatarURL := utils.GetGravatarURL(email, 800)

	// アクセストークンとリフレッシュトークンを生成
	accessToken, refreshToken, err := generateTokens(user.ID)
	if err != nil {
		return nil, "", "", "", "", err
	}

	return user, emailMd5Hash, gravatarURL, accessToken, refreshToken, nil
}

// LoginUserはユーザーのログインを処理
func (u AuthUsecases) LoginUser(email, password string) (string, string, string, error) {
	// メールでユーザーを取得
	user, err := u.AuthStore.GetUserByEmail(email)
	if err != nil {
		return "", "", "", err
	}

	// パスワードを比較
	err = utils.ComparePassword(user.Password, password)
	if err != nil {
		return "", "", "", err
	}

	// アクセストークンとリフレッシュトークンを生成
	accessToken, refreshToken, err := generateTokens(user.ID)
	if err != nil {
		return "", "", "", err
	}

	gravatarURL := utils.GetGravatarURL(email, 800)
	return gravatarURL, accessToken, refreshToken, nil
}

// RefreshTokenはリフレッシュトークンを使用して新しいアクセストークンを発行
func (u AuthUsecases) RefreshToken(refreshTokenString string) (string, string, error) {
	// リフレッシュトークンを解析
	token, err := jwt.Parse(refreshTokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return "", "", fmt.Errorf("invalid refresh token")
	}

	// トークンのクレームからユーザーIDを取得
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if idFloat, ok := claims["id"].(float64); ok {
			userId := uint(idFloat)

			// 新しいアクセストークンとリフレッシュトークンを生成
			accessToken, refreshToken, err := generateTokens(userId)
			if err != nil {
				return "", "", err
			}

			return accessToken, refreshToken, nil
		}
	}

	return "", "", fmt.Errorf("invalid refresh token claims")
}

// トークンを生成するヘルパー関数
func generateTokens(userID uint) (string, string, error) {
	// アクセストークンを生成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Minute * 30).Unix(), // 30分の有効期限
	})
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	// リフレッシュトークンを生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userID,
		"exp": time.Now().Add(time.Hour * 24 * 90).Unix(), // 90日の有効期限
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}
