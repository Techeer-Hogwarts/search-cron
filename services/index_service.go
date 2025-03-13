package services

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/Techeer-Hogwarts/search-cron/repositories"
	"github.com/meilisearch/meilisearch-go"
)

type IndexService interface {
	CreateIndex(post *models.BaiscIndex) error
	CreateUserIndex(post *models.UserIndex) (*meilisearch.TaskInfo, error)
	CreateProjectIndex(post *models.ProjectIndex) (*meilisearch.TaskInfo, error)
	CreateStudyIndex(post *models.StudyIndex) (*meilisearch.TaskInfo, error)
	CreateBlogIndex(post *models.BlogIndex) (*meilisearch.TaskInfo, error)
	CreateResumeIndex(post *models.ResumeIndex) (*meilisearch.TaskInfo, error)
	CreateSessionIndex(post *models.SessionIndex) (*meilisearch.TaskInfo, error)
	CreateEventIndex(post *models.EventIndex) (*meilisearch.TaskInfo, error)
	CreateStackIndex() (*meilisearch.TaskInfo, error)
	CreateAllIndex(condition string) (*meilisearch.TaskInfo, error)
	DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error)
	DeleteIndex(index string) (*meilisearch.TaskInfo, error)
	DeleteAllIndex() (*meilisearch.TaskInfo, error)
	DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error)
}

type indexService struct {
	repo repositories.IndexRepository
}

func NewIndexService(repo repositories.IndexRepository) IndexService {
	return &indexService{repo}
}

func (s *indexService) CreateIndex(post *models.BaiscIndex) error {
	var temp models.UserIndex
	_, err := s.repo.CreateUserIndex(&temp)
	if err != nil {
		return err
	}
	return nil
}

func (s *indexService) CreateUserIndex(post *models.UserIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateUserIndex(post)
}

func (s *indexService) CreateProjectIndex(post *models.ProjectIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateProjectIndex(post)
}

func (s *indexService) CreateStudyIndex(post *models.StudyIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateStudyIndex(post)
}

func (s *indexService) CreateBlogIndex(post *models.BlogIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateBlogIndex(post)
}

func (s *indexService) CreateResumeIndex(post *models.ResumeIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateResumeIndex(post)
}

func (s *indexService) CreateSessionIndex(post *models.SessionIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateSessionIndex(post)
}

func (s *indexService) CreateEventIndex(post *models.EventIndex) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateEventIndex(post)
}

func (s *indexService) CreateStackIndex() (*meilisearch.TaskInfo, error) {
	return s.repo.CreateStackIndex()
}

func (s *indexService) CreateAllIndex(condition string) (*meilisearch.TaskInfo, error) {
	return s.repo.CreateAllIndex(condition)
}

func (s *indexService) DeleteIndexDocument(id, index string) (*meilisearch.TaskInfo, error) {
	return s.repo.DeleteIndexDocument(id, index)
}

func (s *indexService) DeleteIndex(index string) (*meilisearch.TaskInfo, error) {
	return s.repo.DeleteIndex(index)
}

func (s *indexService) DeleteAllIndex() (*meilisearch.TaskInfo, error) {
	return s.repo.DeleteAllIndex()
}

func (s *indexService) DeleteAllDocuments(index string) (*meilisearch.TaskInfo, error) {
	return s.repo.DeleteAllDocuments(index)
}
