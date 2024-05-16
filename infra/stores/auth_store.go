package stores

import (
	"github.com/eternaleight/go-backend/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// AuthStoreInterface defines the interface for authentication store operations
type AuthStoreInterface interface {
	RegisterUser(username, email, password string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// AuthStore manages database operations related to authentication
type AuthStore struct {
	DB *gorm.DB
}

// NewAuthStore creates a new instance of AuthStore
func NewAuthStore(db *gorm.DB) *AuthStore {
	return &AuthStore{DB: db}
}

// RegisterUser registers a new user in the database
func (s *AuthStore) RegisterUser(username, email, password string) (*models.User, error) {
	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	}

	// Save the user information in the database
	user := &models.User{
		Username: username,
		Email:    email,
		Password: string(hashedPassword),
	}
	result := s.DB.Create(user)
	return user, result.Error
}

// GetUserByEmail retrieves user information based on email
func (s *AuthStore) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := s.DB.Where("email = ?", email).First(&user).Error
	return &user, err
}
