package services

import (
	"context"
	"fmt"
	"link-in-bio-api/config"
	"link-in-bio-api/internal/models"
	"link-in-bio-api/internal/repositories"
	"os"
	"strconv"
	"sync"
	"time"
)

type LinkService struct {
	repo         repositories.LinkRepositoryInterface
	clickChannel chan ClickRequest
	wg           sync.WaitGroup
	Config       *config.Config
}

type ClickRequest struct {
	LinkID    string
	IPAddress string
}

func NewLinkService(repo repositories.LinkRepositoryInterface, cfg *config.Config) *LinkService {
	clickChannel := make(chan ClickRequest, 10000) // Increased buffer for burst handling
	service := &LinkService{
		repo:         repo,
		clickChannel: clickChannel,
		Config:       cfg,
	}

	// Start multiple worker Goroutines for parallel click processing
	numWorkers, err := strconv.Atoi(os.Getenv("WORKER_COUNT")) // Tune this based on load testing

	if err != nil {
		numWorkers = 10
	}

	for i := 0; i < numWorkers; i++ {
		service.wg.Add(1)
		go func(workerID int) {
			defer service.wg.Done()
			for req := range clickChannel {
				_ = repo.IncrementClicks(context.Background(), req.LinkID, req.IPAddress)
			}
		}(i)
	}

	return service
}

func (s *LinkService) CreateLink(ctx context.Context, link *models.Link) error {
	return s.repo.CreateLink(ctx, link)
}

func (s *LinkService) UpdateLink(ctx context.Context, link *models.Link) error {
	return s.repo.UpdateLink(ctx, link)
}

func (s *LinkService) DeleteLink(ctx context.Context, id string) error {
	return s.repo.DeleteLink(ctx, id)
}

func (s *LinkService) GetLink(ctx context.Context, id string) (*models.Link, error) {
	return s.repo.GetLink(ctx, id)
}

func (s *LinkService) TrackClick(ctx context.Context, linkID string, ipAddress string) error {

	select {
	case s.clickChannel <- ClickRequest{LinkID: linkID, IPAddress: ipAddress}:
		return nil
	default:
		return fmt.Errorf("failed to increment clicks for linkID %s: %v", linkID, time.Now())
	}

}

func (s *LinkService) StopClickProcessing() {
	close(s.clickChannel) // Ensure the channel is closed before waiting
	s.wg.Wait()
}
