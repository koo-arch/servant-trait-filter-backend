package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/internal/di"
)

type Handler struct {
	*di.Services
}

func NewHandler(services *di.Services) *Handler {
	return &Handler{
		services,
	}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api/v1")
	{
		servant := api.Group("/servants")
		{
			servant.GET("", h.FetchListServants)
			servant.POST("/search", h.SearchServants)
		}

		trait := api.Group("/traits")
		{
			trait.GET("", h.FetchListTraits)
		}

		master := api.Group("/master")
		{
			master.GET("/classes", h.FetchListClasses)
			master.GET("/attributes", h.FetchListAttributes)
			master.GET("/order-alignments", h.FetchListOrderAlignments)
			master.GET("/moral-alignments", h.FetchListMoralAlignments)
		}
	}
}