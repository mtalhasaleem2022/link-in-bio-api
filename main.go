package main

import (
	"link-in-bio-api/api/v1/routes"
	"link-in-bio-api/config"
	"link-in-bio-api/internal/repositories"
	"link-in-bio-api/internal/services"
	"link-in-bio-api/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load configuration
	cfg := config.LoadConfig()

	// Initialize MongoDB repository
	repo := repositories.NewLinkRepository(cfg.MongoURI)

	// Initialize services with the repository
	linkService := services.NewLinkService(repo, cfg)

	// Ensure workers are stopped when app shuts down
	defer linkService.StopClickProcessing()
	// Setup Gin router
	r := gin.Default()

	// Middleware
	r.Use(middleware.Logger())

	// Setup routes
	routes.SetupLinkRoutes(r, linkService)

	// Start server
	r.Run(":" + cfg.Port) // Start server on the configured port
}
