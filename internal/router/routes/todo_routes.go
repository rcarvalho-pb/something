package routes

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	todo_controller "github.com/rcarvalho-pb/todo-app-go/internal/controller/todo"
	todo_service "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/service"
	todo_repository "github.com/rcarvalho-pb/todo-app-go/internal/infra/repository/todo"
)

const TODO_RESOURCE = "todos"

var todoController *todo_controller.TodoController

func InitTodoRoutes(db *sqlx.DB) []Route {
	todoController = todo_controller.NewTodoController(todo_service.NewTodoService(todo_repository.NewTodoRepository(db)))

	return todoRoutes
}

var todoRoutes = []Route{
	{
		Uri:      fmt.Sprintf("/%s", TODO_RESOURCE),
		Method:   http.MethodPost,
		Function: todoController.CreateTodo,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:   http.MethodPut,
		Function: todoController.UpdateTodo,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:   http.MethodPatch,
		Function: todoController.UpdateTodoUsers,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/all-todos", TODO_RESOURCE),
		Method:   http.MethodGet,
		Function: todoController.FindAllTodos,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s", TODO_RESOURCE),
		Method:   http.MethodGet,
		Function: todoController.FindAllActiveTodos,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}/user", TODO_RESOURCE),
		Method:   http.MethodGet,
		Function: todoController.FindByUserId,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:   http.MethodGet,
		Function: todoController.FindById,
		Auth:     false,
	},
	{
		Uri:      fmt.Sprintf("/%s/{id}", TODO_RESOURCE),
		Method:   http.MethodDelete,
		Function: todoController.DeleteById,
		Auth:     false,
	},
}
