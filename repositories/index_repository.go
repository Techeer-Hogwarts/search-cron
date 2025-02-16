package repositories

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"gorm.io/gorm"
)

type IndexRepository interface {
	CreateIndex(post *models.BaicIndex) error
}

type indexRepository struct {
	db *gorm.DB
}

func NewIndexRepository(db *gorm.DB) IndexRepository {
	return &indexRepository{db}
}

func (r *indexRepository) CreateIndex(post *models.BaicIndex) error {
	return r.db.Create(post).Error // this should connect query db and make index in meilisearch
}
