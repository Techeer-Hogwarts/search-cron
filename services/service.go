package services

import "github.com/Techeer-Hogwarts/search-cron/repositories"

// Service groups all individual services.
type Service struct {
	IndexService IndexService
}

// NewService creates a new instance of Service with all required services.
func NewService(repo *repositories.Repository) *Service {
	return &Service{
		IndexService: NewIndexService(repo.IndexRepository),
	}
}
