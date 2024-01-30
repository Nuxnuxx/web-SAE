package storage

import (
	"backend/types"
	"fmt"
)

type ListStorage interface {
	// List
	CreateList(string, string) (*types.APIResponse, error)
	GetList(string) (*types.APIResponse, error)
	DeleteList(int) (*types.APIResponse, error)
	UpdateList(int, string) (*types.APIResponse, error)
	CheckListAlreadyExist(string, string) error
	CheckListBelongToUser(int, string) error

	// Recipe List
	CreateRecipeLiked(string, int) (*types.APIResponse, error)
	CreateRecipeList(string, int, int) (*types.APIResponse, error)
	DeleteRecipeList(int, int) (*types.APIResponse, error)
	GetRecipeList(int) (*types.APIResponse, error)
}

func (s *Neo4jStore) DeleteRecipeList(id int, idList int) (*types.APIResponse, error) {
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

	finalResult := types.APIResponse{
		Result: "The recipe has been deleted",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetRecipeList(id int) (*types.APIResponse, error) {
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
	recipes := make([]types.RecipeDetail, 0)
	for resp.Next(s.ctx) {
		recipe := CreateRecipeDetail(*resp.Record(), "r")
		recipes = append(recipes, recipe)
	}

	if err != nil {
		return nil, err
	}

	finalResult := types.APIResponse{
		Result: recipes,
	}

	return &finalResult, nil
}

func (s *Neo4jStore) CreateRecipeList(mail string, id, idList int) (*types.APIResponse, error) {
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

	finalResult := types.APIResponse{
		Result: "The recipe has been added",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) CreateRecipeLiked(mail string, idRecipe int) (*types.APIResponse, error) {
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

	finalResult := types.APIResponse{
		Result: "The recipe has been added",
	}

	return &finalResult, nil
}

func (s *Neo4jStore) GetList(mail string) (*types.APIResponse, error) {
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

	playlistList := make([]types.PlaylistDetail, 0)

	for resp.Next(s.ctx) {
		playlist := CreatePlaylistDetail(*resp.Record(), "p")
		playlistList = append(playlistList, playlist)
	}

	if err != nil {
		return nil, err
	}

	finalResult := types.APIResponse{
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

func (s *Neo4jStore) UpdateList(id int, name string) (*types.APIResponse, error) {
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

	result := types.APIResponse{
		Result: "Playlist has been updated",
	}
	return &result, nil
}

func (s *Neo4jStore) DeleteList(id int) (*types.APIResponse, error) {
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

	result := types.APIResponse{
		Result: "Playlist has been deleted",
	}
	return &result, nil
}

func (s *Neo4jStore) CreateList(name string, mail string) (*types.APIResponse, error) {
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

	result := types.APIResponse{
		Result: "Playlist has been created",
	}
	return &result, nil
}
