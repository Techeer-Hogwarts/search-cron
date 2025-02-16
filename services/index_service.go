package services

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/Techeer-Hogwarts/search-cron/repositories"
)

type IndexService interface {
	CreateIndex(post *models.BaicIndex) error
}

type indexService struct {
	repo repositories.IndexRepository
}

func NewIndexService(repo repositories.IndexRepository) IndexService {
	return &indexService{repo}
}

func (s *indexService) CreateIndex(post *models.BaicIndex) error {
	return s.repo.CreateIndex(post)
}
