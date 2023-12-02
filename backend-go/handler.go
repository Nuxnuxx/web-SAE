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

	return writeJSON(w, http.StatusOK, req)
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccountRequest)

	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	account, err := newAccount(req.FirstName, req.LastName, req.Mail, req.Password)

	if err != nil {
		return err
	}

	if err := s.store.CreateAccount(account); err != nil {
		return err
	}

	token, err := createJWT(account)

	return writeJSON(w, http.StatusOK, token)
}

func (s *APIServer) handleGetRecipe(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, "Please give a id")
	}

	resp, err := s.store.GetRecipeById(id)

	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, "Internal Server error")
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleRecipes(w http.ResponseWriter, r *http.Request) error {
	page, err := strconv.Atoi(mux.Vars(r)["page"])

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, "Please give a id")
	}

	query := r.URL.Query()

	if query != nil {
		resp, err := s.store.GetRecipes(page, query)

		if err != nil {
			return writeJSON(w, http.StatusInternalServerError, ApiError{Error: err.Error()})
		}

		return writeJSON(w, http.StatusOK, resp)
	}

	resp, err := s.store.GetRecipes(page, nil)

	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, "Internal Server Error")
	}

	return writeJSON(w, http.StatusOK, resp)
}
