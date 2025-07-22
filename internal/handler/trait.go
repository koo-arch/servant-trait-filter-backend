package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchListTraits(c *gin.Context) {
	ctx := c.Request.Context()

	traits, err := h.Trait.GetAllTraits(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch traits"})
		return
	}

	c.JSON(http.StatusOK, traits)
}