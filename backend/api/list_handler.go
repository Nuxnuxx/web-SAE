package api

import (
	"backend/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) registerListRoutes(e *echo.Echo) {
	listRouter := e.Group("/list")

	listRouter.Use(JWTMiddleware(s.store))
	listRouter.GET("", s.handleGetList)
	listRouter.POST("", s.handleCreateList)
	listRouter.PUT("", s.handleUpdateList)
	listRouter.DELETE("", s.handleDeleteList)

	listRouter.GET("/recipe", s.handleGetListRecipe)
	listRouter.POST("/recipe", s.handleCreateListRecipe)
	listRouter.PUT("/recipe", s.handleUpdateListRecipe)
	listRouter.DELETE("/recipe", s.handleDeleteListRecipe)
}

func (s *APIServer) handleCreateListRecipe(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if query := c.QueryParams(); len(query) > 0 {
		id, err := strconv.Atoi(query.Get("id"))

		if err != nil {
			return err
		}

		resp, err := s.store.CreateRecipeLiked(mail, id)

		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}
		return c.JSON(http.StatusOK, resp)
	}

	return echo.NewHTTPError(http.StatusBadRequest, "Please give a id")
}

func (s *APIServer) handleGetListRecipe(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if query := c.QueryParams(); len(query) > 0 {
		id, err := strconv.Atoi(query.Get("id"))

		check := s.store.CheckListBelongToUser(id, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorDoesntBelong)
		}

		if err != nil {
			return err
		}

		resp, err := s.store.GetRecipeList(id)

		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}
		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a id")
	}
}

func (s *APIServer) handleUpdateListRecipe(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if query := c.QueryParams(); len(query) > 0 {
		id, err := strconv.Atoi(query.Get("id"))

		if err != nil {
			return err
		}

		idList, err := strconv.Atoi(query.Get("idlist"))

		if err != nil {
			return err
		}

		check := s.store.CheckListBelongToUser(idList, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorDoesntBelong)
		}

		resp, err := s.store.CreateRecipeList(mail, id, idList)

		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}
		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a id")
	}
}

func (s *APIServer) handleDeleteListRecipe(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if query := c.QueryParams(); len(query) > 0 {
		idList, err := strconv.Atoi(query.Get("idlist"))

		if err != nil {
			return err
		}

		id, err := strconv.Atoi(query.Get("id"))

		if err != nil {
			return err
		}

		check := s.store.CheckListBelongToUser(idList, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorDoesntBelong)
		}

		if err != nil {
			return err
		}
		resp, err := s.store.DeleteRecipeList(id, idList)

		if err != nil {
			fmt.Println(err)
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}
		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a id")
	}
}

func (s *APIServer) handleCreateList(c echo.Context) error {
	mail := c.Request().Header.Get("mail")
	if query := c.QueryParams(); len(query) > 0 {
		check := s.store.CheckListAlreadyExist(query.Get("name"), mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Playlist already exist with this name")
		}

		resp, err := s.store.CreateList(query.Get("name"), mail)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a name for the playlist")
	}
}
func (s *APIServer) handleGetList(c echo.Context) error {
	mail := c.Request().Header.Get("mail")
	resp, err := s.store.GetList(mail)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}
func (s *APIServer) handleUpdateList(c echo.Context) error {
	mail := c.Request().Header.Get("mail")
	if query := c.QueryParams(); len(query) > 0 {
		id, err := strconv.Atoi(query.Get("id"))

		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Please give a id for the playlist")
		}

		name := query.Get("name")

		check := s.store.CheckListBelongToUser(id, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorDoesntBelong)
		}

		check = s.store.CheckListAlreadyExist(name, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnprocessableEntity, "Playlist already exist with this name")
		}

		resp, err := s.store.UpdateList(id, name)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a name for the playlist")
	}
}
func (s *APIServer) handleDeleteList(c echo.Context) error {
	mail := c.Request().Header.Get("mail")
	if query := c.QueryParams(); len(query) > 0 {

		id, err := strconv.Atoi(query.Get("id"))

		check := s.store.CheckListBelongToUser(id, mail)

		if check != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorDoesntBelong)
		}

		resp, err := s.store.DeleteList(id)

		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
		}

		return c.JSON(http.StatusOK, resp)
	} else {
		return echo.NewHTTPError(http.StatusBadRequest, "Please give a name for the playlist")
	}
}
