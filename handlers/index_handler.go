package handlers

import (
	"github.com/Techeer-Hogwarts/search-cron/models"
	"github.com/Techeer-Hogwarts/search-cron/services"
	"github.com/gin-gonic/gin"
)

type IndexHandler struct {
	service services.IndexService
}

func NewIndexHandler(service services.IndexService) *IndexHandler {
	return &IndexHandler{service}
}

// CreateIndex creates index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Query Meilisearch and return results
// @Tags search
// @Accept json
// @Produce json
// @Param index query string true "name of index"
// @Param query query string true "Search query string"
// @Param limit query int false "Number of results to return (default 10)"
// @Param offset query int false "Offset for pagination (default 0)"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /search [get]
func (h *IndexHandler) CreateIndex(c *gin.Context) {
	// Create index in meilisearch
	var post *models.BaicIndex
	err := h.service.CreateIndex(post)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Index created successfully"})
}
