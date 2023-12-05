package main

import (
	"context"
	"log"
)

func main() {

	ctx := context.Background()
	err := DotEnvInit()

	if err != nil {
		log.Fatal(err)
	}

	err = initEnvVar()

	if err != nil {
		log.Fatal(err)
	}


	store, err := NewNeo4jStore(ctx)

	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3001", store)
	server.Run()
}
