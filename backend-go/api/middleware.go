package api

import (
	"fmt"
	"net/http"
	"strings"

	"backend/config"
	"backend/storage"
	"backend/types"

	"github.com/golang-jwt/jwt/v4"
)

func withJWTAuth(handlerFunc http.HandlerFunc, s storage.Storage) http.HandlerFunc {

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
		r.Header.Set("Gender", fmt.Sprintf("%v", claims["Gender"]))
		r.Header.Set("Password", fmt.Sprintf("%v", claims["Password"]))
		r.Header.Set("price", fmt.Sprintf("%v", claims["price"]))
		r.Header.Set("difficulty", fmt.Sprintf("%v", claims["difficulty"]))

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