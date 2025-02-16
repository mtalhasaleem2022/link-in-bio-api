package repositories

import (
	"context"
	"link-in-bio-api/internal/models"
	"time"
)

// MockLinkRepository is a mock implementation of the LinkRepositoryInterface.
type MockLinkRepository struct {
	Links  map[string]models.Link
	Visits []models.Visit
}

// NewMockLinkRepository creates a new instance of MockLinkRepository.
func NewMockLinkRepository() *MockLinkRepository {
	return &MockLinkRepository{
		Links:  make(map[string]models.Link),
		Visits: []models.Visit{},
	}
}

func (m *MockLinkRepository) CreateLink(ctx context.Context, link *models.Link) error {
	m.Links[link.ID] = *link
	return nil
}

func (m *MockLinkRepository) UpdateLink(ctx context.Context, link *models.Link) error {
	if _, exists := m.Links[link.ID]; !exists {
		return nil
	}
	m.Links[link.ID] = *link
	return nil
}

func (m *MockLinkRepository) DeleteLink(ctx context.Context, id string) error {
	if _, exists := m.Links[id]; !exists {
		return nil
	}
	delete(m.Links, id)
	return nil
}

func (m *MockLinkRepository) GetLink(ctx context.Context, id string) (*models.Link, error) {
	if link, exists := m.Links[id]; exists {
		return &link, nil
	}
	return nil, nil
}

func (m *MockLinkRepository) IncrementClicks(ctx context.Context, linkID string, ipAddress string) error {
	if link, exists := m.Links[linkID]; exists {
		link.Clicks++
		m.Links[linkID] = link
		m.Visits = append(m.Visits, models.Visit{
			LinkID:    linkID,
			IPAddress: ipAddress,
			Timestamp: time.Now(),
		})
	}
	return nil
}
