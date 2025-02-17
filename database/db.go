package database

import (
	"fmt"
	"log"

	"github.com/Techeer-Hogwarts/search-cron/config"
	"github.com/meilisearch/meilisearch-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupDatabase() *gorm.DB {
	config.LoadEnvFile(".env")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetEnvVarAsString("DB_HOST", "localhost"),
		config.GetEnvVarAsString("DB_USER", "postgres"),
		config.GetEnvVarAsString("DB_PASSWORD", "postgres"),
		config.GetEnvVarAsString("DB_NAME", "postgres"),
		config.GetEnvVarAsString("DB_PORT", "5432"),
		config.GetEnvVarAsString("DB_SSLMODE", "disable"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected")
	return db
}

// Initialize Meilisearch client
func InitMeilisearch() *meilisearch.ServiceManager {
	clientAddr := config.GetEnvVarAsString("MEILISEARCH_ADDR", "http://localhost:7700")
	clientAPIKey := config.GetEnvVarAsString("MEILISEARCH_API_KEY", "masterKey")
	client := meilisearch.New(clientAddr, meilisearch.WithAPIKey(clientAPIKey))
	MeiliClient := &client
	log.Println("Meilisearch client initialized")
	return MeiliClient
}
