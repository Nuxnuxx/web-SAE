# Parameters : email of user for each

## Create List

### Parameters : namePlaylist : string

-   Create a list for the user with the specific email

MERGE (id:UniqueId {name: 'Playlist'})
ON CREATE SET id.count = 100
ON MATCH SET id.count = id.count + 1

WITH id.count AS uid
CREATE (p:Playlist {idPlaylist: uid, name: 'salutBg'})


WITH p, '0' AS userIdValue
MATCH (u:User {idUser: userIdValue})

// Créer un lien entre le nœud Playlist et le nœud User
CREATE (u)-[:A_UNE]->(p)

## Modify List

### Parameters : idplaylist : number , namePlaylist : string

-   fetch the id and check that it belongs to the user email
Match (p:Playlist) where p.idPlaylist = '19' 
match (u:User) where u.mail='Toni.Brown@gmail.com'
with p,u 
match (u)-[:A_UNE]->(p)
return u,p
-   modify the list name of the specific idPlaylist
   match (p:Playlist{idPlaylist:100}) SET p.name = 'nouveau nom'

## Delete list

### Parameters : idplaylist : number

-   fetch the id and check that it belongs to the user email
-   delete the playlist with the id
  match (p:Playlist{idPlaylist:100}) DELETE p

## Display List

### Parameters : idplaylist : number

MATCH (u:User{mail:'Amanda.Bernard@gmail.com'}) MATCH (p:Playlist) MATCH (u)-[:A_UNE]->(p) return u,p
-   fetch the id and check that it belongs to the user email
-   modify the list name of the specific idPlaylist
