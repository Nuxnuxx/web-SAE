package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		finalToken := strings.Split(tokenString, " ")

		token, err := validateJWT(finalToken[1])
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}

		claims := token.Claims.(jwt.MapClaims)
		r.Header.Set("Mail", fmt.Sprintf("%v", claims["Mail"]))
		r.Header.Set("FirstName", fmt.Sprintf("%v", claims["FirstName"]))
		r.Header.Set("LastName", fmt.Sprintf("%v", claims["LastName"]))
		r.Header.Set("Password", fmt.Sprintf("%v", claims["Password"]))

		if err := s.FindAccountByMail(r.Header.Get("Mail")); err == nil {
			permissionDenied(w)
			return
		}

		handlerFunc(w, r)
	}
}

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

func validateJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})
}

func createJWT(account *Account) (string, error) {
	claims := &jwt.MapClaims{
		"ExpiresAt": 15000,
		"Mail":      account.Mail,
		"FirstName": account.FirstName,
		"LastName":  account.LastName,
		"Password":  account.EncryptedPassword,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
