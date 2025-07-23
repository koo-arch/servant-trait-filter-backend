package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/internal/di"
	"github.com/koo-arch/servant-trait-filter-backend/internal/handler"
)

func InitRouter(services *di.Services) *gin.Engine {
	r := gin.Default()
	// ハンドラーの初期化
	h := handler.NewHandler(services)
	// ルーティングの登録
	h.RegisterRoutes(r)
	return r
}