package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleList(w http.ResponseWriter, r *http.Request) error {
	mail := r.Header.Get("Mail")
	if r.Method == "POST" {
		if query := r.URL.Query(); len(query) > 0 {
			check := s.store.CheckListAlreadyExist(query.Get("name"), mail)

			if check != nil {
				return fmt.Errorf("Playlist already exist with this name")
			}

			resp, err := s.store.CreateList(query.Get("name"), mail)

			if err != nil {
				return fmt.Errorf("Internal server error")
			}

			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf("Please give a name for the playlist")
		}
	}

	if r.Method == "GET" {
		resp, err := s.store.GetList(mail)

		if err != nil {
			return err
		}

		return writeJSON(w, http.StatusOK, resp)
	}

	if r.Method == "PUT" {
		if query := r.URL.Query(); len(query) > 0 {
			id, err := strconv.Atoi(query.Get("id"))

			if err != nil {
				return fmt.Errorf("Please give a id")
			}

			name := query.Get("name")

			check := s.store.CheckListBelongToUser(id, mail)

			if check != nil {
				return fmt.Errorf("List doesnt belong to the user")
			}

			check = s.store.CheckListAlreadyExist(name, mail)

			if check != nil {
				return fmt.Errorf("List with this name already exists")
			}

			resp, err := s.store.UpdateList(id, name)

			if err != nil {
				return writeJSON(w, http.StatusInternalServerError, "Internal Server Error")
			}

			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf("Please give id and name for the playlist")
		}
	}

	if r.Method == "DELETE" {
		if query := r.URL.Query(); len(query) > 0 {

			id, err := strconv.Atoi(query.Get("id"))

			check := s.store.CheckListBelongToUser(id, mail)

			if check != nil {
				return fmt.Errorf("List doesnt belong to the user")
			}

			resp, err := s.store.DeleteList(id)

			if err != nil {
				return fmt.Errorf("Internal server error")
			}

			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf("Please give a name for the playlist")
		}
	}

	return writeJSON(w, http.StatusInternalServerError, "Internal Server Error")
}

func (s *APIServer) handleUpdateProfil(w http.ResponseWriter, r *http.Request) error {
	var req UpdateProfilRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	account, err := NewAccount(r.Header.Get("FirstName"), r.Header.Get("LastName"), r.Header.Get("Mail"), req.NewPassWord)

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

	response := APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, response)
}

func (s *APIServer) handleDeleteProfil(w http.ResponseWriter, r *http.Request) error {
	if err := s.store.FindAccountByMail(r.Header.Get("Mail")); err == nil {
		return fmt.Errorf("Account not found")
	}

	if err := s.store.DeleteAccount(r.Header.Get("Mail")); err != nil {
		return fmt.Errorf("Internal Server Error")
	}

	response := APIResponse{
		Result: "Account deleted",
	}

	return writeJSON(w, http.StatusOK, response)
}

func (s *APIServer) handleProfil(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "PUT" {
		return s.handleUpdateProfil(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteProfil(w, r)
	}

	return writeJSON(w, http.StatusBadRequest, "Method not allowed")
}
func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(req.Mail); err == nil {
		return fmt.Errorf("Invalid Credentials")
	}

	account, err := s.store.Login(req)

	if err != nil {
		return fmt.Errorf("Invalid Credentials")
	}

	token, err := createJWT(account)

	if err != nil {
		return fmt.Errorf("Invalid Credentials")
	}

	finalResponse := APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResponse)
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	req := new(CreateAccountRequest)

	// decode the body and store it in the req variable
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	account, err := NewAccount(req.FirstName, req.LastName, req.Mail, req.Password)

	if err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(account.Mail); err != nil {
		return fmt.Errorf("Account already exists")
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
		return writeJSON(w, http.StatusBadRequest, "Please give a page")
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
