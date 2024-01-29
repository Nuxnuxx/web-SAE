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

	authRouter := e.Group("/auth")
	authRouter.POST("/login", s.handleLogin)
	authRouter.POST("/register", s.handleRegister)
	authRouter.GET("/profil", withJWTAuth(s.handleProfil, s.store))
	authRouter.POST("/coldstart", withJWTAuth(s.handleColdStart, s.store))

	recipeRouter := e.Group("/recipe")
	recipeRouter.GET("/:id", s.handleGetRecipe)
	recipeRouter.GET("/page/:page", s.handleRecipes)

	listRouter := e.Group("/list")
	listRouter.GET("", withJWTAuth(s.handleList, s.store))
	listRouter.GET("/recipe", withJWTAuth(s.handleListRecipe, s.store))

	//TODO: Add query and implement similarRecipes drystart and recommended
	e.GET("/trending", s.handleTrending)
	e.GET("/mostliked", s.handleMostLiked)
	e.GET("/recommended", s.handleMostLiked)
	e.GET("/similarRecipes/:id/:number", s.handleSimilarRecipes)

	log.Println("Piratecook api server running on port:", s.listenAddr)
	e.Logger.Fatal(e.Start(s.listenAddr))
}
