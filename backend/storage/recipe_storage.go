package storage

import (
	"backend/types"
	"net/url"
	"strings"
)

type RecipeStorage interface {
	// Recipe
	GetRecipeById(int) (*types.APIResponse, error)
	GetRecipes(int) (*types.APIResponse, error)
	GetRecipesWithFilter(int, url.Values) (*types.APIResponse, error)
	GetMetadata(int, string, map[string]interface{}) (*types.Pagination, error)
}

func (s *Neo4jStore) GetRecipes(page int) (*types.APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n SKIP $page LIMIT $limit"

	params := map[string]interface{}{
		"page":  page * 9,
		"limit": 9,
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

	recipeList := make([]types.RecipeDetail, 0)

	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "n")

		recipeList = append(recipeList, recipe)
	}

	recipePagination, err := s.GetMetadata(page, query, params)

	if err != nil {
		return nil, err
	}

	finalResult := types.APIResponse{
		Result:     recipeList,
		Pagination: recipePagination,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetRecipesWithFilter(page int, query url.Values) (*types.APIResponse, error) {
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

	recipeList := make([]types.RecipeDetail, 0)

	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "n")

		recipeList = append(recipeList, recipe)
	}

	recipePagination, err := s.GetMetadata(page, queryString, params)

	if err != nil {
		return nil, err
	}

	finalResult := types.APIResponse{
		Result:     recipeList,
		Pagination: recipePagination,
	}

	return &finalResult, nil

}

func (s *Neo4jStore) GetMetadata(page int, query string, params map[string]interface{}) (*types.Pagination, error) {

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

func (s *Neo4jStore) GetRecipeById(id int) (*types.APIResponse, error) {

	query := `
			MATCH (recipe:Recipe {idRecipe: $id})
			MATCH (recipe)-[:CONTAINS]->(step:Step)
			MATCH (recipe)-[valueIngredient]->(ingredient:Ingredient)
			RETURN recipe, ingredient, step, properties(valueIngredient) as realIngredient
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

		recipe := createRecipe(data, []string{"recipe", "step", "ingredient", "realIngredient"})

		if err != nil {
			return nil, err
		}

		finalResult := types.APIResponse{
			Result: recipe,
		}

		return &finalResult, nil
	}

	return nil, err
}
