package usecases

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/eternaleight/go-backend/domain/models"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/utils"
)

// 認証に関するユースケースのインターフェースを定義します
type AuthUsecasesInterface interface {
	RegisterUser(username, email, password string) (*models.User, string, string, string, string, error)
	LoginUser(email, password string) (string, string, string, error)
	RefreshToken(refreshTokenString string) (string, string, error)
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
func (u AuthUsecases) RegisterUser(username, email, password string) (*models.User, string, string, string, string, error) {
	// ユーザーを登録
	user, err := u.AuthStore.RegisterUser(username, email, password)
	if err != nil {
		return nil, "", "", "", "", err
	}

	// GravatarのURLとメールのハッシュを生成
	emailMd5Hash := fmt.Sprintf("%x", utils.GetGravatarURL(email, 800))
	gravatarURL := utils.GetGravatarURL(email, 800)

	// アクセストークンを生成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(), // 30分の有効期限
	})
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, "", "", "", "", err
	}

	// リフレッシュトークンを生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 90).Unix(), // 90日の有効期限
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return nil, "", "", "", "", err
	}

	return user, emailMd5Hash, gravatarURL, accessTokenString, refreshTokenString, nil
}

// LoginUserはユーザーのログインを処理します
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

	// アクセストークンを生成
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Minute * 30).Unix(), // 30分の有効期限
	})
	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", "", err
	}

	// リフレッシュトークンを生成
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 90).Unix(), // 90日の有効期限
	})
	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", "", "", err
	}

	gravatarURL := utils.GetGravatarURL(email, 800)
	return gravatarURL, accessTokenString, refreshTokenString, nil
}

// RefreshTokenはリフレッシュトークンを使用して新しいアクセストークンを発行します
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

			// 新しいアクセストークンを生成
			accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":  userId,
				"exp": time.Now().Add(time.Minute * 30).Unix(), // 30分の有効期限
			})
			accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
			if err != nil {
				return "", "", err
			}

			// 新しいリフレッシュトークンを生成
			refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"id":  userId,
				"exp": time.Now().Add(time.Hour * 24 * 90).Unix(), // 90日の有効期限
			})
			refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
			if err != nil {
				return "", "", err
			}

			return accessTokenString, refreshTokenString, nil
		}
	}

	return "", "", fmt.Errorf("invalid refresh token claims")
}
