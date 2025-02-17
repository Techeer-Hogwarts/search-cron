package repositories

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/gorm"
)

type IndexRepository interface {
	CreateUserIndex(post *models.UserIndex) error
	CreateProjectIndex(post *models.ProjectIndex) error
	CreateStudyIndex(post *models.StudyIndex) error
	CreateBlogIndex(post *models.BlogIndex) error
	CreateResumeIndex(post *models.ResumeIndex) error
	CreateSessionIndex(post *models.SessionIndex) error
	CreateEventIndex(post *models.EventIndex) error
	DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error)
	DeleteIndex(index string) (*meilisearch.TaskInfo, error)
	DeleteAllIndex() (*meilisearch.TaskInfo, error)
	DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error)
}

type indexRepository struct {
	db          *gorm.DB
	meiliclient *meilisearch.ServiceManager
}

func NewIndexRepository(db *gorm.DB, meiliclient *meilisearch.ServiceManager) IndexRepository {
	return &indexRepository{db, meiliclient}
}

func (r *indexRepository) CreateUserIndex(post *models.UserIndex) error {
	r.db.Create(post)
	return nil
}

func (r *indexRepository) CreateProjectIndex(post *models.ProjectIndex) error {
	return nil
}

func (r *indexRepository) CreateStudyIndex(post *models.StudyIndex) error {
	return nil
}

func (r *indexRepository) CreateBlogIndex(post *models.BlogIndex) error {
	return nil
}

func (r *indexRepository) CreateResumeIndex(post *models.ResumeIndex) error {
	return nil
}

func (r *indexRepository) CreateSessionIndex(post *models.SessionIndex) error {
	return nil
}

func (r *indexRepository) CreateEventIndex(post *models.EventIndex) error {
	return nil
}

func (r *indexRepository) DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index(index).DeleteDocument(id)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) DeleteIndex(index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).DeleteIndex(index)
	if err != nil {
		return info, err
	}
	return info, nil
}

func (r *indexRepository) DeleteAllIndex() (*meilisearch.TaskInfo, error) {
	param := &meilisearch.IndexesQuery{}
	index, err := (*r.meiliclient).ListIndexes(param)
	if err != nil {
		return nil, err
	}
	for _, v := range index.Results {
		info, err := (*r.meiliclient).DeleteIndex(v.UID)
		if err != nil {
			return info, err
		}
	}
	return nil, nil
}

func (r *indexRepository) DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error) {
	info, err := (*r.meiliclient).Index(index).DeleteAllDocuments()
	if err != nil {
		return info, err
	}
	return info, nil
}
