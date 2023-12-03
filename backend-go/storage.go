package main

import (
	"context"
	"net/url"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Storage interface {
	CreateAccount(*Account) error
	GetRecipeById(int) (any, error)
	GetRecipes(int, url.Values) ([]RecipeDetail, error)
}

type Neo4jStore struct {
	db  neo4j.SessionWithContext
	ctx context.Context
}

func NewNeo4jStore(ctx context.Context) (*Neo4jStore, error) {

	db, err := neo4j.NewDriverWithContext("neo4j://localhost:7687", neo4j.BasicAuth("neo4j", "Wawa02290.", ""))

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

func (s *Neo4jStore) CreateAccount(acc *Account) error {

	resp, err := s.db.Run(s.ctx, "CREATE (u:User {name: $name, mail: $mail}) RETURN u",
		map[string]interface{}{
			"name": acc.FirstName + " " + acc.LastName,
			"mail": acc.Mail},
	)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil

}

func (s *Neo4jStore) GetRecipeById(id int) (any, error) {

	resp, err := s.db.Run(s.ctx, `
			MATCH (recipe:Recipe {idRecipe: $id})
			MATCH (recipe)-[:CONTAINS]->(step:Step)
			MATCH (recipe)-[:INGREDIENTS]->(ingredient:Ingredient)
			RETURN recipe, ingredient, step
			ORDER BY step.step ASC;
		`,
		map[string]interface{}{
			"id": id},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		data, err := resp.Collect(s.ctx)

		if err != nil {
			return nil, err
		}

		recipe := createRecipe(data, []string{"recipe", "step", "ingredient"})

		if err != nil {
			return nil, err
		}

		return recipe, nil
	}

	return nil, err
}

func (s *Neo4jStore) GetRecipes(page int, query url.Values) ([]RecipeDetail, error) {
	resp, err := s.db.Run(s.ctx, "MATCH (n:Recipe) RETURN n SKIP $page LIMIT $limit",
		map[string]interface{}{
			"page":  page * 10,
			"limit": 11,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		// create a array of recipes
		recipeList := make([]RecipeDetail, 0)

		for resp.Next(s.ctx) {
			recipe := createRecipeDetail(*resp.Record(), "n")

			recipeList = append(recipeList, recipe)
		}

		return recipeList, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, err
}

func extractProperty(record neo4j.Record, key string, propertyName string) interface{} {
	return record.AsMap()[key].(neo4j.Node).Props[propertyName]
}

func createRecipeDetail(record neo4j.Record, key string) RecipeDetail {
	return RecipeDetail{
		Difficulty: extractProperty(record, key, "difficulty").(string),
		Quantity:   extractProperty(record, key, "quantity").(string),
		Price:      extractProperty(record, key, "price").(string),
		Name:       extractProperty(record, key, "name").(string),
		IdRecipe:   extractProperty(record, key, "idRecipe").(int64),
	}
}

func createRecipeStep(record []*neo4j.Record, key string) RecipeStep {
	recipe := make(RecipeStep)
	for _, r := range record {
		recipeId := extractProperty(*r, key, "step").(int64)

		recipe[recipeId] = struct {
			Step string
		}{
			Step: extractProperty(*r, key, "name").(string),
		}
	}
	return recipe
}

func createRecipeIngredient(record []*neo4j.Record, key string) RecipeIngredients {
	recipe := make(RecipeIngredients)

	for _, r := range record {
		recipeId := extractProperty(*r, key, "idIngredient").(int64)

		recipe[recipeId] = IngredientInfo{
			Name:       extractProperty(*r, key, "name").(string),
			URLPicture: extractProperty(*r, key, "urlPicture").(string),
		}
	}

	return recipe
}

func createRecipe(record []*neo4j.Record, key []string) Recipe {
	detailRecord := record[0]
	return Recipe{
		RecipeDetail:      createRecipeDetail(*detailRecord, key[0]),
		RecipeStep:        createRecipeStep(record, key[1]),
		RecipeIngredients: createRecipeIngredient(record, key[2]),
	}
}
