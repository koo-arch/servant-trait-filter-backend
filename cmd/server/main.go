package main

import (
	"log"
	"context"
	
	"github.com/koo-arch/servant-trait-filter-backend/pkg/config"
	_ "github.com/koo-arch/servant-trait-filter-backend/ent/runtime" // Import the generated ent runtime
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	config.LoadEnv()

	databaseURL := config.GetEnv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	ctx := context.Background()
	// データベース接続
	client := connectDB(ctx, databaseURL)
	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()

	// リポジトリとサービスの初期化
	repos := InitRepos(client)
	services := InitServices(repos)

	// ルーターの初期化
	router := InitRouter(services)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}