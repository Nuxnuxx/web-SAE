package api

import (
	"log"
	"net/http"
	"os"

	"backend/storage"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      storage.Storage
}

func NewAPIServer(listenAddr string, store storage.Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	// CORS Settings
	credentials := handlers.AllowCredentials()
	ttl := handlers.MaxAge(3600)
	origin := handlers.AllowedOrigins([]string{os.Getenv("FRONT_URL")})

	router := mux.NewRouter()

	router.HandleFunc("/auth/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/auth/register", makeHTTPHandleFunc(s.handleRegister))
	router.HandleFunc("/auth/profil", withJWTAuth(makeHTTPHandleFunc(s.handleProfil), s.store))
	router.HandleFunc("/auth/coldstart", withJWTAuth(makeHTTPHandleFunc(s.handleColdStart), s.store))

	router.HandleFunc("/recipe/{id}", makeHTTPHandleFunc(s.handleGetRecipe))
	router.HandleFunc("/recipe/page/{page}", makeHTTPHandleFunc(s.handleRecipes))

	router.HandleFunc("/list", withJWTAuth(makeHTTPHandleFunc(s.handleList), s.store))
	router.HandleFunc("/list/recipe", withJWTAuth(makeHTTPHandleFunc(s.handleListRecipe), s.store))

	//TODO: Add query and implement similarRecipes drystart and recommended
	router.HandleFunc("/mostliked", makeHTTPHandleFunc(s.handleMostLiked))
	router.HandleFunc("/trending", makeHTTPHandleFunc(s.handleTrending))
	router.HandleFunc("/similarRecipes/{id}/{number}", makeHTTPHandleFunc(s.handleSimilarRecipes))
	router.HandleFunc("/drystart", makeHTTPHandleFunc(s.handleMostLiked))
	router.HandleFunc("/recommended", makeHTTPHandleFunc(s.handleMostLiked))

	log.Println("Piratecook api server running on port:", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, handlers.CORS(credentials, ttl, origin)(router)); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
