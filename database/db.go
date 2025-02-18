package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Techeer-Hogwarts/search-cron/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/meilisearch/meilisearch-go"
)

func SetupDatabase() *sql.DB {
	// config.LoadEnvFile(".env")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.GetEnvVarAsString("DB_HOST", "localhost"),
		config.GetEnvVarAsString("DB_USER", "postgres"),
		config.GetEnvVarAsString("DB_PASSWORD", "postgres"),
		config.GetEnvVarAsString("DB_NAME", "postgres"),
		config.GetEnvVarAsString("DB_PORT", "5432"),
		config.GetEnvVarAsString("DB_SSLMODE", "disable"),
	)
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	log.Println("Database connected")
	return db
}

// Initialize Meilisearch client
func InitMeilisearch() *meilisearch.ServiceManager {
	clientAddr := config.GetEnvVarAsString("MEILISEARCH_ADDR", "http://localhost:7700")
	clientAPIKey := config.GetEnvVarAsString("MEILISEARCH_API_KEY", "someapikey")
	client := meilisearch.New(clientAddr, meilisearch.WithAPIKey(clientAPIKey))
	MeiliClient := &client
	log.Println("Meilisearch client initialized")
	return MeiliClient
}
