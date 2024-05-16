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
// Gravatarの仕様がmd5を要求するため、Gravatar用のハッシュにはmd5を使用することが標準的な実装。
// MD5はセキュリティ上の問題があるため、他の用途には使用しないでください。

// ComparePasswordは、ハッシュ化されたパスワードとプレーンなパスワードを比較します
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
