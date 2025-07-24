package main

import (
	"strings"

	"github.com/koo-arch/servant-trait-filter-backend/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/pkg/config"
	"github.com/koo-arch/servant-trait-filter-backend/internal/di"
	"github.com/koo-arch/servant-trait-filter-backend/internal/handler"

)

func InitRouter(services *di.Services) *gin.Engine {
	r := gin.Default()

	// corsの設定
	allowedOrigins := strings.Split(config.GetEnv("CORS_ALLOW_ORIGINS"), ",")
	r.Use(middleware.CORS(allowedOrigins))
	// ハンドラーの初期化
	h := handler.NewHandler(services)
	// ルーティングの登録
	h.RegisterRoutes(r)
	return r
}