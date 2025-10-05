package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/internal/httperr"
)

func (h *Handler) FetchListTraits(c *gin.Context) {
	ctx := c.Request.Context()

	traits, err := h.Trait.GetAllTraits(ctx)
	if err != nil {
		httperr.Write(c, err)
		return
	}

	c.JSON(http.StatusOK, traits)
}