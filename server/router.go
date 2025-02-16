package server

import (
	docs "github.com/Techeer-Hogwarts/search-cron/docs"
	"github.com/Techeer-Hogwarts/search-cron/handlers"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// var (
// 	indexAccessCounter = promauto.NewCounterVec(
// 		prometheus.CounterOpts{
// 			Name: "",
// 			Help: "Total number of times search handler was accessed",
// 		},
// 		[]string{"status"}, // Success or failure status
// 	)

// 	searchDurationHistogram = promauto.NewHistogramVec(
// 		prometheus.HistogramOpts{
// 			Name:    "search_processing_duration_seconds",
// 			Help:    "Time taken to process search requests",
// 			Buckets: prometheus.DefBuckets, // Default Prometheus buckets
// 		},
// 		[]string{"status"},
// 	)
// )

// 라우터 설정
func setupRouter(handler *handlers.Handler) *gin.Engine {
	router := gin.New()
	router.Use(gin.Recovery())
	apiGroup := router.Group("/api/v1")

	docs.SwaggerInfo.Title = "Techeerzip Search API"
	docs.SwaggerInfo.Description = "Search Engine API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v2"
	{
		// User routes
		userGroup := apiGroup.Group("/index")
		{
			userGroup.POST("", handler.IndexHandler.CreateIndex)
		}
	}
	swagger := router.Group("/swagger")
	swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))
	return router
}
