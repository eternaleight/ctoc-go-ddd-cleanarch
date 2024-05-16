package main

import (
	"github.com/eternaleight/go-backend/config"
	"github.com/eternaleight/go-backend/interfaces"
)

func main() {
	dsn := config.LoadConfig() // dsn Data Source Name
	db := config.InitializeDatabase(dsn)

	// ルーターを設定
	r := interfaces.SetupRouter(db)
	r.Run(":8001")
}
