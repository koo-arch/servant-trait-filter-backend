package main

import (
	"log"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/pkg/config"
	"github.com/koo-arch/servant-trait-filter-backend/ent"
	_ "github.com/koo-arch/servant-trait-filter-backend/ent/runtime" // Import the generated ent runtime
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	config.LoadEnv()

	databaseURL := config.GetEnv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	client, err := ent.Open("postgres", databaseURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer func() {
		if err := client.Close(); err != nil {
			log.Fatalf("failed closing connection to postgres: %v", err)
		}
	}()

	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	router.Run(":8080")
}