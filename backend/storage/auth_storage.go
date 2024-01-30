package storage

import (
	"backend/types"
	"backend/utils"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AuthStorage interface {
	// Account
	CreateAccount(*types.Account) error
	GetAccountByMail(string) (*types.Account, error)
	DeleteAccount(string) error
	UpdateAccount(*types.Account) error
	FindAccountByMail(string) error
	Login(types.LoginRequest) (*types.Account, error)
	CreateColdstart(*types.CreateColdstartRequest, string) error
}

func (s *Neo4jStore) CreateAccount(acc *types.Account) error {

	resp, err := s.db.Run(s.ctx, "CREATE (u:User {name: $name, mail: $mail, password: $password, gender: $gender, price: $price, difficulty: $difficulty}) RETURN u",
		map[string]interface{}{
			"name":       acc.FirstName + " " + acc.LastName,
			"mail":       acc.Mail,
			"gender":     acc.Gender,
			"password":   acc.EncryptedPassword,
			"price":      acc.Price,
			"difficulty": acc.Difficulty,
		},
	)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil
}

func (s *Neo4jStore) Login(req types.LoginRequest) (*types.Account, error) {
	query := "MATCH (u:User {mail: $mail}) RETURN u"

	resp, err := s.db.Run(s.ctx, query,
		map[string]interface{}{
			"mail": req.Mail},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		data := extractProperty(*resp.Record(), "u", "password").(string)

		if err := bcrypt.CompareHashAndPassword([]byte(data), []byte(req.Password)); err != nil {
			return nil, fmt.Errorf(utils.ErrorInvalidCredentials)
		}

		account := CreateAccount(*resp.Record(), "u")

		return &account, nil
	}

	return nil, err
}

func (s *Neo4jStore) CreateColdstart(req *types.CreateColdstartRequest, mail string) error {
	query := `
		MATCH (u:User)
		WHERE u.mail = $mail
		SET u.price = $price
		SET u.difficulty = $difficulty
	`

	params := map[string]interface{}{
		"mail":       mail,
		"price":      req.Price,
		"difficulty": req.Difficulty,
	}

	resp, err := s.db.Run(s.ctx, query, params)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil
}

func (s *Neo4jStore) GetAccountByMail(mail string) (*types.Account, error) {
	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) return u;",
		map[string]interface{}{
			"mail": mail,
		},
	)

	if err != nil {
		return nil, err
	}

	if resp.Err() != nil {
		return nil, err
	}

	if resp.Next(s.ctx) {
		account := CreateAccount(*resp.Record(), "u")

		return &account, nil
	}

	return nil, nil
}

func (s *Neo4jStore) DeleteAccount(mail string) error {
	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) DETACH DELETE u;",
		map[string]interface{}{
			"mail": mail,
		},
	)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil
}

func (s *Neo4jStore) UpdateAccount(acc *types.Account) error {

	resp, err := s.db.Run(s.ctx, "MATCH (u:User {mail: $mail}) SET u.password = $password, u.name = $name RETURN u",
		map[string]interface{}{
			"name":     acc.FirstName + " " + acc.LastName,
			"mail":     acc.Mail,
			"password": acc.EncryptedPassword,
		},
	)

	if err != nil {
		return err
	}

	if resp.Err() != nil {
		return err
	}

	return nil
}

func (s *Neo4jStore) FindAccountByMail(mail string) error {
	query := "MATCH (u:User {mail: $mail}) RETURN u"

	resp, err := s.db.Run(s.ctx, query,
		map[string]interface{}{
			"mail": mail},
	)

	if err != nil {

		return err
	}

	if resp.Err() != nil {
		return err
	}

	if resp.Next(s.ctx) {
		return fmt.Errorf("Email found")
	}

	return nil
}
