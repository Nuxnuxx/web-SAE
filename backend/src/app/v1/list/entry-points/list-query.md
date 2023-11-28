# Parameters : email of user for each
### POUR CHNAGER EN INT :
- MATCH (n:Ingredient) FOREACH (node IN [1] | SET n.idIngredient = toInteger(n.idIngredient))
- MATCH (n:Recipe) FOREACH (node IN [1] | SET n.idRecipe = toInteger(n.idRecipe))
- MATCH (n:Playlist) FOREACH (node IN [1] | SET n.idPlaylist = toInteger(n.idPlaylist))
- MATCH (n:User) FOREACH (node IN [1] | SET n.idUser= toInteger(n.idUser))
## Create List

### Parameters : namelist : string

-   Create a list for the user with the specific email
MERGE (id:UniqueId {name: 'Playlist'}) ON CREATE SET id.count = 100 ON MATCH SET id.count = id.count + 1
WITH id.count AS uid CREATE (p:Playlist {idPlaylist: uid, name: 'salutBg'})
WITH p, '0' AS userIdValue MATCH (u:User {idUser: userIdValue})
CREATE (u)-[:A_UNE]->(p)

## Modify List

### Parameters : idplaylist : number , namelist : string

-   fetch the id and check that it belongs to the user email
Match (p:Playlist) where p.idPlaylist = '19' match (u:User) where u.mail='Toni.Brown@gmail.com' with p,u match (u)-[:A_UNE]->(p) return u,p
-   modify the list name of the specific idlist
match (p:Playlist{idPlaylist:100}) SET p.name = 'nouveau nom'
## Delete list FIXED

### Parameters : idplaylist : number

-   fetch the id and check that it belongs to the user email
-   idem
-   delete the playlist with the id
-   MATCH (p:Playlist{idPlaylist:'1'}) MATCH ()-[l:A_UNE]->(p) MATCH ()-[l2:est_dans]->(p) delete l,l2,p

## Display List

### Parameters : idplaylist : number

-   fetch the id and check that it belongs to the user email
IDEM
-   modify the list name of the specific idlist J'imagine que tu veux que je display les playlist de l'utilisateur
MATCH (u:User{mail:'Amanda.Bernard@gmail.com'}) MATCH (p:Playlist) MATCH (u)-[:A_UNE]->(p) return u,p
## Create recipe in liked

### Paramaters : idRecipe : number FIXME WORKING
- Pour créer la playlist liked :
-  match (u:User) where u.mail='Toni.Brown@gmail.com' create (p:Playlist{name:'liked'}) with p,u create (u)-[:A_UNE]->(p) 

- add the recipe to liked list
  Match (p:Playlist{name:'liked'})  match (u:User) where u.mail='Toni.Brown@gmail.com' MATch (u)-[:A_UNE]->(p)   Match (r:Recipe{idRecipe:'16'})with p,u,r CREATE (r)-[:est_dans]->(p) 

## Add recipe in a list

### Paramaters : idRecipe : number FIXME

- add the recipe to list
- Match (p:Playlist{idPlaylist:'18'})  match (u:User) where u.mail='Toni.Brown@gmail.com' MATch (u)-[:A_UNE]->(p)   Match (r:Recipe{idRecipe:'16'})with p,u,r CREATE (r)-[:est_dans]->(p) 

## Get all Recipes in a list

### Parameters : idList : number FIXME

- get all recipe from a playlist
  MATCH (r:Recipe)-[:est_dans]->(p:Playlist{idPlaylist:'0'}) return r

## Delete one recipe from a list FIXME

### Parameters : idrecipe: number , idlist : number

- delete one recipe from a list
MATCH (r:Recipe{idRecipe:'286'})-[l:est_dans]->(p:Playlist{idPlaylist:'0'}) delete l
