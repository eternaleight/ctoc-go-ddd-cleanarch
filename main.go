package main

import (
	"github.com/eternaleight/go-backend/config"
	"github.com/eternaleight/go-backend/infra/stores"
	"github.com/eternaleight/go-backend/interfaces"
	"github.com/eternaleight/go-backend/interfaces/api/handlers"
)

func main() {
	dsn := config.LoadConfig() // dsn Data Source Name
	db := config.InitializeDatabase(dsn)

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

	// ルーターを設定
	r := interfaces.SetupRouter(authHandler, productHandler, purchaseHandler, userHandler, postHandler)

	r.Run(":8001")
}
