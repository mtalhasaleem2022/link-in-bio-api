package services

import (
	"context"
	"testing"
	"time"

	"link-in-bio-api/config"
	"link-in-bio-api/internal/models"
	"link-in-bio-api/internal/repositories"
)

func TestCreateLink(t *testing.T) {

	// Create a mock repository
	mockRepo := repositories.NewMockLinkRepository()

	// Create a mock config with a timeout value
	mockConfig := &config.Config{
		RequestTimeout: 5 * time.Second, // Set a reasonable timeout
	}
	// Create the service with the mock repository
	service := NewLinkService(mockRepo, mockConfig)

	// Define a test link
	link := &models.Link{
		ID:        "1",
		Title:     "Test Link",
		URL:       "https://example.com",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Clicks:    0,
	}

	// Test creating the link
	err := service.CreateLink(context.Background(), link)
	if err != nil {
		t.Fatalf("Failed to create link: %v", err)
	}

	// Verify the link was created in the mock repository
	if _, exists := mockRepo.Links[link.ID]; !exists {
		t.Errorf("Link was not created in the repository")
	}
}

func TestTrackClick(t *testing.T) {
	// Create a mock repository
	mockRepo := repositories.NewMockLinkRepository()

	mockConfig := &config.Config{
		RequestTimeout: 5 * time.Second, // Set a reasonable timeout
	}
	// Create the service with the mock repository
	service := NewLinkService(mockRepo, mockConfig)

	// Define a test link
	link := &models.Link{
		ID:        "1",
		Title:     "Test Link",
		URL:       "https://example.com",
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Clicks:    0,
	}

	// Create the link first
	err := service.CreateLink(context.Background(), link)
	if err != nil {
		t.Fatalf("Failed to create link: %v", err)
	}

	// Track a click
	err = service.TrackClick(context.Background(), link.ID, "127.0.0.1")
	if err != nil {
		t.Fatalf("Failed to track click: %v", err)
	}

	// Wait for the click processing to complete
	service.StopClickProcessing()

	// Verify the click count was incremented
	if mockRepo.Links[link.ID].Clicks != 1 {
		t.Errorf("Click count was not incremented")
	}

	// Verify the visit was logged
	if len(mockRepo.Visits) != 1 {
		t.Errorf("Visit was not logged")
	}
}
