package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"golang.org/x/crypto/bcrypt"
)

type Storage interface {
	// Suggestion
	GetMostLiked() (*APIResponse, error)
	GetTrending() (*APIResponse, error)

	// Recipe List
	CreateRecipeLiked(string, int) (*APIResponse, error)
	CreateRecipeList(string, int, int) (*APIResponse, error)
	DeleteRecipeList(int, int) (*APIResponse, error)
	GetRecipeList(int) (*APIResponse, error)

	// List
	CreateList(string, string) (*APIResponse, error)
	GetList(string) (*APIResponse, error)
	DeleteList(int) (*APIResponse, error)
	UpdateList(int, string) (*APIResponse, error)
	CheckListAlreadyExist(string, string) error
	CheckListBelongToUser(int, string) error

	// Account
	CreateAccount(*Account) error
	GetAccountByMail(string) (*Account, error)
	DeleteAccount(string) error
	UpdateAccount(*Account) error
	FindAccountByMail(string) error
	Login(LoginRequest) (*Account, error)

	// Recipe
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
	uri := os.Getenv("DB_URL")
	auth := neo4j.NoAuth()

	if Stage == "dev" {
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

func (s *Neo4jStore) GetTrending() (*APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n LIMIT 10"

	resp, err := s.db.Run(s.ctx, query, nil)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	recipeList := make([]RecipeDetail, 0)

	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "n")

		recipeList = append(recipeList, recipe)
	}

	finalResult := APIResponse{
		Result: recipeList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetMostLiked() (*APIResponse, error) {
	query := "MATCH (n:Recipe) RETURN n LIMIT 10"

	resp, err := s.db.Run(s.ctx, query, nil)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	recipeList := make([]RecipeDetail, 0)

	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "n")

		recipeList = append(recipeList, recipe)
	}

	finalResult := APIResponse{
		Result: recipeList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) DeleteRecipeList(id int, idList int) (*APIResponse, error) {
	query := "MATCH (r:Recipe{idRecipe:$idRecipe})-[l:est_dans]->(p:Playlist{idPlaylist:$idList}) delete l"

	params := map[string]interface{}{
		"idRecipe": id,
		"idList":   idList,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	finalResult := APIResponse{
		Result: "The recipe has been deleted",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetRecipeList(id int) (*APIResponse, error) {
	query := "MATCH (r:Recipe)-[:est_dans]->(p:Playlist{idPlaylist:$id}) return r"

	params := map[string]interface{}{
		"id": id,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	// create array of recipeDetail
	recipes := make([]RecipeDetail, 0)
	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "r")
		recipes = append(recipes, recipe)
	}

	if err != nil {
		return nil, err
	}

	finalResult := APIResponse{
		Result: recipes,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) CreateRecipeList(mail string, id, idList int) (*APIResponse, error) {
	query := `
		MATCH (p:Playlist{idPlaylist:$idList})
		MATCH (u:User) WHERE u.mail = $mail
		MATCH (u)-[:A_UNE]->(p)
		MATCH (r:Recipe{idRecipe: $idRecipe})
		CREATE (r)-[:est_dans]->(p)
	`

	params := map[string]interface{}{
		"idList":   idList,
		"mail":     mail,
		"idRecipe": id,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	finalResult := APIResponse{
		Result: "The recipe has been added",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) CreateRecipeLiked(mail string, idRecipe int) (*APIResponse, error) {
	query := `
		MATCH (p:Playlist{name:'liked'})
		MATCH (u:User) WHERE u.mail = $mail
		MATCH (u)-[:A_UNE]->(p)
		MATCH (r:Recipe{idRecipe: $id})
		CREATE (r)-[:est_dans]->(p)
	`

	params := map[string]interface{}{
		"mail": mail,
		"id":   idRecipe,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	finalResult := APIResponse{
		Result: "The recipe has been added",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetList(mail string) (*APIResponse, error) {
	query := `
		MATCH (u:User {mail: $mail})-[:A_UNE]->(p:Playlist)
		OPTIONAL MATCH (r:Recipe)-[:est_dans]->(p)
		WITH p, COUNT(r) AS numberOfRecipes, COLLECT(r)[0] AS firstRecipe
		SET p.numberOfRecipes = numberOfRecipes
		SET p.image = toString(firstRecipe.image)
		RETURN p
	`

	params := map[string]interface{}{
		"mail": mail,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	playlistList := make([]PlaylistDetail, 0)

	for resp.Next(s.ctx) {
		playlist := CreatePlaylistDetail(*resp.Record(), "p")
		playlistList = append(playlistList, playlist)
	}

	if err != nil {
		return nil, err
	}

	finalResult := APIResponse{
		Result: playlistList,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) CheckListBelongToUser(id int, mail string) error {
	query := "MATCH (p:Playlist) where p.idPlaylist = $id match (u:User) where u.mail=$mail with p,u match (u)-[:A_UNE]->(p) return u,p"

	params := map[string]interface{}{
		"mail": mail,
		"id":   id,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	if resp.Next(s.ctx) {
		return nil
	}

	return fmt.Errorf("Playlist doesnt belong to the user")
}

func (s *Neo4jStore) CheckListAlreadyExist(name, mail string) error {
	query := "MATCH (:User {mail: $mail})-[:A_UNE]->(p:Playlist {name: $name}) RETURN p"

	resp, err := s.db.Run(s.ctx, query, map[string]interface{}{
		"name": name,
		"mail": mail,
	})

	if err != nil {

		return err
	}

	if resp.Err() != nil {
		return err
	}

	if resp.Next(s.ctx) {
		return fmt.Errorf("Playlist found")
	}

	return nil

}

func (s *Neo4jStore) UpdateList(id int, name string) (*APIResponse, error) {
	query := "match (p:Playlist{idPlaylist:$id}) SET p.name = $name"

	params := map[string]interface{}{
		"id":   id,
		"name": name,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	result := APIResponse{
		Result: "Playlist has been updated",
	}
	return &result, nil
}

func (s *Neo4jStore) DeleteList(id int) (*APIResponse, error) {
	queryString := `
		MATCH (p:Playlist{idPlaylist:$id}) MATCH ()-[l:A_UNE]->(p) OPTIONAL MATCH ()-[l2:est_dans]->(p) delete l,l2,p
	`

	params := map[string]interface{}{
		"id": id,
	}

	resp, err := s.db.Run(s.ctx, queryString, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	result := APIResponse{
		Result: "Playlist has been deleted",
	}
	return &result, nil
}

func (s *Neo4jStore) CreateList(name string, mail string) (*APIResponse, error) {
	queryString := `
		MERGE(id: UniqueId { name: 'Playlist' })
		ON CREATE SET id.count = 100
		ON MATCH SET id.count = id.count + 1

		WITH id.count AS uid
			CREATE(p: Playlist { idPlaylist: uid, name: $name })


		WITH p, $mail AS userMailValue
			MATCH(u: User { mail: userMailValue})

		// Créer un lien entre le nœud Playlist et le nœud User
		CREATE(u) - [: A_UNE] -> (p)
	`

	params := map[string]interface{}{
		"name": name,
		"mail": mail,
	}

	resp, err := s.db.Run(s.ctx, queryString, params)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	result := APIResponse{
		Result: "Playlist has been created",
	}
	return &result, nil
}

//INFO: If it return a error then the email has been found
// 			If response is nil then the account is not found
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

func (s *Neo4jStore) GetAccountByMail(mail string) (*Account, error) {
	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) return u;",
		map[string]interface{}{
			"mail": mail,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		account := CreateAccount(*resp.Record(), "u")

		return &account, nil
	}

	return nil, nil
}

func (s *Neo4jStore) DeleteAccount(mail string) error {
	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) DETACH DELETE u;",
		map[string]interface{}{
			"mail": mail,
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

func (s *Neo4jStore) UpdateAccount(acc *Account) error {

	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) SET u.password = $password RETURN u",
		map[string]interface{}{
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

func (s *Neo4jStore) CreateAccount(acc *Account) error {

	resp, err := s.db.Run(s.ctx, "CREATE (u:User {name: $name, mail: $mail, password: $password, gender: $gender}) RETURN u",
		map[string]interface{}{
			"name":     acc.FirstName + " " + acc.LastName,
			"mail":     acc.Mail,
			"gender":   acc.Gender,
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
