package main

import (
	"github.com/Techeer-Hogwarts/search-cron/config"
	"github.com/Techeer-Hogwarts/search-cron/server"
)

// 앱에서 가장 먼저 실행되는 함수
func main() {
	port := config.GetEnvVarAsString("PORT", "8081")
	server.Start(port)
}
