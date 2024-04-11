package middleware

import (
	"log"
	"net/http"

	"github.com/rcarvalho-pb/todo-app-go/internal/authentication"
	response_status "github.com/rcarvalho-pb/todo-app-go/internal/core/user/entity_status_response"
	json_response "github.com/rcarvalho-pb/todo-app-go/internal/infra/response"
)

func Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, r.Host)

		next(w, r)
	}
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := authentication.ValidateToken(r); err != nil {
			json_response.ERROR(w, response_status.Unauthorized, err)
			return
		}
		next(w, r)
	}
}
