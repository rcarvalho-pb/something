package todo_iservice

import (
	todo_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/todo/dto"
	user_dto "github.com/rcarvalho-pb/todo-app-go/internal/core/user/dto"
)

type TodoService interface {
	CreateTodo(*)
}
