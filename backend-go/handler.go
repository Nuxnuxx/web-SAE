package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(req.Mail); err == nil {
		permissionDenied(w)
	}

	account, err := s.store.Login(req)

	if err != nil {
		permissionDenied(w)
	}

	token, err := createJWT(account)

	if err != nil {
		permissionDenied(w)
	}

	finalResponse := APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResponse)
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	account, err := NewAccount(req.FirstName, req.LastName, req.Mail, req.Password)

	if err != nil {
		return err
	}
	if err := s.store.FindAccountByMail(account.Mail); err != nil {
		return err
	}

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	token, err := createJWT(account)

	finalResult := APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResult)
}

func (s *APIServer) handleGetRecipe(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, "Please give a id")
	}

	resp, err := s.store.GetRecipeById(id)

	if err != nil {
		fmt.Println(err)
		return writeJSON(w, http.StatusInternalServerError, "Internal Server error")
	}
	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleRecipes(w http.ResponseWriter, r *http.Request) error {
	// vars from web are always a string need to convert it to Integer
	page, err := strconv.Atoi(mux.Vars(r)["page"])

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, "Please give a id")
	}

	// if queries present then filter the recipe
	if query := r.URL.Query(); len(query) > 0 {
		resp, err := s.store.GetRecipesWithFilter(page, query)

		if err != nil {
			fmt.Println(err)
			return writeJSON(w, http.StatusInternalServerError, "Internal Server Error")
		}

		return writeJSON(w, http.StatusOK, resp)
	}

	resp, err := s.store.GetRecipes(page)

	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, "Internal Server Error")
	}

	return writeJSON(w, http.StatusOK, resp)
}
