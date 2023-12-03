package main

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	CreateAccount(*Account) error
	FindAccountByMail(string) error
	Login(LoginRequest) (*Account, error)
	GetRecipeById(int) (*APIResponse, error)
	GetRecipes(int) (*APIResponse, error)
	GetRecipesWithFilter(int, url.Values) (*APIResponse, error)
	GetMetadata(int, string, map[string]interface{}) (*Pagination, error)
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

func (s *Neo4jStore) FindAccountByMail(mail string) error {
	query := "MATCH (u:User {mail: $mail}) RETURN u"

	resp, err := s.db.Run(s.ctx, query,
		map[string]interface{}{
			"mail": mail},
	)

	if err != nil {

		return err
	}

	if resp.Err() != nil {
		return err
	}

	if resp.Next(s.ctx) {
		return fmt.Errorf("Email found")
	}

	return nil
}

func (s *Neo4jStore) CreateAccount(acc *Account) error {

	resp, err := s.db.Run(s.ctx, "CREATE (u:User {name: $name, mail: $mail, password: $password}) RETURN u",
		map[string]interface{}{
			"name":     acc.FirstName + " " + acc.LastName,
			"mail":     acc.Mail,
			"password": acc.EncryptedPassword,
		},
	)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil
}

func (s *Neo4jStore) Login(req LoginRequest) (*Account, error) {
	query := "MATCH (u:User {mail: $mail}) RETURN u"

	resp, err := s.db.Run(s.ctx, query,
		map[string]interface{}{
			"mail": req.Mail},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		data := extractProperty(*resp.Record(), "u", "password").(string)

		if err := bcrypt.CompareHashAndPassword([]byte(data), []byte(req.Password)); err != nil {
			return nil, fmt.Errorf("Invalid Credentials")
		}

		account := CreateAccount(*resp.Record(), "u")

		return &account, nil
	}

	return nil, err
}

func (s *Neo4jStore) GetRecipeById(id int) (*APIResponse, error) {

	query := `
			MATCH (recipe:Recipe {idRecipe: $id})
			MATCH (recipe)-[:CONTAINS]->(step:Step)
			MATCH (recipe)-[:INGREDIENTS]->(ingredient:Ingredient)
			RETURN recipe, ingredient, step
			ORDER BY step.step ASC;
	`

	params := map[string]interface{}{
		"id": id,
	}

	resp, err := s.db.Run(s.ctx, query,
		params,
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

		finalResult := APIResponse{
			Result: recipe,
		}

		return &finalResult, nil
	}

	return nil, err
}

func (s *Neo4jStore) GetRecipes(page int) (*APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n SKIP $page LIMIT $limit"

	params := map[string]interface{}{
		"page":  page * 10,
		"limit": 10,
	}

	resp, err := s.db.Run(s.ctx, query,
		params,
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
			recipe := CreateRecipeDetail(*resp.Record(), "n")

			recipeList = append(recipeList, recipe)
		}

		recipePagination, err := s.GetMetadata(page, query, params)

		if err != nil {
			return nil, err
		}

		finalResult := APIResponse{
			Result:     recipeList,
			Pagination: recipePagination,
		}

		return &finalResult, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (s *Neo4jStore) GetRecipesWithFilter(page int, query url.Values) (*APIResponse, error) {
	queryString, params := buildQueryAndParams(query, page)

	resp, err := s.db.Run(s.ctx, queryString,
		params,
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
			recipe := CreateRecipeDetail(*resp.Record(), "n")

			recipeList = append(recipeList, recipe)
		}

		recipePagination, err := s.GetMetadata(page, queryString, params)

		if err != nil {
			return nil, err
		}

		finalResult := APIResponse{
			Result:     recipeList,
			Pagination: recipePagination,
		}

		return &finalResult, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (s *Neo4jStore) GetMetadata(page int, query string, params map[string]interface{}) (*Pagination, error) {

	query = strings.ReplaceAll(query, "RETURN n", "RETURN count(n) as total")

	query = strings.ReplaceAll(query, "SKIP $page LIMIT $limit", "")

	query = strings.ReplaceAll(query, "UNION", "\n$1")

	resp, err := s.db.Run(s.ctx, query,
		params,
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		data := resp.Record().AsMap()["total"].(int64)

		pagination := CreatePagination(data, page)
		return &pagination, nil
	}

	return nil, nil
}
