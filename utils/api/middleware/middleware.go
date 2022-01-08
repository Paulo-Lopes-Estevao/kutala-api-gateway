package middleware

import (
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/bootstrap"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/interface/controller"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegistryMiddlewareJWT() echo.MiddlewareFunc {

	signingKey := bootstrap.GoDotEnvVariable("JWT_SECRET_KEY")

	configJWT := middleware.JWTConfig{
		TokenLookup: "header:Authorization",
		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
			keyFunc := func(t *jwt.Token) (interface{}, error) {
				if t.Method.Alg() != "HS256" {
					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
				}
				return signingKey, nil
			}

			token, err := jwt.Parse(auth, keyFunc)
			if err != nil {
				return nil, err
			}
			if !token.Valid {
				return nil, errors.New("invalid token")
			}
			return token, nil
		},
	}

	return middleware.JWTWithConfig(configJWT)

}

func RegistryMiddlewareBasicAuth(e *echo.Echo, c controller.AppController) echo.MiddlewareFunc {

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(context echo.Context) error {

			username, password, ok := context.Request().BasicAuth()

			if !ok {
				return context.JSON(401, "Unauthorized.")
			}

			users, err := c.User.AuthBasicUser(username, password, context)

			if err != nil {
				return context.JSON(401, "Unauthorized.")
			}

			result := users.VerifyPassword(password)

			if subtle.ConstantTimeCompare([]byte(username), []byte(users.Username)) != 1 && !result {
				return context.JSON(401, "Unauthorized.")
			}
			return next(context)
		}
	}
}
