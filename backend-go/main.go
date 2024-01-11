package main

import (
	"backend/api"
	"backend/config"
	"backend/storage"
	"context"
	"log"
	"os"
)

func main() {

	ctx := context.Background()
	stage := os.Getenv("STAGE")
	if stage != "prod" {
		err := config.DotEnvInit()

		if err != nil {
			log.Fatal(err)
		}
	}

	err := config.InitEnvVar()

	if err != nil {
		log.Fatal(err)
	}

	store, err := storage.NewNeo4jStore(ctx)

	if err != nil {
		log.Fatal(err)
	}

	server := api.NewAPIServer(":3001", store)
	server.Run()
}
