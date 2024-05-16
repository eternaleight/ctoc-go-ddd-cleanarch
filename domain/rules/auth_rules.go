package rules

import (
	"crypto/md5"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// getGravatarURLは、指定されたメールアドレスとサイズに基づいてGravatarのURLを生成します
func GetGravatarURL(email string, size int) string {
	emailHash := fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(strings.TrimSpace(email)))))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=%d&d=identicon", emailHash, size)
}

// ComparePasswordは、ハッシュ化されたパスワードとプレーンなパスワードを比較します
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
