package storage

import "backend/types"

type SuggestionStorage interface {
	// Suggestion
	GetMostLiked() (*types.APIResponse, error)
	GetTrending() (*types.APIResponse, error)
	GetSimilarRecipes(int, int) (*types.APIResponse, error)
	SetTimeSpentRecipe(string, int, int) error
}

func (s *Neo4jStore) GetMostLiked() (*types.APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n LIMIT 10"

	resp, err := s.db.Run(s.ctx, query, nil)

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

	finalResult := types.APIResponse{
		Result: recipeList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetTrending() (*types.APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n LIMIT 10"

	resp, err := s.db.Run(s.ctx, query, nil)

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

	finalResult := types.APIResponse{
		Result: recipeList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetSimilarRecipes(id int, number int) (*types.APIResponse, error) {
	query := `CALL gds.knn.stream('myGraph', {
		topK: $number,
		nodeProperties: ['difficultyNumeric','priceNumeric','preparationTimeNumeric','totalTimeNumeric'],
		randomSeed: 1337,
		concurrency: 1,
		sampleRate: 1.0,
		deltaThreshold: 0.0
	})
	YIELD node1, node2, similarity
	WITH gds.util.asNode(node1) AS recipe1, gds.util.asNode(node2) AS n, similarity
	WHERE recipe1.idRecipe = $id
	RETURN n
	ORDER BY similarity DESCENDING, recipe1, n`

	params := map[string]interface{}{
		"number": number,
		"id":     id,
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

	if err != nil {
		return nil, err
	}

	finalResult := types.APIResponse{
		Result: recipeList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) SetTimeSpentRecipe(mail string, idRecipe int, timeSpent int) error {
	/*query := `
		// Put the query here


	params := map[string]interface{}{
		"mail":      account.Mail,
		"idRecipe":  idRecipe,
		"timeSpent": timeSpent,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}
	*/
	return nil
}
