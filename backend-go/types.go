package main

import "golang.org/x/crypto/bcrypt"

type Recipe struct {
	RecipeDetail      RecipeDetail
	RecipeStep        RecipeStep
	RecipeIngredients RecipeIngredients
}

type IngredientInfo struct {
	Name       string
	URLPicture string
}

type RecipeIngredients map[int64]IngredientInfo

type RecipeStep map[int64]struct {
	Step string
}

type RecipeDetail struct {
	Difficulty string `json:"difficulty"`
	Quantity   string `json:"quantity"`
	Price      string `json:"price"`
	Name       string `json:"name"`
	IdRecipe   int64  `json:"idrecipe"`
}

type LoginRequest struct {
	Mail     string `json:"mail"`
	Password string `json:"password"`
}

type CreateAccountRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Mail      string `json:"mail"`
	Password  string `json:"password"`
}

type Account struct {
	IdUser            int    `json:"idUser"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Gender            string `json:"gender"`
	Price             string `json:"price"`
	EncryptedPassword string `json:"-"`
	Mail              string `json:"mail"`
}

func newAccount(firstName, lastName, mail, password string) (*Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &Account{
		FirstName:         firstName,
		LastName:          lastName,
		Mail:              mail,
		EncryptedPassword: string(encpw),
	}, nil
}
