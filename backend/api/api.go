package api

import (
	"log"
	"os"

	"backend/storage"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		MaxAge:           3600,
		AllowOrigins:     []string{os.Getenv("FRONT_URL")},
	}))

	s.registerAuthRoutes(e)
	s.registerListRoutes(e)
	s.registerRecipeRoutes(e)

	log.Println("Piratecook api server running on port:", s.listenAddr)
	e.Logger.Fatal(e.Start(s.listenAddr))
}
