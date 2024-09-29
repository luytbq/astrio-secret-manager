package main

import (
	"log"

	"github.com/luytbq/astrio-secret-manager/internal/database"
	"github.com/luytbq/astrio-secret-manager/pkg/secret"
)

func main() {
	db, err := database.NewPosgresDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := secret.NewRepo(db)
	key, err := repo.GetKey()
	// key, err := repo.GetKeyById(1)
	// key, err := repo.NewKey()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("key: %+v", key)
}
