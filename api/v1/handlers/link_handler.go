package handlers

import (
	"context"
	"net/http"

	"link-in-bio-api/internal/models"
	"link-in-bio-api/internal/services"

	"github.com/gin-gonic/gin"
)

type LinkHandler struct {
	service *services.LinkService
}

func NewLinkHandler(service *services.LinkService) *LinkHandler {
	return &LinkHandler{service: service}
}

func (h *LinkHandler) CreateLink(c *gin.Context) {
	// Apply timeout from .env file before calling service
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.service.Config.RequestTimeout)
	defer cancel()
	var link models.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.CreateLink(ctx, &link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, link)
}

func (h *LinkHandler) UpdateLink(c *gin.Context) {
	// Apply timeout from .env file before calling service
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.service.Config.RequestTimeout)
	defer cancel()

	id := c.Param("id")
	var link models.Link
	if err := c.ShouldBindJSON(&link); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	link.ID = id
	if err := h.service.UpdateLink(ctx, &link); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, link)
}

func (h *LinkHandler) DeleteLink(c *gin.Context) {

	// Apply timeout from .env file before calling service
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.service.Config.RequestTimeout)
	defer cancel()

	id := c.Param("id")

	if err := h.service.DeleteLink(ctx, id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted"})
}

func (h *LinkHandler) GetLink(c *gin.Context) {

	// Apply timeout from .env file before calling service
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.service.Config.RequestTimeout)
	defer cancel()
	id := c.Param("id")

	link, err := h.service.GetLink(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, link)
}

func (h *LinkHandler) TrackClick(c *gin.Context) {
	// Apply timeout from .env file before calling service
	ctx, cancel := context.WithTimeout(c.Request.Context(), h.service.Config.RequestTimeout)
	defer cancel()
	linkID := c.Param("id")
	ipAddress := c.ClientIP()

	if err := h.service.TrackClick(ctx, linkID, ipAddress); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Click tracked"})
}
