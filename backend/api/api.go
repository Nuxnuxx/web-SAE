package api

import (
	"io"
	"os"

	"backend/config"
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

func (s *APIServer) SetupApi(e *echo.Echo) {
	// setup a single file of log to store it
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		e.Logger.Fatal("Error opening log file:", err)
	}
	defer logFile.Close()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{Output: logFile}))
	e.Logger.SetOutput(io.MultiWriter(os.Stdout, logFile))

	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowCredentials: true,
		MaxAge:           3600,
		AllowOrigins:     []string{os.Getenv("FRONT_URL")},
	}))

	if config.Stage == "prod" {
		e.HideBanner = true
		e.HidePort = true
	}
}

func (s *APIServer) Run() {
	e := echo.New()

	s.SetupApi(e)

	s.registerAuthRoutes(e)
	s.registerListRoutes(e)
	s.registerRecipeRoutes(e)
	s.registerSuggestionRoutes(e)

	e.Logger.Fatal(e.Start(s.listenAddr))
}
