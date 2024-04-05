package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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

	routes = append(routes, todoRoutes...)
	routes = append(routes, userRoutes...)

	for _, route := range routes {
		if route.Auth {
			// implement
		} else {
			router.HandleFunc(route.Uri, route.Function).Methods(route.Method)
		}
	}

	fileServe := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServe))
}
