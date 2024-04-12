package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/rcarvalho-pb/todo-app-go/internal/middleware"
)

type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func ConfigRouter(router *mux.Router, db *sqlx.DB) {
	var routes []Route

	todoRoutes := InitTodoRoutes(db)
	userRoutes := InitUserRoutes(db)
	authRoutes := InitAuthRoutes(db)

	routes = append(routes, todoRoutes...)
	routes = append(routes, userRoutes...)
	routes = append(routes, authRoutes...)

	for _, route := range routes {

		if route.Auth {
			router.HandleFunc(route.Uri, middleware.Logger(middleware.Authenticate(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.Uri, middleware.Logger(route.Function)).Methods(route.Method)
		}
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))
}
