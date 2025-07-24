package main

import (
	"context"
	"log"

	"github.com/koo-arch/servant-trait-filter-backend/internal/atlas"
	"github.com/koo-arch/servant-trait-filter-backend/internal/etl"
	"github.com/koo-arch/servant-trait-filter-backend/internal/scheduler"
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

	// Atlas APIとの同期
	atlas := atlas.NewClient()
	etl := etl.NewSyncAtlas(client, atlas, repos)
	if err := etl.Sync(ctx); err != nil {
		log.Fatalf("failed to sync atlas api: %v", err)
	}
	// スケジューラーの設定
	sched := scheduler.NewScheduler(etl)
	sched.SetupJobs(ctx)
	sched.Start()
	defer sched.Stop()

	// ルーターの初期化
	router := InitRouter(services)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}