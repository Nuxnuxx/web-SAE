package api

import (
	"fmt"
	"net/http"
	"strings"

	"backend/config"
	"backend/storage"
	"backend/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func withJWTAuth(next echo.HandlerFunc, s storage.Storage) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		finalToken := strings.Split(tokenString, " ")

		token, err := validateJWT(finalToken[1])
		if err != nil || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "permission denied")
		}

		claims := token.Claims.(jwt.MapClaims)
		c.Request().Header.Set("Mail", fmt.Sprintf("%v", claims["Mail"]))
		c.Request().Header.Set("FirstName", fmt.Sprintf("%v", claims["FirstName"]))
		c.Request().Header.Set("LastName", fmt.Sprintf("%v", claims["LastName"]))
		c.Request().Header.Set("Gender", fmt.Sprintf("%v", claims["Gender"]))
		c.Request().Header.Set("Password", fmt.Sprintf("%v", claims["Password"]))
		c.Request().Header.Set("price", fmt.Sprintf("%v", claims["price"]))
		c.Request().Header.Set("difficulty", fmt.Sprintf("%v", claims["difficulty"]))

		if err := s.FindAccountByMail(c.Request().Header.Get("Mail")); err == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "permission denied")
		}

		return next(c)
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
