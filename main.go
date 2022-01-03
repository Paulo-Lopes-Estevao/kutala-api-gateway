package main

import (
	"fmt"
	"log"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/injection"

	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/utils/api/handler"
	"github.com/Paulo-Lopes-Estevao/kutala-api-gateway/utils/database"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	server := echo.New()

	db := database.ConnectionDB()
	i := injection.NewRegistry(db)

	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	p := prometheus.NewPrometheus("echo", nil)

	handler.RouteMicroserviceHandler(server, i.NewAppController())

	handler.RouteUserHandler(server, i.NewAppController())

	p.Use(server)

	/* 	signingKey := []byte("secret")

	   	config := middleware.JWTConfig{
	   		TokenLookup: "query:token",
	   		ParseTokenFunc: func(auth string, c echo.Context) (interface{}, error) {
	   			keyFunc := func(t *jwt.Token) (interface{}, error) {
	   				if t.Method.Alg() != "HS256" {
	   					return nil, fmt.Errorf("unexpected jwt signing method=%v", t.Header["alg"])
	   				}
	   				return signingKey, nil
	   			}

	   			// claims are of type `jwt.MapClaims` when token is created with `jwt.Parse`
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

	   	server.Use(middleware.JWTWithConfig(config)) */

	fmt.Println("Server started at port 9000")
	if err := server.Start(":9000"); err != nil {
		log.Println("Not Running Server A...", err.Error())
	}
}
