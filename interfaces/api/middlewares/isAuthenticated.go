package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
)

func IsAuthenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Authorizationヘッダーからトークンを取得
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}

		// Bearerトークンを取り出す
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Bearer token is required"})
			c.Abort()
			return
		}

		// JWTトークンを解析
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// トークンのクレームからユーザーIDを取得
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if idFloat, ok := claims["id"].(float64); ok {
				userId := uint(idFloat)
				c.Set("userID", userId)
				fmt.Println("UserID set in middleware:", userId)
				c.Next()
				return
			}
		}

		// 何らかの理由で認証が失敗した場合
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		c.Abort()
	}
}
