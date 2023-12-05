package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

func extractProperty(record neo4j.Record, key string, propertyName string) interface{} {
	return record.AsMap()[key].(neo4j.Node).Props[propertyName]
}

func CreateAccount(record neo4j.Record, key string) Account {
	name := strings.Split(extractProperty(record, key, "name").(string), " ")
	return Account{
		FirstName:         name[0],
		LastName:          name[1],
		Mail:              extractProperty(record, key, "mail").(string),
		EncryptedPassword: extractProperty(record, key, "password").(string),
	}
}

func CreatePagination(total int64, currentPage int) Pagination {
	return Pagination{
		CurrentPage: currentPage,
		TotalPage:   int(total / 10),
		TotalResult: int(total),
	}
}

func CreatePlaylistDetail(record neo4j.Record, key string) PlaylistDetail {
	return PlaylistDetail{
		Name:       extractProperty(record, key, "name").(string),
		IdPlaylist: extractProperty(record, key, "idPlaylist").(int64),
	}
}

func CreateRecipeDetail(record neo4j.Record, key string) RecipeDetail {
	return RecipeDetail{
		Difficulty: extractProperty(record, key, "difficulty").(string),
		Images:     extractProperty(record, key, "image").(string),
		Quantity:   extractProperty(record, key, "quantity").(string),
		Price:      extractProperty(record, key, "price").(string),
		Name:       extractProperty(record, key, "name").(string),
		IdRecipe:   extractProperty(record, key, "idRecipe").(int64),
	}
}

func CreateRecipeStep(record []*neo4j.Record, key string) RecipeStep {
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

func CreateRecipeIngredient(record []*neo4j.Record, key string) RecipeIngredients {
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
		RecipeDetail:      CreateRecipeDetail(*detailRecord, key[0]),
		RecipeStep:        CreateRecipeStep(record, key[1]),
		RecipeIngredients: CreateRecipeIngredient(record, key[2]),
	}
}

func buildQueryAndParams(query url.Values, page int) (string, map[string]interface{}) {
	queryString := "MATCH (n:Recipe) WHERE"
	values := url.Values{}

	for key := range query {
		if key == "name" {
			values.Add(key, fmt.Sprintf("n.%s CONTAINS $%s", key, key))
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
		"page":  page * 10,
		"limit": 10,
	}

	for key, value := range query {
		params[key] = value[0]
	}

	return queryString, params
}

func NewAccount(firstName, lastName, mail, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Mail:              mail,
		EncryptedPassword: string(encpw),
	}, nil
}
