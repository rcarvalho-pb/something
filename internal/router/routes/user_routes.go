package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	user_controller "github.com/rcarvalho-pb/todo-app-go/internal/controller/user"
	user_repository "github.com/rcarvalho-pb/todo-app-go/internal/core/user/port/repository"
	user_service "github.com/rcarvalho-pb/todo-app-go/internal/core/user/port/service"
	user_iservice "github.com/rcarvalho-pb/todo-app-go/internal/core/user/service"
	user_irepository "github.com/rcarvalho-pb/todo-app-go/internal/infra/repository/user"
)

const USER_RESOURCE = "users"

var userController user_controller.UserController

func InitUserRoutes(db *sqlx.DB) []Route {
	var rep user_repository.UserRepository = user_irepository.NewUserRepository(db)
	var serv user_service.UserService = user_iservice.NewUserService(rep)
	userController = *user_controller.NewUserControler(serv)

	return UserRoutes
}

var UserRoutes = []Route{
	{
		Uri:      fmt.Sprintf("/%s", USER_RESOURCE),
		Method:   http.MethodPost,
		Function: userController.CreateUser,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:   http.MethodPut,
		Function: userController.UpdateUser,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:   http.MethodPatch,
		Function: userController.UpdateUserPassword,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/all-users", USER_RESOURCE),
		Method:   http.MethodGet,
		Function: userController.FindAllUsers,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s", USER_RESOURCE),
		Method:   http.MethodGet,
		Function: userController.FindAllActiveUsers,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:   http.MethodGet,
		Function: userController.FindUserById,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", USER_RESOURCE),
		Method:   http.MethodDelete,
		Function: userController.DeleteUserById,
		Auth:     false,
	},
}
