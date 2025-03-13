package handlers

import (
	"encoding/json"
	"log"
	"net/http"

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

// CreateUserIndexHandler creates user index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add User Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.UserIndex true "User Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/user [post]
func (h *IndexHandler) CreateUserIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var userIndex models.UserIndex

	if err := c.ShouldBindJSON(&userIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateUserIndex(&userIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("User Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateProjectIndexHandler creates project index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Project Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.ProjectIndex true "Project Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/project [post]
func (h *IndexHandler) CreateProjectIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var projectIndex models.ProjectIndex

	if err := c.ShouldBindJSON(&projectIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateProjectIndex(&projectIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Project Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateStudyIndexHandler creates study index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Study Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.StudyIndex true "Study Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/study [post]
func (h *IndexHandler) CreateStudyIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var studyIndex models.StudyIndex

	if err := c.ShouldBindJSON(&studyIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateStudyIndex(&studyIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Study Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateBlogIndexHandler creates blog index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Blog Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.BlogIndex true "Blog Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/blog [post]
func (h *IndexHandler) CreateBlogIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var blogIndex models.BlogIndex

	if err := c.ShouldBindJSON(&blogIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateBlogIndex(&blogIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Blog Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateResumeIndexHandler creates resume index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Resume Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.ResumeIndex true "Resume Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/resume [post]
func (h *IndexHandler) CreateResumeIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var resumeIndex models.ResumeIndex

	if err := c.ShouldBindJSON(&resumeIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateResumeIndex(&resumeIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Resume Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateSessionIndexHandler creates session index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Session Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.SessionIndex true "Session Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/session [post]
func (h *IndexHandler) CreateSessionIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var sessionIndex models.SessionIndex

	if err := c.ShouldBindJSON(&sessionIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateSessionIndex(&sessionIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Session Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateEventIndexHandler creates event index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Event Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Param index body models.EventIndex true "Event Index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/event [post]
func (h *IndexHandler) CreateEventIndexHandler(c *gin.Context) {
	// Create index in meilisearch
	var eventIndex models.EventIndex

	if err := c.ShouldBindJSON(&eventIndex); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.CreateEventIndex(&eventIndex)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Event Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// CreateStackIndexHandler creates stack index in meilisearch from request
// @Summary Creates index in meilisearch
// @Description Add Stack Index to Meilisearch
// @Tags index
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /index/stack [post]
func (h *IndexHandler) CreateStackIndexHandler(c *gin.Context) {
	info, err := h.service.CreateStackIndex()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Stack Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// DeleteIndexHandler deletes an index
// @Summary Delete index in meilisearch
// @Description Delete intended index in Meilisearch
// @Tags delete
// @Produce json
// @Param index query string true "name of index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /delete/index [delete]
func (h *IndexHandler) DeleteIndexHandler(c *gin.Context) {
	var indexReq models.BaiscIndex

	if err := c.ShouldBindQuery(&indexReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.DeleteIndex(indexReq.Index)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// DeleteDocumentHandler deletes an document in index
// @Summary Delete document in index in meilisearch
// @Description Delete intended document in Meilisearch
// @Tags delete
// @Produce json
// @Param index query string true "name of index"
// @Param id query string true "id of document"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /delete/document [delete]
func (h *IndexHandler) DeleteDocumentHandler(c *gin.Context) {
	var indexReq models.BaiscIndex

	if err := c.ShouldBindQuery(&indexReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.DeleteIndexDocument(indexReq.ID, indexReq.Index)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// DeleteAllIndexHandler deletes all index
// @Summary Delete all index in meilisearch
// @Description Delete all index in Meilisearch. Use at your own risk
// @Tags delete
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /delete/index/all [delete]
func (h *IndexHandler) DeleteAllIndexHandler(c *gin.Context) {
	info, err := h.service.DeleteAllIndex()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// DeleteAllDocumentHandler deletes all documents in index
// @Summary Delete all document in meilisearch
// @Description Delete all documents in an index in Meilisearch. Use at your own risk
// @Tags delete
// @Produce json
// @Param index query string true "name of index"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /delete/document/all [delete]
func (h *IndexHandler) DeleteAllDocumentHandler(c *gin.Context) {
	var indexReq models.BaiscIndex

	if err := c.ShouldBindQuery(&indexReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	info, err := h.service.DeleteAllDocuments(indexReq.Index)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// SyncAllIndexHandler syncs all index with database
// @Summary Sync all index in meilisearch
// @Description Sync all index with database
// @Tags sync
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sync/all [post]
func (h *IndexHandler) SyncAllIndexHandler(c *gin.Context) {
	info, err := h.service.CreateAllIndex("all")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}

// SyncNewIndexHandler syncs new index with database
// @Summary Sync new index in meilisearch
// @Description Sync new index with database
// @Tags sync
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /sync/new [post]
func (h *IndexHandler) SyncNewIndexHandler(c *gin.Context) {
	info, err := h.service.CreateAllIndex("new")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	jsonData, _ := json.MarshalIndent(info, "", "  ")
	log.Printf("Index created successfully: %v", string(jsonData))
	c.JSON(200, gin.H{"message": "Index created successfully", "info": info})
}
