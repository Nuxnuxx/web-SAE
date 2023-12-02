package main

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s Storage) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("x-jwt-token")

		token, err := validateJWT(tokenString)
		if err != nil {
			permissionDenied(w)
			return
		}

		if !token.Valid {
			permissionDenied(w)
			return
		}
		// userId, err := getId(r)
		//
		// if err != nil {
		// 	permissionDenied(w)
		// }
		// account, err := s.GetAccountById(userId)
		//
		// if err != nil {
		// 	permissionDenied(w)
		// 	return
		// }
		//
		// claims := token.Claims.(jwt.MapClaims)
		//
		// if account.IdUser != claims["idUser"] {
		// 	permissionDenied(w)
		// 	return
		// }
		//
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
		"Id":        account.IdUser,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(jwtSecret))
}
