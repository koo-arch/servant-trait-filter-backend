package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/koo-arch/servant-trait-filter-backend/internal/httperr"
)

func (h *Handler) FetchListClasses(c *gin.Context) {
	ctx := c.Request.Context()

	classes, err := h.Class.GetAllClasses(ctx)
	if err != nil {
		httperr.Write(c, err)
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h *Handler) FetchListAttributes(c *gin.Context) {
	ctx := c.Request.Context()

	attributes, err := h.Attribute.GetAllAttributes(ctx)
	if err != nil {
		httperr.Write(c, err)
		return
	}

	c.JSON(http.StatusOK, attributes)
}

func (h *Handler) FetchListOrderAlignments(c *gin.Context) {
	ctx := c.Request.Context()

	orderAlignments, err := h.OrderAlign.GetAllOrderAlignments(ctx)
	if err != nil {
		httperr.Write(c, err)
		return
	}

	c.JSON(http.StatusOK, orderAlignments)
}

func (h *Handler) FetchListMoralAlignments(c *gin.Context) {
	ctx := c.Request.Context()

	moralAlignments, err := h.MoralAlign.GetAllMoralAlignments(ctx)
	if err != nil {
		httperr.Write(c, err)
		return
	}

	c.JSON(http.StatusOK, moralAlignments)
}