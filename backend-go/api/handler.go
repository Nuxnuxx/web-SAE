package api

import (
	"fmt"
	"net/http"
	"strconv"

	"backend/types"
	"backend/utils"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) handleListRecipe(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if c.Request().Method == "POST" {
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
		} else {
			return echo.NewHTTPError(http.StatusBadRequest, "Please give a id")
		}
	}

	if c.Request().Method == "PUT" {
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

	if c.Request().Method == "GET" {
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

	if c.Request().Method == "DELETE" {
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

	return echo.NewHTTPError(http.StatusMethodNotAllowed, utils.ErrorMethodNotAllowed)
}

func (s *APIServer) handleList(c echo.Context) error {
	mail := c.Request().Header.Get("Mail")
	if c.Request().Method == "POST" {
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

	if c.Request().Method == "GET" {
		resp, err := s.store.GetList(mail)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, resp)
	}

	if c.Request().Method == "PUT" {
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

	if c.Request().Method == "DELETE" {
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

	return echo.NewHTTPError(http.StatusMethodNotAllowed, utils.ErrorMethodNotAllowed)
}

func (s *APIServer) handleUpdateProfil(c echo.Context) error {
	var req types.UpdateProfilRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	account, err := utils.NewAccount(c.Request().Header.Get("Gender"), req.FirstName, req.LastName, c.Request().Header.Get("Mail"), req.NewPassWord, c.Request().Header.Get("price"), c.Request().Header.Get("difficulty"))

	if err != nil {
		return err
	}

	//FIXME: need to check if the account exists and the current password is correct

	if err := s.store.UpdateAccount(account); err != nil {
		return err
	}

	token, err := createJWT(account)

	if err != nil {
		return err
	}

	response := types.APIResponse{
		Result: token,
	}

	return c.JSON(http.StatusOK, response)
}

func (s *APIServer) handleDeleteProfil(c echo.Context) error {
	if err := s.store.FindAccountByMail(c.Request().Header.Get("Mail")); err == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorInvalidCredentials)
	}

	if err := s.store.DeleteAccount(c.Request().Header.Get("Mail")); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	response := types.APIResponse{
		Result: "Account deleted",
	}

	return c.JSON(http.StatusOK, response)
}

func (s *APIServer) handleGetProfil(c echo.Context) error {
	user := types.Account{
		FirstName:  c.Request().Header.Get("firstName"),
		LastName:   c.Request().Header.Get("lastName"),
		Gender:     c.Request().Header.Get("gender"),
		Mail:       c.Request().Header.Get("mail"),
		Price:      c.Request().Header.Get("price"),
		Difficulty: c.Request().Header.Get("difficulty"),
	}

	return c.JSON(http.StatusOK, user)
}

func (s *APIServer) handleProfil(c echo.Context) error {
	if c.Request().Method == "PUT" {
		return s.handleUpdateProfil(c)
	}
	if c.Request().Method == "DELETE" {
		return s.handleDeleteProfil(c)
	}

	if c.Request().Method == "GET" {
		return s.handleGetProfil(c)
	}

	return c.JSON(http.StatusMethodNotAllowed, utils.ErrorMethodNotAllowed)
}

func (s *APIServer) handleLogin(c echo.Context) error {
	if c.Request().Method != "POST" {
		return echo.NewHTTPError(http.StatusMethodNotAllowed, "method not allowed")
	}

	var req types.LoginRequest
	if err := c.Bind(&req); err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(req.Mail); err == nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorInvalidCredentials)
	}

	account, err := s.store.Login(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorInvalidCredentials)
	}

	token, err := createJWT(account)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
	}

	finalResponse := types.APIResponse{
		Result: token,
	}

	return c.JSON(http.StatusOK, finalResponse)
}

func (s *APIServer) handleRegister(c echo.Context) error {
	req := new(types.CreateAccountRequest)

	// decode the body and store it in the req variable
	if err := c.Bind(&req); err != nil {
		return err
	}

	account, err := utils.NewAccount(req.Gender, req.FirstName, req.LastName, req.Mail, req.Password, "0", "0")

	if err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(account.Mail); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, utils.ErrorInvalidCredentials)
	}

	if err := s.store.CreateAccount(account); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, utils.ErrorInternal)
	}

	token, err := createJWT(account)

	finalResult := types.APIResponse{
		Result: token,
	}

	return c.JSON(http.StatusOK, finalResult)
}

func (s *APIServer) handleColdStart(c echo.Context) error {
	req := new(types.CreateColdstartRequest)

	// decode the body and store it in the req variable
	if err := c.Bind(req); err != nil {
		return err
	}

	// the account is already created in the register function

	// create the coldstart
	if err := s.store.CreateColdstart(req, c.Request().Header.Get("Mail")); err != nil {
		return err
	}

	// make token with the all the old information + the new one
	account, err := utils.NewAccount(c.Request().Header.Get("Gender"), c.Request().Header.Get("FirstName"), c.Request().Header.Get("LastName"),
		c.Request().Header.Get("Mail"), c.Request().Header.Get("Password"), req.Price, req.Difficulty)

	if err != nil {
		return err
	}

	token, err := createJWT(account)

	finalResult := types.APIResponse{
		Result: token,
	}

	return c.JSON(http.StatusOK, finalResult)
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
