# Cahier des charges

## Origine du projet


### Situation de départ

- Il existe une ribambelle de services de recettes culinaires qui ne sont pas assez orientés sur l'expérience utilisateur. Lorsque l'on va sur ce type de site, nous avons déjà une idée de la recette que nous allons faire. Au final le côté "Quelle recette vais-je faire ?" a été fait en amont et c'est le moteur de recherche qui nous redirige sur le site. Nous pouvons faire une comparaison avec une application de musique qui permet d'en écouter exclusivement grâce à un titre rentré par l'utilisateur. Ce genre d'application n'est pas centrée sur l'utilisateur mais sur les recettes en elles-mêmes. La problématique est de savoir comment proposer un site qui gère tout le chemin de la recette à l'utilisateur en se recentrant sur ses goûts. 

### Développement de la problématique

- Notre objectif est de proposer une expérience personnalisée. Par exemple, sur les applications de musique, on nous permet de découvrir de nouveaux groupes selon les styles que nous aimons. Dans notre cas, nous pourrons proposer de nouveaux plats selon ce que l'utilisateur a fait et a aimé. Nos goûts en musique évoluent et s'élargissent alors pourquoi ne pas avoir la même chose avec la cuisine ? Combien de fois s'est-on demandé "Qu'est ce que j'achète et qu'est-ce que je mange cette semaine ?" pour au final manger des plats habituels ? Combien de personnes sont bloquées sur les mêmes plats en boucle comme une playlist de 8 morceaux ? Notre objectif n'est pas seulement de faire découvrir à l'utilisateur de nouvelles manières de cuisiner et de découvrir de nouveaux plats, mais aussi de s'assurer qu'il apprécie ces nouvelles choses grâce à une expérience personnalisée. 


![Diagramme Bete a corne](bete_a_corne.png)

### Comment répondre à cette problématique ?

#### Tout d'abord nous devons proposer un site de recettes fonctionnel :

- Création d'un site web, qui comme sur un site web de recette lambda, permet de rechercher des recettes et les filtrer. 

- Ajout d'une interface utilisateur simple et intuitive coordonnée avec une expérience utilisateur optimisée pour fidéliser les visiteurs du site. Ajout d'un principe d'authentification qui permet de récupérer ses sauvegardes sur n'importe quel appareil.

#### Ensuite nous devons y ajouter le côté personnalisation :

L'objectif est donc de proposer un site qui gère tout le chemin de la recette à l'utilisateur en se recentrant sur ses goûts via une expérience personnalisée. Nous intégrerons des "livres de recettes" où les gens peuvent créer des listes et y ajouter des recettes. De plus nous ajouterons des suggestions personnalisées tout au long du site établies en fonction des habitudes des utilisateurs ainsi que de leur profil individuel mais aussi de la similitude des recettes. Tous les détails de personnalisation sont inscrits dans la partie "Fonctionnalités".



## Fonctionnalités

![Diagramme Pieuvre](diagramme_pieuvre.png)

### Personnalisation

#### Sauvegarde de(s) recette(s)

- Pour sauvegarder les recettes dans des playlists (like ou autre playlists), il sera obligatoire de se connecter, et de s'inscrire au préalable. Grâce à son nouveau compte, l'utilisateur aura plusiseurs possibilités : sauvegarder la recette dans sa playlist "Like" (playlist par défaut non supprimable) et/ou dans une liste créée spécialement par l'utilisateur avec un nom personnalisé. L'utilisateur pourra créer autant de playlists qu'il le souhaitera. 

#### Suggestion

##### Recommandations globales

- L'idée de la création de recommandations globales semble idéale pour donner un avant-goût des fonctionnalités aux nouveaux utilisateurs lors de leur première session. De plus, cela nous permettra de présenter les tendances globales du site. En effet, n'ayant pas de profil, il n'y a pas de données sur l'utilisateur. Des recommandations globales basées sur la popularité des produits et sur les tendances du moment permettent de donner envie à l'utilisateur de s'inscrire pour avoir accès à la recommandation personnalisée.

Pour la recommandation globale, nous allons nous concentrer sur deux modules précis qui serviront de base pour guider tous les utilisateurs (connectés ou non).

###### Trending (Tendances)

Le système de trending fonctionnera sur le même principe la technique sera différente. Le Trending, ou tendances, pourra afficher à l'utilisateur des recettes qui sont populaires, qui suscitent un intérêt pour les utilisateurs. Grâce au Trending, il sera possible de donner des recettes dites "du moment". Avec cette approche, il sera possible de susciter l'engagement des utilisateurs plus facilement dès leur première session : en plus des recommandations gloables, ils auront accès aux tendances. Cela les incitera à explorer davantage le site web.

Il faut retenir que le module de "Trending" enregistre les recettes les plus vues et "likées" sur les 30 derniers jours, contrairement au module suivant "Most Liked". L'affichage du "Trending" est donc voué à être fluide sur l'affichage de différents contenus sur le long terme.

###### Most Liked (Les plus aimées)

Le module "Most Liked" suggèrera des produits, des contenus ou des éléments en se basant sur le nombre de "likes" reçus par les divers utilisateurs. Avec cette approche, la liste des recettes appréciées par un grand nombre d'utilisateurs s'affichera et sera susceptible d'attirer l'attention de nouveaux utilisateurs et de les fidéliser.

Il faut retenir que le module "Most Liked" enregistre les recettes, les plus "likées" depuis le début du site, contrairement au module de "Trending" qui réalise les tendances sur les 30 derniers jours.

##### Recommandation contextuelle

Le module de Recommandation contextuelle va venir personnaliser les suggestions à un utilisateurs selon un large éventail de contextes et d'attribut liés aux recettes et à l'utilisateur. Dans le cadre de ce projet, nous nous concentrerons sur les attribut liés aux recettes (vu dans la partie suivante : "Item by item") ainsi que sur les utilisateurs en leur permettant d'insérer, lors de la création de leurs compte, leurs préférences alimentaires, un suivi de leurs habitude de recherche,... lors de l'utilisation du site (une plus grande explication est disponible au point "Dry Start" plus loin).

###### Item by Item (Personnalisé par caractéristique)

Le module "Item by item", ou personnalisation par caractéristiques, est une stratégie de recommandation qui se concentre sur la personnalisation des suggestions. Il se base sur les caractéristiques spécifiques de chaque recette. Une approche telle que celle-ci va permettre d'interpréter le contenu des recettes pour en sortir des relations. Ces relations vont permettre de proposer des recettes en rapport avec celles recherchées par l'utilisateur. De plus nous pourrons proposer de nouvelles recettes à ajouter dans un livre de recettes en fonction de ce qu'il y a déjà.

###### Dry Start (Démarrage à vide)

Le module de "Dry Start" sera très important pour réaliser le système de suggestions pour les nouveaux utilisateurs. En effet, un formulaire lors de l'inscription permettra de définir dès lors les préférences de l'utilisateur. Grâce à cette fonctionnalité, malgré le fait que l'utilisateur soit nouveau, il sera possible de lui présenter des suggestions plus affinées que celles globales.

##### Recommandations personnalisées

Le module de recommandation personnalisée sera un aspect clé de l'expérience utilisateur. Les recommandations vont tenir compte des préférences et des comportements de l'utilisateur pour offrir des suggestions sur mesure. Une recommandation personnalisée se base sur le suivi en temps réel de leurs habitudes de recherche et de navigation sur le site. Avec cette approche, il est garanti d'avoir des recommandations plus pertinentes et donc une expérience utilisateur améliorée.

Dans le cadre de notre projet, nous nous reposerons sur l'analyse des actions et des choix antérieurs de l'utilisateur pour lui proposer des recettes qui corresponderont à ses préférences. L'algorithme veillera à ce que chaque utilisateur bénéficie de suggestions adaptées à son profil, améliorant ainsi leur satisfaction et leur engagement sur le site.

### Expérience utilisateur

#### Interface

L'interface graphique va être le point clé pour l'utilisateur : c'est ce qu'il verra en premier sur notre site. Pour garantir une expérience optimale, nous concentrerons nos efforts sur une interface intuitive et attrayante. La simplicité d'utilisation sera cruciale : tout type d'utilisateur doit pouvoir naviguer facilement, sans accroc ni confusion. 

L'interface devra être : propre, structurée et esthétique pour attirer davantage l'attention des utilisateurs en plus de renforcer leur confiance et augmenter le temps passé sur le site. Les éléments visuels, tels que des icônes intuitives, une mise en page cohérente et des visuels attrayants, contribuent à rendre l'interface conviviale (voir le point Maquette pour un visuel).

#### Navigation

À travers la navigation, les utilisateurs devront pouvoir trouver rapidement ce qu'ils recherchent. Avec une navigation intuitive et bien organisée, des menus clairs et des liens pertinents, l'expérience de l'utilisateur sera simplifiée et donc beaucoup plus attrayante. L'utilisation d'un moteur de recherche avec des filtres pertinents permettra de satisfaire l'utilisateur et l'augmenter sur le long terme. Une navigation fluide incitera davantage les utilisateurs à rester sur le site, puis l'explorer.

Pour les utilisateurs non connectés, cela permettra de créer une envie d'avoir accès aux fonctionnalités exclusives d'un compte telle que la pertinence des suggestions.

#### Multiplateforme

Pour pouvoir toucher un large public, il est essentiel d'avoir une grande accessibilité. Nous devrons nous assurer que l'expérience utilisateur est cohérente sur toutes les plateformes (ordinateurs, tablettes, mobiles...). Cela permettra à l'utilisateur d'accéder à toutes ses recettes avec son compte via n'importe quelle plateforme à partir du moment où il a accès à internet.

## La cible

### L’étudiant
Paul Biard, 20 ans, est un étudiant vivant à Amiens. Il est en études supérieures pour la première fois et vit donc seul. Il étudie en informatique et est très sédentaire. Il aime jouer avec ses amis en ligne et commande beaucoup de nourriture via des applications telles qu'Uber Eats. La nourriture est très peu équilibrée, grasse et ne lui confère aucune énergie. Paul souhaite donc utiliser PirateCook pour se mettre à cuisiner. Il peut tester des recettes et en découvrir de nouvelles avec le système de suggestion. Il pourra créer des playlists de recettes selon ses envies. Les sites tels que Marmiton ou Jow ne lui conviennent pas : il n’est pas possible de garder les recettes facilement pour l’un et l’application a des difficultés sur téléphone pour l’autre. Paul est frustré de perdre autant d’argent dans l’utilisation d'applications de livraison. En effet, les prix sont beaucoup plus élevés. Il veut donc pouvoir cuisiner facilement et dépenser peu. Paul souhaite pouvoir utiliser le site facilement sur son téléphone étant donné que sa cuisine est trop petite pour poser son ordinateur. 

### Le retraité
Jean-Bernard Delacroix est un retraité de 69 ans. Il aime beaucoup faire de la pêche toute la journée avec ses petits-enfants. Malheureusement, Jean-Bernard ne cuisine pas ses prises car il ne sait pas comment faire. Il se retrouve donc avec ses poissons sans pouvoir les cuisiner. Ses enfants lui conseillent d’utiliser son ordinateur et d’aller sur PirateCook pour apprendre des recettes. De plus, grâce aux suggestions, il pourra en découvrir de nouvelles. C’est difficile pour Jean-Bernard d’utiliser un ordinateur mais il a la motivation d’apprendre grâce à la facilité d’utilisation de PirateCook. Grâce à PirateCook, il pourra réaliser des repas avec ses poissons à ses petits-enfants après la pêche de la journée.

###  Salarié
José Bertier, 42 ans, travaille en tant que salarié. Le soir, c’est sa femme qui cuisine pour eux deux. Il souhaite donc cuisiner pour la décharger de ce travail. Il veut cuisiner sans l’aide de sa femme. Cependant, il ne veut pas refaire encore et encore les mêmes repas. Marmiton n’offre pas une suggestion assez poussée et il souhaite enregistrer les repas selon des catégories précises. Avec PirateCook il peut facilement apprendre grâce aux étapes de la recette et garder celles qu’il aime. Il souhaite l’utiliser sur téléphone la journée lors de ses pauses au travail et cuisiner avec son ordinateur.

### Mère de famille
Clothilde Marie est une mère de 5 enfants. Elle n’a donc pas beaucoup de temps pour elle. Il lui est long de réaliser des repas. Elle utilise Marmiton mais elle prépare des recettes en boucle et a du mal à les retrouver. Il n’est pas rare qu’elle utilise aussi d’autres sites moins connus de cuisine. Elle souhaite donc trouver un site qui lui permet de voir des recettes peu chères et faciles à faire pour les noter et y avoir accès à tout moment sur ordinateur et téléphone. Grâce à PirateCook, elle peut retrouver ses recettes facilement et gagner du temps grâce à des recettes faciles à réaliser pour elle et ses 5 enfants.

### Cadre
Jérémie Albion est un cadre supérieur dans une entreprise de tech. Avec son travail, il n’a pas le temps de cuisiner par lui-même : il paie un service de livraison de repas équilibré. Cependant, il a rencontré une femme et souhaite l’impressionner lors d’un dîner en cuisinant un repas complexe. Sur Marmiton, il ne trouve pas de recette intéressante et se tourne donc vers PirateCook pour trouver ce qu’il souhaite : une recette complexe à n’importe quel prix. Grâce à PirateCook, il peut “liker” la recette et la retrouver facilement sur toutes les plateformes.

## La concurrence

![Tableau des concurrence ](concurrence.png)

## Aspects techniques

### Infrastructure

#### Docker

Dans l'optique d'améliorer la praticité du développement et de l'utilisation du site web et de l'application, nous utiliserons Docker. Les caractéristiques propres aux conteneurs sont avantageuses car elles permettent de faciliter l'assemblage, la maintenance, et la portabilité, grâce à l'isolation des applications entre elles et du système sous-jacent. Par rapport à une machine virtuelle, les conteneurs utilisent plus efficacement les ressources à leur disposition et en nécessitent moins, parce qu'ils partagent le même noyau que le système hôte au lieu de démarrer leur propre système d'exploitation. Conjointement, le démarrage est plus rapide. Enfin, Docker propose des outils qui facilitent la gestion des versions pour mettre à jour facilement ainsi que des options de sauvegarde des conteneurs et des volumes simplifiées.

#### Serveur Web

Afin d'héberger le site web, nous utiliserons le serveur web NGINX. Puisque nous utiliserons Node.js, il semble judicieux de l'utiliser à la place d'Apache pour plusieurs raisons : Nginx est réputé pour son efficacité en matière de gestion de connexions simultanées, ce qui le rend adapté aux applications en temps réel comme Node.js. La configuration de Nginx est généralement plus légère que celle d'Apache, signifiant qu'il nécessitera moins de ressources. Nginx est également excellent pour traiter les requêtes statiques, ce qui réduit aussi la charge serveur. 

#### Base de données



### Architecture

- Le choix du modèle MVC (Modèle-Vue-Contrôleur) pour ce projet se justifie sur plusieurs points importants : il sera possible de séparer le travail entre tous les développeurs facilement, avec une maintenabilité et une réutilisation du code facilitée, une bonne scalabilité (capacité à être mis à l'échelle) et enfin une conformité aux bonnes pratiques avec l'utilisation de frameworks (boîte à outils pour un développeur web).

Cette approche architecturale permet une organisation claire des composantes du projet : la gestion des données (Modèle), l'interface utilisateur (Vue) et la logique de l'application (Contrôleur). De plus, étant donné que nous souhaitons apporter des mises à jour, telles que des ajouts futurs de nouvelles fonctionnalités, il est nécessaire d'avoir une architecture modulable. Le modèle MVC est donc un choix solide pour garantir un développement structuré et évolutif de ce site internet complexe.

#### View

- Pour le choix de la Vue, nous avons opté pour le framework front end "**Svelte**" (https://svelte.dev/). Ce framework a été choisi pour sa distinction par rapport aux autres frameworks. La création d'une web app (SPA) se justifie par l'ajout de réactivité dans le site web et une navigation plus fluide. La facilité de réutilisation des composants permettra un dévelopemment plus organisé et plus rapide avec moins de répétitions et une maintenance plus aisée. La gestion des états de l'application nous permettra de gérer plus facilement les métriques utilisateur pour personnaliser leur expérience. Cela contribuera aussi à la compatibilité interplateforme et à la simplification d'intégration de l'application sur l'App Store ou le Play Store. Svelte possède une grosse communauté qui permet l'ajout de bibliothèques aidant au dévelopemment, de modules de transition et d'animation integrés à Svelte qui, par défaut, permettront de produire une interface attrayante.

- Le backend sera en Nodejs - Express, donc en javascript. L'utilisation de ce même langage en front permet l'homogénéité de l'ensemble du code source.

#### Controller

- Le choix du backend sera Express dans l'écosystème NodeJS. Ce choix est justifié par sa réputation de framework solide, stable et fiable pour le développement web. Cela facilitera l'implémentation de plusieurs de nos modules : la gestion des utilisateurs, des recettes, la sécurité des données,... Grâce à ses fonctionnalités de middleware (passerelle entre plusieurs applications, outils et bases de données), Express simplifie la gestion de l'authentification, de l'autorisation et de la validation des données, garantissant ainsi la qualité des données et la sécurité de l'application. Pour notre projet visant un UX performant, Express est optimisé pour les performances et une navigation du site plus fluide.

#### Model

- Dans le cadre d'une recette de cuisine, les informations auront besoin en priorité d'être cohérentes et relationnelles. En effet, des informations telles que la recette, les ingrédients ou encore les tags pourront être étudiés pour réaliser les suggestions. Dans cet idée, l'utilisation d'une base de donnée relationnelle est avisée. La technologie utilisée sera MySQL pour sa facilité d'utilisation. En effet, la simplicité de cette technologie permettra de garder l'essentiel et de faire fonctionner notre base de données d'une manière propre et efficace.

![Diagramme base de données](diagrammeBaseDeDonnee.png)

- Dans le cadre des recommendations nous devons représenter les relations entre les différentes recettes (item-per-item). Pour cela nous utiliserons neoforJ qui permet de faire une base de donnée orienté graph et de faire des liens entre entitées (dans notre cas des recettes). NeoForJ permet d'explorer rapidement les données, et est optimisé pour modifier les poids rapidements.

### Analyse Fonctionnelle des besoins

![Diagramme Priorite des besoins](priorite_besoin.png)

#### Module : Apercu recette

##### Affichage d'une recette

- Front
  - components : recipe/:id
    - Composant affichant une recette
- Back
  - /recipe/id/:id
    - Récupère les informations d'une recette
- Model
  - table recette
    - entites
      - idRecette
      - nomRecette
      - urlImage
      - nbrLike

##### Affichage d'une recette pour carousel

- Front
    - components : recipeCard
    - composant affichant une recette dans une card
    - Paramètre : Recipe (données d'une recette (ex: nomRecette/imageRecette...))
- Back
    - /recipe?id=
- Model
  - table recette
    - entites
      - idRecette
      - nomRecette
      - urlImage
      - nbrLike
      - user a liké ou pas
      - user a mis dans liste ou pas


##### Affichage des recettes

- Front
  - components : recipes
    - Composant affichant la liste des recettes avec une pagination
    - Paramètre : listRecipe (liste de recette transmis par la searchBar et/ou les filtres)
- Back
  - /recipe/page/:page
    - Récupère un nombre X de recettes (pour pagination exemple : 13)
- Model
  - table recette
    - entites
      - idRecette
      - nomRecette
      - urlImage
      - nbrLike
  - table user && recette
    - entites
      - ? ajouter a une liste ou pas
      - ? ajouter au liker ou pas

#### Module : Recherche recette

##### Recherche selon mots-clés :

- Front
  - Components : searchBar
  - Composant avec un input permettant aux utilisateurs de saisir des mots-clés pour rechercher des recettes
- Back
  - /search/:keyword
  - Reçoit les mots-clés saisis par l'utilisateur et demande à la base de données de retourner les résultats correspondants
- Model
- Table recette
  - Entités
    - nomRecette
    - description

##### Recherche filtré :

- Front
  - Components : filter
  - Composant permettant aux utilisateurs de filtrer leurs résultats de recherche par divers critères (ex. : durée de préparation, ingrédients, etc.)
- Back
  - /filter?critere=value?critere=value?
  - Reçoit les critères de filtrage choisis par l'utilisateur et demande à la base de données de retourner les résultats filtrés
- Model
- Table recette
  - Entités
      - difficulté
      - temps
      - prix
      - nbrLike
- Table Ingrédient
  - Entités
    - nomIngrédient
- Table tags
    - tags

##### Fonctionnalités de recherche et de filtrage

Les critères de recherche et de filtrage peuvent inclure des paramètres tels que :

- Tags
- Temps total
- ? Ingrédients
- Niveau de difficulté
- Évaluation
- etc.

#### Module : Authentification

##### Création du compte

- Front
  - components : Register
  - Composant avec formulaire d'inscription
- Back
  - /register
  - creer l'utilisateur sans ses preferences
- Model
- table user
  - entites
    - id
    - email
    - nom
    - prenom
    - sexe
    - password


##### Modification du compte

- Front
  - composant : modifyProfil
  - composant affichant le compte ainsi que des input pour y modifier ses information lors de clique du boutton
  - Contexte : User
- Back
  - /profil method PUT
- Model
    - table user
      - addresseMail
      - nom
      - prenom
      - mdpCrypte
      - sexe

##### Suppression du compte

- Front
    - composant : Profil
    - bouton pour supprimer son profil
    - Contexte : User
- Back
    - /profil method DELETE
- Model
    - table user

##### Connexion au compte
- Front
    - composant : Login
    - composant affichant le formulaire de connexion
    - Contexte : User
- Back
    - /login
    - verifier la validite des informations et envoyer un token jwt
- Model
    - table user
        - mail
        - password

#### Module : Sauvegarde recette

##### Ajout recette Like

- Front
  - composant : likeButton
  - Composant permettant d'ajouter la recette dans sa liste d'aimé (consultable dans les listes de recettes dans une section spéciale "Like")
  - Paramètre : idRecette
- Back
  - /addLike?idRecette=& (user)
  - Composant ajoutant la recette dans les likés de l'utilisateur
- Model
  - Table Playlist
    - Entités
      - idUser
      - NomPlaylist

##### Ajout recette liste

- Front
  - composant : addListButton
  - Composant permettant d'ajouter une rectte dans une (ou plusieurs) liste(s) que l'utilisateur aura créé précédemment
  - Paramètre : idRecette
- Back
  - /addList?idRecette=&idPlaylist=
  - Permet l'ajout de la recette dans une liste choisi par l'utilisateur dans la base de données
- Model
  - Table Playlist
  - Entités
    - idPlaylist

##### Gestion des listes de recettes

###### Affichage listes

- Front
  - page : displayLists
  - Page permettant l'affichage des listes de recettes liées au compte de l'utilisateur
  - Context : User
- Back
  - /displayLists
  - retourne les listes de recettes de l'utilisateur (y compris la liste des liké disposé différement des autres)
- Model
  - Table Playlist
    - Entités
      - idPlaylist
      - nomPlaylist
      - ?imagePlaylist

###### Affichage liste

- Front
  - page : playlist/:id
  - Page permettant l'affichage de la liste des recettes liées a la playlist
  - Paramètre : idPlaylist
- Back
  - /playlist?id=
  - retourne les recettes de la playlist
- Model
  - Table Playlist
    - Entités
      - idPlaylist
      - nomPlaylist
  - Table Recette
    - Entite
        - idRecette
        - nomRecette
        - voir tom design


##### Affichage d'une recette dans une liste

- Front
    - components : recipeListCard
    - composant affichant une recette dans une liste (like ou pas)
    - Paramètre : Recipe (données d'une recette (ex: nomRecette/imageRecette...))
- Model
  - table recette
    - entites
      - idRecette
      - nomRecette
      - urlImage
      - nbrLike
      - user a liker ou pas


##### Suppresion d'une recette dans une liste liste

- Front
    - components : deleteListRecipeButton
    - composant qui supprime la recipe d'une liste
    - Paramètre : idRecette
- Back
    - /playlist?id=&?=recipe& method post
- Model
    - table playlist
        - idplaylist
        - idRecipe


###### Création d'une liste :

- Front
  - composant : createListButton
  - Composant permettant à l'utilisateur de créer une nouvelle liste avec le nom de son choix
  - Contexte : User
- Back
  - /createList?nomPlaylist=&
  - Composant permettant l'ajout d'une nouvelle listes pour l'utilisateur (avec nom choisi par l'utilisateur)
- Model
  - Table Playlist
    - Entités
      - IdPlaylist
      - nomPlaylist

###### Modification d'une liste :

- Front
  - composant : modifyList
  - Composant permettant à l'utilisateur de modifier le nom
  - Paramètre : idPlaylist
- Back
  - /modifyList?idPlaylist=&?nomPlaylist=& (les informations)
  - Composant permettant de modifier le nom d'une liste
- Model
  - Table Playlist
    - Entités
      - idPlaylist
      - nomPlaylist

###### Supprimer une liste :

- Front
  - composant : deleteList
  - Composant permettant à l'utilisateur de supprimer une de ses liste
  - Paramètre : idPlaylist
- Back
  - /deleteList?idPlaylist=
  - Composant permettant la suppression total de la liste de recettes
- Model
  - Table Playlist

#### Module : suggestion

##### Global

###### Most liked
- Front
    - composant : MostLiked
    - composant affichant les recettes les plus likees
- Back
    - /mostLiked
    - retourne un nombre X de recettes les plus likees 
- Model
    - table recette
        - nomRecette
        - imageRecette
        - nbrLike
        - variable mostLiked (calcule)
    - table recette && user
        - user a liker ou pas
        - a ajouter a la playlist ou pas

###### Trending
- Front
    - composant : Trending
    - composant affichant les recettes les plus en vogue sur les 30 derniers jours
- Back
    - /trending
    - retourne les recettes les plus clique sur les 30 derniers jours
- Model
    - table statRecettes

#####  Context

###### Item by Item
- Front
    - composant : SimilarRecipes
    - composant affichant les recettes similaire a celle en ce moment
    - Paramètre : idPlaylist
- Back
    - /similarRecipes?idRecette=&
    - retourne les recettes les plus similaire en lien avec l'app python
- Model
    - neo4j graph

###### Drystart

- Front
  - composant : Drystart
  - composant posant les multiples questions demandées à la suite et envoye les donnees une fois toutes les informations reunies
- Back
  - /drystart
- Model
  - table user
    - prix
    - difficulte
    - epicee

##### Personnalise

###### Recommended for you
- Front
    - composant : Recommended for you
    - Contexte : User
- Back
    - /recommended
    - calcul magique
- Model
    - table user
        - preferences
    - neo4j graph

### Attente technique des besoins

## Organisation
