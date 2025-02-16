package repositories

import (
	"context"
	"link-in-bio-api/internal/models"
)

// LinkRepositoryInterface defines the methods that a link repository must implement.
type LinkRepositoryInterface interface {
	CreateLink(ctx context.Context, link *models.Link) error
	UpdateLink(ctx context.Context, link *models.Link) error
	DeleteLink(ctx context.Context, id string) error
	GetLink(ctx context.Context, id string) (*models.Link, error)
	IncrementClicks(ctx context.Context, linkID string, ipAddress string) error
}
