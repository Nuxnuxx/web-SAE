package storage

import (
	"fmt"
	"net/url"
	"strings"

	"backend/types"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func extractProperty(record neo4j.Record, key string, propertyName string) interface{} {
	value, ok := record.AsMap()[key]
	if !ok {
		return nil
	}

	if node, ok := value.(neo4j.Node); ok {
		return node.Props[propertyName]
	} else {
		return value.(map[string]interface{})[propertyName]
	}
}

func CreateAccount(record neo4j.Record, key string) types.Account {
	name := strings.Split(extractProperty(record, key, "name").(string), " ")
	return types.Account{
		FirstName:         name[0],
		LastName:          name[1],
		Mail:              extractProperty(record, key, "mail").(string),
		EncryptedPassword: extractProperty(record, key, "password").(string),
	}
}

func CreatePagination(total int64, currentPage int) types.Pagination {
	return types.Pagination{
		CurrentPage: currentPage,
		TotalPage:   int(total / 9),
		TotalResult: int(total),
	}
}

func CreatePlaylistDetail(record neo4j.Record, key string) types.PlaylistDetail {
	var image string

	if prop := extractProperty(record, key, "image"); prop != nil {
		if imageVal, ok := prop.(string); ok {
			image = imageVal
		}
	}

	return types.PlaylistDetail{
		Name:          extractProperty(record, key, "name").(string),
		Image:         image,
		IdPlaylist:    extractProperty(record, key, "idPlaylist").(int64),
		NumberRecipes: extractProperty(record, key, "numberOfRecipes").(int64),
	}
}

func CreateRecipeDetail(record neo4j.Record, key string) types.RecipeDetail {
	return types.RecipeDetail{
		Difficulty:      extractProperty(record, key, "difficulty").(string),
		Images:          extractProperty(record, key, "image").(string),
		Quantity:        extractProperty(record, key, "quantity").(string),
		Price:           extractProperty(record, key, "price").(string),
		PreparationTime: extractProperty(record, key, "preparationTime").(string),
		Name:            extractProperty(record, key, "name").(string),
		IdRecipe:        extractProperty(record, key, "idRecipe").(int64),
	}
}

func CreateRecipeStep(record []*neo4j.Record, key string) types.RecipeStep {
	recipe := make(types.RecipeStep)
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

func CreateRecipeIngredient(record []*neo4j.Record, key string, key2 string) types.RecipeIngredients {
	recipe := make(types.RecipeIngredients)

	for _, r := range record {
		recipeId := extractProperty(*r, key, "idIngredient").(int64)

		recipe[recipeId] = types.IngredientInfo{
			Name:       extractProperty(*r, key2, "libelleIngredient").(string),
			URLPicture: extractProperty(*r, key, "urlPicture").(string),
		}
	}

	return recipe
}

func createRecipe(record []*neo4j.Record, key []string) types.Recipe {
	detailRecord := record[0]
	return types.Recipe{
		RecipeDetail:      CreateRecipeDetail(*detailRecord, key[0]),
		RecipeStep:        CreateRecipeStep(record, key[1]),
		RecipeIngredients: CreateRecipeIngredient(record, key[2], key[3]),
	}
}

func buildQueryAndParams(query url.Values, page int) (string, map[string]interface{}) {
	queryString := "MATCH (n:Recipe) WHERE"
	values := url.Values{}

	for key := range query {
		if key == "name" {
			values.Add(key, fmt.Sprintf("n.%s =~ '(?i).*'+$%s+'.*'", key, key))
		} else {
			values.Add(key, fmt.Sprintf("n.%s = $%s", key, key))
		}
	}

	i := 0
	for _, value := range values {

		if i == len(values)-1 {
			queryString += " " + value[0]
			break
		}
		queryString += " " + value[0] + " AND"
		i++
	}

	queryString += " RETURN n SKIP $page LIMIT $limit"

	params := map[string]interface{}{
		"page":  page * 9,
		"limit": 9,
	}

	for key, value := range query {
		params[key] = value[0]
	}

	return queryString, params
}
