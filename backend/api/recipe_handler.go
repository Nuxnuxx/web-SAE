package api

import (
	"backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) registerRecipeRoutes(e *echo.Echo) {
	recipeRouter := e.Group("/recipe")
	recipeRouter.GET("/:id", s.handleGetRecipe)
	recipeRouter.GET("/page/:page", s.handleRecipes)
}

func (s *APIServer) handleGetRecipe(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a page")
	}

	resp, err := s.store.GetRecipeById(id)

	if err != nil {
		fmt.Println(err)
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}
	return c.JSON(http.StatusOK, resp)
}

func (s *APIServer) handleRecipes(c echo.Context) error {
	// vars from web are always a string need to convert it to Integer
	page, err := strconv.Atoi(c.Param("page"))

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a page")
	}

	// if queries present then filter the recipe
	if query := c.QueryParams(); len(query) > 0 {
		resp, err := s.store.GetRecipesWithFilter(page, query)

		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}

		return c.JSON(http.StatusOK, resp)
	}

	resp, err := s.store.GetRecipes(page)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	return c.JSON(http.StatusOK, resp)
}
