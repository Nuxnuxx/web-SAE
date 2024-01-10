package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"backend/types"
	"backend/utils"

	"github.com/gorilla/mux"
)

func (s *APIServer) handleListRecipe(w http.ResponseWriter, r *http.Request) error {
	mail := r.Header.Get("Mail")
	if r.Method == "POST" {
		if query := r.URL.Query(); len(query) > 0 {
			id, err := strconv.Atoi(query.Get("id"))

			if err != nil {
				return err
			}

			resp, err := s.store.CreateRecipeLiked(mail, id)

			if err != nil {
				fmt.Println(err)
				return fmt.Errorf(utils.ErrorInternal)
			}
			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf(utils.ErrorNoId)
		}
	}

	if r.Method == "PUT" {
		if query := r.URL.Query(); len(query) > 0 {
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
				return fmt.Errorf(utils.ErrorDoesntBelong)
			}

			resp, err := s.store.CreateRecipeList(mail, id, idList)

			if err != nil {
				fmt.Println(err)
				return fmt.Errorf(utils.ErrorInternal)
			}
			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf(utils.ErrorNoId)
		}
	}

	if r.Method == "GET" {
		if query := r.URL.Query(); len(query) > 0 {
			id, err := strconv.Atoi(query.Get("id"))

			check := s.store.CheckListBelongToUser(id, mail)

			if check != nil {
				return fmt.Errorf(utils.ErrorDoesntBelong)
			}

			if err != nil {
				return err
			}
			resp, err := s.store.GetRecipeList(id)

			if err != nil {
				fmt.Println(err)
				return fmt.Errorf(utils.ErrorInternal)
			}
			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf(utils.ErrorNoId)
		}
	}

	if r.Method == "DELETE" {
		if query := r.URL.Query(); len(query) > 0 {
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
				return fmt.Errorf(utils.ErrorDoesntBelong)
			}

			if err != nil {
				return err
			}
			resp, err := s.store.DeleteRecipeList(id, idList)

			if err != nil {
				fmt.Println(err)
				return fmt.Errorf(utils.ErrorInternal)
			}
			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf(utils.ErrorNoId)
		}
	}

	return writeJSON(w, http.StatusMethodNotAllowed, "Method not allowed")
}

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
				return fmt.Errorf(utils.ErrorInternal)
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
				return fmt.Errorf(utils.ErrorNoId)
			}

			name := query.Get("name")

			check := s.store.CheckListBelongToUser(id, mail)

			if check != nil {
				return fmt.Errorf(utils.ErrorDoesntBelong)
			}

			check = s.store.CheckListAlreadyExist(name, mail)

			if check != nil {
				return fmt.Errorf("List with this name already exists")
			}

			resp, err := s.store.UpdateList(id, name)

			if err != nil {
				return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
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
				return fmt.Errorf(utils.ErrorDoesntBelong)
			}

			resp, err := s.store.DeleteList(id)

			if err != nil {
				return fmt.Errorf(utils.ErrorInternal)
			}

			return writeJSON(w, http.StatusOK, resp)
		} else {
			return fmt.Errorf("Please give a name for the playlist")
		}
	}

	return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
}

func (s *APIServer) handleUpdateProfil(w http.ResponseWriter, r *http.Request) error {
	var req types.UpdateProfilRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	account, err := utils.NewAccount(r.Header.Get("Gender"), req.FirstName, req.LastName, r.Header.Get("Mail"), req.NewPassWord, "0", "0")

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

	return writeJSON(w, http.StatusOK, response)
}

func (s *APIServer) handleDeleteProfil(w http.ResponseWriter, r *http.Request) error {
	if err := s.store.FindAccountByMail(r.Header.Get("Mail")); err == nil {
		return fmt.Errorf("Account not found")
	}

	if err := s.store.DeleteAccount(r.Header.Get("Mail")); err != nil {
		return fmt.Errorf(utils.ErrorInternal)
	}

	response := types.APIResponse{
		Result: "Account deleted",
	}

	return writeJSON(w, http.StatusOK, response)
}

func (s *APIServer) handleGetProfil(w http.ResponseWriter, r *http.Request) error {
	user := types.Account{
		FirstName:  r.Header.Get("firstName"),
		LastName:   r.Header.Get("lastName"),
		Gender:     r.Header.Get("gender"),
		Mail:       r.Header.Get("mail"),
		Price:      r.Header.Get("price"),
		Difficulty: r.Header.Get("difficulty"),
	}

	return writeJSON(w, http.StatusOK, user)
}

func (s *APIServer) handleProfil(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "PUT" {
		return s.handleUpdateProfil(w, r)
	}
	if r.Method == "DELETE" {
		return s.handleDeleteProfil(w, r)
	}

	if r.Method == "GET" {
		return s.handleGetProfil(w, r)
	}

	return writeJSON(w, http.StatusBadRequest, "Method not allowed")
}
func (s *APIServer) handleLogin(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		return fmt.Errorf("method not allowed %s", r.Method)
	}

	var req types.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return err
	}

	if err := s.store.FindAccountByMail(req.Mail); err == nil {
		return fmt.Errorf(utils.ErrorInvalidCredentials)
	}

	account, err := s.store.Login(req)

	if err != nil {
		return fmt.Errorf(utils.ErrorInvalidCredentials)
	}

	token, err := createJWT(account)

	if err != nil {
		return fmt.Errorf(utils.ErrorInvalidCredentials)
	}

	finalResponse := types.APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResponse)
}

func (s *APIServer) handleRegister(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateAccountRequest)

	// decode the body and store it in the req variable
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	account, err := utils.NewAccount(req.Gender, req.FirstName, req.LastName, req.Mail, req.Password, "0", "0")

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

	finalResult := types.APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResult)
}

func (s *APIServer) handleColdStart(w http.ResponseWriter, r *http.Request) error {
	req := new(types.CreateColdstartRequest)

	// decode the body and store it in the req variable
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		return err
	}

	// the account is already created in the register function

	// create the coldstart
	if err := s.store.CreateColdstart(req, r.Header.Get("Mail")); err != nil {
		return err
	}

	// make token with the all the old information + the new one
	account, err := utils.NewAccount(r.Header.Get("Gender"), r.Header.Get("FirstName"), r.Header.Get("LastName"),
		r.Header.Get("Mail"), r.Header.Get("Password"), req.Price, req.Difficulty)

	if err != nil {
		return err
	}

	token, err := createJWT(account)

	finalResult := types.APIResponse{
		Result: token,
	}

	return writeJSON(w, http.StatusOK, finalResult)
}

func (s *APIServer) handleGetRecipe(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(mux.Vars(r)["id"])

	if err != nil {
		return writeJSON(w, http.StatusBadRequest, utils.ErrorNoId)
	}

	resp, err := s.store.GetRecipeById(id)

	if err != nil {
		fmt.Println(err)
		return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
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
			return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
		}

		return writeJSON(w, http.StatusOK, resp)
	}

	resp, err := s.store.GetRecipes(page)

	if err != nil {
		return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleTrending(w http.ResponseWriter, r *http.Request) error {
	resp, err := s.store.GetTrending()

	if err != nil {
		fmt.Println(err)
		return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
	}

	return writeJSON(w, http.StatusOK, resp)
}

func (s *APIServer) handleMostLiked(w http.ResponseWriter, r *http.Request) error {
	resp, err := s.store.GetMostLiked()

	if err != nil {
		fmt.Println(err)
		return writeJSON(w, http.StatusInternalServerError, utils.ErrorInternal)
	}

	return writeJSON(w, http.StatusOK, resp)
}
