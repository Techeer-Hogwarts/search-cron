package server

import (
	"fmt"

	"github.com/Techeer-Hogwarts/search-cron/database"
	"github.com/Techeer-Hogwarts/search-cron/handlers"
	"github.com/Techeer-Hogwarts/search-cron/repositories"
	"github.com/Techeer-Hogwarts/search-cron/services"
)

func Start(port string) {
	db := database.SetupDatabase()
	repo := repositories.NewRepository(db)
	service := services.NewService(repo)
	handler := handlers.NewHandler(service)

	router := setupRouter(handler)
	router.Run(fmt.Sprintf(":%s", port))
}
