package api

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/utils"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) registerSuggestionRoutes(e *echo.Echo) {
	//TODO: Add query and implement similarRecipes drystart and recommended
	e.GET("/trending", s.handleTrending)
	e.GET("/mostliked", s.handleMostLiked)
	e.GET("/recommended", s.handleMostLiked)
	e.GET("/similarRecipes/:id/:number", s.handleSimilarRecipes)
}

func (s *APIServer) handleTrending(c echo.Context) error {
	resp, err := s.store.GetTrending()

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *APIServer) handleSimilarRecipes(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a page")
	}

	number, err := strconv.Atoi(c.Param("number"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a page")
	}

	resp, err := s.store.GetSimilarRecipes(id, number)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *APIServer) handleMostLiked(c echo.Context) error {
	resp, err := s.store.GetMostLiked()

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	return c.JSON(http.StatusOK, resp)
}
