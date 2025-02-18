package repositories

import (
	"database/sql"

	"github.com/meilisearch/meilisearch-go"
)

type Repository struct {
	IndexRepository IndexRepository
}

func NewRepository(db *sql.DB, meiliclient *meilisearch.ServiceManager) *Repository {
	return &Repository{
		IndexRepository: NewIndexRepository(db, meiliclient),
	}
}
