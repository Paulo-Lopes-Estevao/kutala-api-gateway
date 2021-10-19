package middleware

import (
	"net/http"

	"github.com/Paulo-Lopes-Estevao/NZIMBUPAY-api-gateway/controller"
)

func RegistryMiddlewareBasicAuth(next http.HandlerFunc, c controller.AppController) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		username, password, ok := r.BasicAuth()
		if !ok {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		services, err := c.Service.AuthBasicService(username, password)

		if err != nil {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		result := services.VerifyPassword(password)

		isValid := (username == services.Username) && result
		if !isValid {
			http.Error(w, "Unauthorized.", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

	})

}
