package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"strings"
)

// Gravatarの仕様がmd5を要求するため、Gravatar用のハッシュにはmd5を使用することが標準的な実装。
// MD5はセキュリティ上の問題があるため、他の用途には使用しないでください。
// getGravatarURLは、指定されたメールアドレスとサイズに基づいてGravatarのURLを生成します
func GetGravatarURL(email string, size int) string {
	emailHash := fmt.Sprintf("%x", md5.Sum([]byte(strings.ToLower(strings.TrimSpace(email)))))
	return fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=%d&d=identicon", emailHash, size)
}

// ComparePasswordは、ハッシュ化されたパスワードとプレーンなパスワードを比較します
func ComparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Encryptは、指定されたテキストをAES暗号化して返します
func Encrypt(text string) (string, error) {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	b := base64.StdEncoding.EncodeToString([]byte(text))
	ciphertext := make([]byte, aes.BlockSize+len(b))
	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(b))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decryptは、AES暗号化されたテキストを復号化して返します
func Decrypt(cryptoText string) (string, error) {
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	data, err := base64.StdEncoding.DecodeString(string(ciphertext))
	if err != nil {
		return "", err
	}

	return string(data), nil
}
