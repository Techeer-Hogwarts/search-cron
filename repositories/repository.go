package repositories

import (
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

type Repository struct {
	IndexRepository IndexRepository
}

func NewRepository(db *gorm.DB, meiliclient *meilisearch.ServiceManager) *Repository {
	return &Repository{
		IndexRepository: NewIndexRepository(db, meiliclient),
	}
}
