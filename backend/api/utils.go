package api

import (
	"backend/config"
	"backend/types"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(config.Jwt_token), nil
	})
}

func createJWT(account *types.Account) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt":  15000,
		"Mail":       account.Mail,
		"FirstName":  account.FirstName,
		"Gender":     account.Gender,
		"LastName":   account.LastName,
		"Password":   account.EncryptedPassword,
		"price":      account.Price,
		"difficulty": account.Difficulty,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(config.Jwt_token))
}
