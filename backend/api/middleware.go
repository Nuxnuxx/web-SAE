package api

import (
	"fmt"
	"net/http"
	"strings"

	"backend/storage"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(s storage.Storage) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
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
}
