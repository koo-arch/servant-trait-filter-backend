package handler

import (
	"net/http"

	"github.com/koo-arch/servant-trait-filter-backend/internal/search"
	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchListServants(c *gin.Context) {
	ctx := c.Request.Context()

	servants, err := h.Servant.GetAllServants(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch servants"})
		return
	}

	c.JSON(http.StatusOK, servants)
}

func (h *Handler) SearchServants(c *gin.Context) {
	ctx := c.Request.Context()

	var req search.ServantSearchQuery
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	response, err := h.Servant.Search(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search servants"})
		return
	}

	c.JSON(http.StatusOK, response)
}