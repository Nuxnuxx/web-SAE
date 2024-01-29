package api

import (
	"backend/types"
	"backend/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s *APIServer) registerAuthRoutes(e *echo.Echo) {
	authRouter := e.Group("/auth")
	authRouter.POST("/login", s.handleLogin)
	authRouter.POST("/register", s.handleRegister)
	authRouter.GET("/profil", withJWTAuth(s.handleProfil, s.store))
	authRouter.POST("/coldstart", withJWTAuth(s.handleColdStart, s.store))
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
