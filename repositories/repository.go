package repositories

import "gorm.io/gorm"

type Repository struct {
	IndexRepository IndexRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		IndexRepository: NewIndexRepository(db),
	}
}
