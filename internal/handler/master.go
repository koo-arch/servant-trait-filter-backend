package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) FetchListClasses(c *gin.Context) {
	ctx := c.Request.Context()

	classes, err := h.Class.GetAllClasses(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch classes"})
		return
	}

	c.JSON(http.StatusOK, classes)
}

func (h *Handler) FetchListAttributes(c *gin.Context) {
	ctx := c.Request.Context()

	attributes, err := h.Attribute.GetAllAttributes(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attributes"})
		return
	}

	c.JSON(http.StatusOK, attributes)
}

func (h *Handler) FetchListOrderAlignments(c *gin.Context) {
	ctx := c.Request.Context()

	orderAlignments, err := h.OrderAlign.GetAllOrderAlignments(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch order alignments"})
		return
	}

	c.JSON(http.StatusOK, orderAlignments)
}

func (h *Handler) FetchListMoralAlignments(c *gin.Context) {
	ctx := c.Request.Context()

	moralAlignments, err := h.MoralAlign.GetAllMoralAlignments(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch moral alignments"})
		return
	}

	c.JSON(http.StatusOK, moralAlignments)
}