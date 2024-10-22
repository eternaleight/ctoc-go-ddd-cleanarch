package dtos

// ユーザー登録リクエストのDTO（Data Transfer Object）
type RegisterRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// ユーザーログインリクエストのDTO
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// 認証レスポンスのDTO
type AuthResponse struct {
	User         interface{} `json:"user"`
	EmailMd5Hash string      `json:"emailMd5Hash"`
	GravatarURL  string      `json:"gravatarURL"`
	Token        string      `json:"token"`
}
