package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	listenAddr string
	store      Storage
}

func NewAPIServer(listenAddr string, store Storage) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
		store:      store,
	}
}

func (s *APIServer) Run() {
	router := mux.NewRouter()

	router.HandleFunc("/auth/login", makeHTTPHandleFunc(s.handleLogin))
	router.HandleFunc("/auth/register", makeHTTPHandleFunc(s.handleRegister))
	router.HandleFunc("/auth/profil", withJWTAuth(makeHTTPHandleFunc(s.handleProfil), s.store))

	router.HandleFunc("/recipe/{id}", makeHTTPHandleFunc(s.handleGetRecipe))
	router.HandleFunc("/recipe/page/{page}", makeHTTPHandleFunc(s.handleRecipes))

	router.HandleFunc("/list", withJWTAuth(makeHTTPHandleFunc(s.handleList), s.store))

	log.Println("Piratecook api server running on port:", s.listenAddr)

	if err := http.ListenAndServe(s.listenAddr, router); err != nil {
		log.Fatal("Server failed to start: ", err)
	}
}
