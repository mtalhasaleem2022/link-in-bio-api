package routes

import (
	"link-in-bio-api/api/v1/handlers"
	"link-in-bio-api/internal/services"

	"github.com/gin-gonic/gin"
)

func SetupLinkRoutes(r *gin.Engine, linkService *services.LinkService) {
	linkHandler := handlers.NewLinkHandler(linkService)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/links", linkHandler.CreateLink)
		v1.PUT("/links/:id", linkHandler.UpdateLink)
		v1.DELETE("/links/:id", linkHandler.DeleteLink)
		v1.GET("/links/:id", linkHandler.GetLink)
		v1.GET("/visit/:id", linkHandler.TrackClick)
	}
}
