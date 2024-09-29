package main

import (
	"log"

	"github.com/luytbq/astrio-secret-manager/cmd/api"
	"github.com/luytbq/astrio-secret-manager/config"
	"github.com/luytbq/astrio-secret-manager/internal/database"
)

func main() {
	db, err := database.NewPosgresDB()
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(config.App.SERVER_PORT, config.App.SERVER_API_PREFIX, db)

	server.Run()
}
