package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	auth_controller "github.com/rcarvalho-pb/todo-app-go/internal/controller/auth"
)

const AUTH_RESOURCE = "/auth"

var authController auth_controller.AuthController

func InitAuthRoutes(db *sqlx.DB) []Route {
	authController = *auth_controller.NewAuthController(db)

	return authRoutes
}

var authRoutes = []Route{
	{
		Uri:      fmt.Sprintf("%s/login", AUTH_RESOURCE),
		Method:   http.MethodPost,
		Function: authController.Login,
		Auth:     false,
	},
}
