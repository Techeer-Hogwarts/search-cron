package handlers

import "github.com/Techeer-Hogwarts/search-cron/services"

// Handler groups all individual handlers.
type Handler struct {
	IndexHandler *IndexHandler
}

// NewHandler creates a new instance of Handler with all required handlers.
func NewHandler(service *services.Service) *Handler {
	return &Handler{
		IndexHandler: NewIndexHandler(service.IndexService),
	}
}
