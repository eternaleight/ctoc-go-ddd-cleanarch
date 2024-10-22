package di

import (
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/app/usecases"
	"github.com/eternaleight/go-backend/interfaces/api/handlers"
	"gorm.io/gorm"
)

// すべてのハンドラを初期化し、返す
func InitializeHandlers(db *gorm.DB) (*handlers.AuthHandler, *handlers.ProductHandler, *handlers.PurchaseHandler, *handlers.UserHandler, *handlers.PostHandler) {

	// ストアのインスタンスを作成
	authStore := stores.NewAuthStore(db)
	productStore := stores.NewProductStore(db)
	purchaseStore := stores.NewPurchaseStore(db)
	userStore := stores.NewUserStore(db)
	postStore := stores.NewPostStore(db)

	// ユースケースのインスタンスを作成
	authUsecases := usecases.NewAuthUsecases(authStore)
	productUsecases := usecases.NewProductUsecases(productStore)
	purchaseUsecases := usecases.NewPurchaseUsecases(purchaseStore)
	userUsecases := usecases.NewUserUsecases(userStore)
	postUsecases := usecases.NewPostUsecases(postStore)

	// ハンドラのインスタンスを作成
	authHandler := handlers.NewAuthHandler(authUsecases)
	productHandler := handlers.NewProductHandler(productUsecases)
	purchaseHandler := handlers.NewPurchaseHandler(purchaseUsecases)
	userHandler := handlers.NewUserHandler(userUsecases)
	postHandler := handlers.NewPostHandler(postUsecases)

	return authHandler, productHandler, purchaseHandler, userHandler, postHandler
}

