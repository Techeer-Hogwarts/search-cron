package server

import (
	docs "github.com/Techeer-Hogwarts/search-cron/docs"
	"github.com/Techeer-Hogwarts/search-cron/handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	indexCreationCounter = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "index_creation_count",
			Help: "Total number of times index creation was accessed",
		},
		[]string{"status", "index_type"}, // Success/failure status and type of index (user, project, study, etc.)
	)

	indexCreationDurationHistogram = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "index_creation_duration_seconds",
			Help:    "Time taken to create index",
			Buckets: prometheus.DefBuckets, // Default Prometheus buckets
		},
		[]string{"status", "index_type"},
	)
)

// 라우터 설정
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	apiGroup := router.Group("/api/v1")

	docs.SwaggerInfo.Title = "Techeerzip Search API"
	docs.SwaggerInfo.Description = "Search Engine API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	{
		indexGroup := apiGroup.Group("/index")
		{
			indexGroup.POST("/user", func(c *gin.Context) {
				handler.IndexHandler.CreateUserIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/project", func(c *gin.Context) {
				handler.IndexHandler.CreateProjectIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/study", func(c *gin.Context) {
				handler.IndexHandler.CreateStudyIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/blog", func(c *gin.Context) {
				handler.IndexHandler.CreateBlogIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/resume", func(c *gin.Context) {
				handler.IndexHandler.CreateResumeIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/session", func(c *gin.Context) {
				handler.IndexHandler.CreateSessionIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/event", func(c *gin.Context) {
				handler.IndexHandler.CreateEventIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			indexGroup.POST("/stack", func(c *gin.Context) {
				handler.IndexHandler.CreateStackIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
		}
		syncGroup := apiGroup.Group("/sync")
		{
			syncGroup.POST("/all", func(c *gin.Context) {
				handler.IndexHandler.SyncAllIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
			syncGroup.POST("/new", func(c *gin.Context) {
				handler.IndexHandler.SyncNewIndexHandler(c, indexCreationCounter, indexCreationDurationHistogram)
			})
		}
		deleteGroup := apiGroup.Group("/delete")
		{
			deleteGroup.DELETE("/index", handler.IndexHandler.DeleteIndexHandler)
			deleteGroup.DELETE("/index/all", handler.IndexHandler.DeleteAllIndexHandler)
			deleteGroup.DELETE("/document", handler.IndexHandler.DeleteDocumentHandler)
			deleteGroup.DELETE("/document/all", handler.IndexHandler.DeleteAllDocumentHandler)
		}
	}
	swagger := apiGroup.Group("/swagger")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}
