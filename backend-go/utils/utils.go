package utils

import (
	"backend/types"

	"golang.org/x/crypto/bcrypt"
)

func NewAccount(gender, firstName, lastName, mail, password, price, difficulty string) (*types.Account, error) {
	encpw, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &types.Account{
		FirstName:         firstName,
		LastName:          lastName,
		Mail:              mail,
		Gender:            gender,
		EncryptedPassword: string(encpw),
		Price:             price,
		Difficulty:        difficulty,
	}, nil
}

var ErrorNoId = "Please give a id"
var ErrorDoesntBelong = "List doesn't belong to the user"
var ErrorInternal = "Internal Server Error"
var ErrorInvalidCredentials = "Invalid credentials"
