package services

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/Techeer-Hogwarts/search-cron/repositories"
)

type IndexService interface {
	CreateIndex(post *models.BaiscIndex) error
}

type indexService struct {
	repo repositories.IndexRepository
}

func NewIndexService(repo repositories.IndexRepository) IndexService {
	return &indexService{repo}
}

func (s *indexService) CreateIndex(post *models.BaiscIndex) error {
	var temp models.UserIndex
	return s.repo.CreateUserIndex(&temp)
}
