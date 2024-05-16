package interfaces

import (
	"os"

	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/interfaces/api/handlers"
	"github.com/eternaleight/go-backend/interfaces/api/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {

	r := gin.Default()

	// トレーリングスラッシュへのリダイレクトを無効にする
	r.RedirectTrailingSlash = false

	config := cors.DefaultConfig()
	config.AllowCredentials = true

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS") // 環境変数から読み取る
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:3000" // デフォルト値
	}
	config.AllowOrigins = []string{allowedOrigins} // フロントエンドのオリジンに合わせて変更

	r.Use(cors.New(config))

	// 'Authorization'ヘッダーを許可するためにヘッダーを追加
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")

	// ストアのインスタンスを作成
	authStore := stores.NewAuthStore(db)
	productStore := stores.NewProductStore(db)
	purchaseStore := stores.NewPurchaseStore(db)
	userStore := stores.NewUserStore(db)
	postStore := stores.NewPostStore(db)

	// ハンドラのインスタンスを作成
	authHandler := handlers.NewAuthHandler(authStore)
	productHandler := handlers.NewProductHandler(productStore)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseStore)
	userHandler := handlers.NewUserHandler(userStore)
	postHandler := handlers.NewPostHandler(postStore)

	// auth
	auth := r.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// posts
	posts := r.Group("/api/posts").Use(middlewares.IsAuthenticated())
	{
		posts.POST("", postHandler.CreatePost)
		posts.GET("", postHandler.GetLatestPosts)
	}

	// user
	user := r.Group("/api/user").Use(middlewares.IsAuthenticated())
	user.GET("", userHandler.GetUser)

	// products
	products := r.Group("/api/products").Use(middlewares.IsAuthenticated())
	{
		products.POST("", productHandler.CreateProduct)
		products.GET("", productHandler.ListProducts)
		products.GET("/:id", productHandler.GetProductByID)
		products.PUT("/:id", productHandler.UpdateProduct)
		products.DELETE("/:id", productHandler.DeleteProduct)
	}

	// purchase
	purchase := r.Group("/api/purchase").Use(middlewares.IsAuthenticated())
	purchase.POST("", purchaseHandler.CreatePurchase)

	return r
}
