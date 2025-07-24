package middleware

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func CORS(origins []string) gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge: 12 * time.Hour,
	})
}