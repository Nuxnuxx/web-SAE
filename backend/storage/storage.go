package storage

import (
	"context"
	"os"

	"backend/config"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Storage interface {
	AuthStorage
	RecipeStorage
	SuggestionStorage
	ListStorage
}

type Neo4jStore struct {
	db  neo4j.SessionWithContext
	ctx context.Context
}

func NewNeo4jStore(ctx context.Context) (*Neo4jStore, error) {
	uri := os.Getenv("DB_URL")
	auth := neo4j.NoAuth()

	if config.Stage == "dev" {
		auth = neo4j.BasicAuth(os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), "")
	}

	db, err := neo4j.NewDriverWithContext(uri, auth)

	if err != nil {
		return nil, err
	}

	if err := db.VerifyConnectivity(ctx); err != nil {
		return nil, err
	}

	session := db.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})

	return &Neo4jStore{
		db:  session,
		ctx: ctx,
	}, nil
}
