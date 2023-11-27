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
-   modify the list name of the specific idPlaylist

## Delete list

### Parameters : idplaylist : number

-   fetch the id and check that it belongs to the user email
-   delete the playlist with the id

## Display List

### Parameters : idplaylist : number

-   fetch the id and check that it belongs to the user email
-   modify the list name of the specific idPlaylist
